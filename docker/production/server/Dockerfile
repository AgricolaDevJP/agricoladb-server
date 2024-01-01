FROM golang:1.20-bullseye as builder

WORKDIR /opt/app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o "agricoladb-server" ./cmd/server

FROM gcr.io/distroless/base-debian11

ENV PORT 8000

EXPOSE $PORT

COPY --from=builder /opt/app/agricoladb-server /agricoladb-server
CMD ["/agricoladb-server"]
