FROM golang:1.21-bookworm as builder

WORKDIR /opt/app

# RUN go env -w GOPROXY=direct
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o agricoladb-server ./cmd/server-lambda
RUN CGO_ENABLED=0 go run ./cmd/migration

FROM gcr.io/distroless/base-debian12

COPY --from=builder /opt/app/agricoladb-server /agricoladb-server
COPY --from=builder /opt/app/agricoladb.sqlite /agricoladb.sqlite
ENTRYPOINT ["/agricoladb-server"]   
