# rpm-based image may be better for al2023
FROM --platform=$BUILDPLATFORM golang:1.22-bookworm as builder
ARG TARGETARCH

ENV GOTOOLCHAIN=auto

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o agricoladb-server ./cmd/server-lambda && \
  CGO_ENABLED=0 go run ./cmd/migration

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=builder /opt/app/agricoladb-server /agricoladb-server
COPY --from=builder /opt/app/agricoladb.sqlite /agricoladb.sqlite

ENV PORT 8000
EXPOSE $PORT
ENV DB_PATH /agricoladb.sqlite

ENTRYPOINT ["/agricoladb-server"]
