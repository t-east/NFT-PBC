FROM golang:1.16.2-alpine

ENV ROOT=/home
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
ADD ./go.mod ./go.sum ./
RUN go mod download

CMD ["go", "run", "main.go"]