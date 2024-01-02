FROM golang:1.21-bookworm as builder

WORKDIR /opt/app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o "agricoladb-server-migration" ./cmd/migration

FROM gcr.io/distroless/base-debian12

COPY --from=builder /opt/app/agricoladb-server-migration /agricoladb-server-migration
CMD ["/agricoladb-server-migration"]
