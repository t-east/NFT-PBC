FROM golang:1.16.4-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR /work/

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080