FROM public.ecr.aws/lambda/provided:al2 as builder

RUN yum install -y golang

WORKDIR /opt/app

RUN go env -w GOPROXY=direct
ADD src/go.mod src/go.sum ./
RUN go mod download

ADD src .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o agricoladb-server ./cmd/server-lambda
RUN CGO_ENABLED=0 go run ./cmd/migration

FROM public.ecr.aws/lambda/provided:al2

COPY --from=builder /opt/app/agricoladb-server /agricoladb-server
COPY --from=builder /opt/app/agricoladb.sqlite /agricoladb.sqlite
ENTRYPOINT ["/agricoladb-server"]   
