FROM --platform=$BUILDPLATFORM golang:1.21-bookworm as builder
ARG TARGETARCH

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o agricoladb-server ./cmd/server && \
  CGO_ENABLED=0 go run ./cmd/migration

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=builder --chown=nonroot:nonroot /opt/app/agricoladb-server /agricoladb-server
COPY --from=builder --chown=nonroot:nonroot /opt/app/agricoladb.sqlite /agricoladb.sqlite

ENV PORT 8000
EXPOSE $PORT
ENV DB_PATH /agricoladb.sqlite

ENTRYPOINT ["/agricoladb-server"]
