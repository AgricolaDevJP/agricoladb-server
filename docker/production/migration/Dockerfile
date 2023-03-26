FROM golang:1.20-bullseye as builder

WORKDIR /opt/app

COPY src .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o "agricoladb-server-migration" ./cmd/migration

FROM gcr.io/distroless/base-debian11

COPY --from=builder /opt/app/agricoladb-server-migration /agricoladb-server-migration
CMD ["/agricoladb-server-migration"]
