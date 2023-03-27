FROM golang:1.20-bullseye as builder

WORKDIR /opt/app

# RUN go env -w GOPROXY=direct
COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .
RUN ls -a
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o agricoladb-server ./cmd/server-lambda
RUN CGO_ENABLED=0 go run ./cmd/migration

FROM gcr.io/distroless/base-debian11

COPY --from=builder /opt/app/agricoladb-server /agricoladb-server
COPY --from=builder /opt/app/agricoladb.sqlite /agricoladb.sqlite
ENTRYPOINT ["/agricoladb-server"]   
