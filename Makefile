DOCKER_COMPOSE_LOCAL := docker compose -f docker-compose.local.yml

.PHONY: docker-build
docker-build:
	$(DOCKER_COMPOSE_LOCAL) build

.PHONY: docker-start
docker-start:
	$(DOCKER_COMPOSE_LOCAL) up -d

.PHONY: docker-stop
docker-stop:
	$(DOCKER_COMPOSE_LOCAL) down

.PHONY: docker-ps
docker-ps:
	$(DOCKER_COMPOSE_LOCAL) ps

.PHONY: docker-restart-server
docker-restart-server:
	$(DOCKER_COMPOSE_LOCAL) restart server

.PHONY: docker-remove
docker-remove:
	$(DOCKER_COMPOSE_LOCAL) rm -v
