.PHONY: clean
clean:
	$(RM) -r ./graph
	mkdir -p ./graph

.PHONY: generate
generate:
	CGO_ENABLED=0 go generate ./

.PHONY: build
build:
	go build ./cmd/server

.PHONY: build-migration
build-migration:
	go build ./cmd/migration

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: run-migration
run-migration:
	go run cmd/migration/main.go
