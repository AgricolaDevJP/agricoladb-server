FROM golang:1.20-bullseye

ENV CGO_ENABLED 0
WORKDIR /go/src/agricoladb-server

RUN go install github.com/cosmtrek/air@latest

# COPY src/go.mod src/go.sum ./
# RUN go mod download

# COPY src/ ./

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
