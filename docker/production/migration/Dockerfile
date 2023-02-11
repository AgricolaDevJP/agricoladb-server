FROM golang:1.19-bullseye as builder

WORKDIR /opt/app

COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o "agricoladb-server-migration" ./cmd/migration

FROM gcr.io/distroless/base-debian11

COPY --from=builder /opt/app/agricoladb-server-migration /agricoladb-server-migration
COPY ./masterdata/ /
CMD ["/agricoladb-server-migration"]
