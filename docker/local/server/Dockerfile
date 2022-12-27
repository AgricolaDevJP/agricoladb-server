FROM golang:1.19-alpine

ENV CGO_ENABLED 0
WORKDIR /go/src/agricoladb-server

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
