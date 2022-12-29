DOCKER_COMPOSE_LOCAL := docker-compose -f docker-compose.local.yml
DOCKER_COMPOSE_TEST := docker-compose -f docker-compose.test.yml

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

.PHONY: docker-build
docker-build:
	$(DOCKER_COMPOSE_LOCAL) build

.PHONY: docker-start
docker-start:
	$(DOCKER_COMPOSE_LOCAL) up -d

.PHONY: docker-stop
docker-stop:
	$(DOCKER_COMPOSE_LOCAL) down

.PHONY: docker-restart-server
docker-restart-server:
	$(DOCKER_COMPOSE_LOCAL) restart server

.PHONY: docker-remove
docker-remove:
	$(DOCKER_COMPOSE_LOCAL) rm -v

.PHONY: docker-build-test
docker-build-test:
	$(DOCKER_COMPOSE_TEST) build

.PHONY: docker-start-test
docker-start-test:
	$(DOCKER_COMPOSE_TEST) up -d

.PHONY: docker-stop-test
docker-stop-test:
	$(DOCKER_COMPOSE_TEST) down

.PHONY: docker-remove-test
docker-remove-test:
	$(DOCKER_COMPOSE_TEST) rm -v
