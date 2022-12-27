FROM golang:1.19-alpine as builder

ENV CGO_ENABLED 0
WORKDIR /go/src/agricoladb-server

COPY . .

RUN apk add make gcc g++

RUN make build

FROM alpine
COPY --from=builder /go/src/agricoladb-server/server /server

ENV DB_NAME agricoladb
ENV DB_HOST db
ENV DB_USER user
ENV DB_PASSWORD password

EXPOSE 8080

ENTRYPOINT ["/server"]
