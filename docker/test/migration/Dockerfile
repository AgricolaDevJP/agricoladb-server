FROM golang:1.19-alpine as builder

ENV CGO_ENABLED 0
WORKDIR /go/src/agricoladb-server

COPY . .

RUN apk add make gcc g++

RUN make build-migration

FROM gcr.io/distroless/base-debian11
COPY --from=builder /go/src/agricoladb-server/migration /migration

COPY ./masterdata/ ./masterdata/

ENV DB_NAME agricoladb
ENV DB_HOST db
ENV DB_USER user
ENV DB_PASSWORD password

ENTRYPOINT ["/migration"]
