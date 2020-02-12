FROM node:alpine as frontend
WORKDIR /src
ENV PATH /src/node_modules/.bin:$PATH

# Download deps
COPY frontend/package.json package.json
RUN \
    npm install && \
    npm install react-scripts -g
# Build frontend
COPY frontend/ .
RUN npm run build


FROM golang:alpine as static
WORKDIR /go/src/github.com/ramosisw/dashboard

RUN apk add git gcc build-base
# Download deps
ADD go.mod .
ADD go.sum .
RUN go mod download


ADD . .
COPY --from=frontend /src/build/ frontend/build
RUN go get -u github.com/go-bindata/go-bindata/...
RUN go get github.com/google/wire/cmd/wire
RUN go generate ./...
# RUN go test ./... -cover --coverprofile=coverage.out && \
#     go tool cover -func coverage.out

RUN go build -v -ldflags "-extldflags -static" -o /release/dashboard

FROM alpine as production
LABEL maintainer="Julio Ramos <ramos.isw@gmail.com>"

RUN apk add --no-cache ca-certificates
COPY --from=static /release/dashboard /usr/bin/dashboard

EXPOSE 80
VOLUME ["/data/"]
ENTRYPOINT dashboard