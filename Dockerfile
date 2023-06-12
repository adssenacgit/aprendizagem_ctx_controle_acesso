ARG GO_VERSION=1.19.2
ARG ALPINE_VERSION=3.16

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./cmd/main.go

FROM alpine:${ALPINE_VERSION}

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .
COPY ./certs ./certs/

EXPOSE 8080

ENTRYPOINT ["./app"]