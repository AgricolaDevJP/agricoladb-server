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

.PHONY: ecr-login
ecr-login:
	aws --profile=AgricolaDevJP-admin ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 131551699606.dkr.ecr.ap-northeast-1.amazonaws.com

.PHONY: ecr-push-lambda
ecr-push-lambda:
	docker buildx build --platform linux/arm64/v8 -f ./docker/production/server-lambda/Dockerfile --push -t 131551699606.dkr.ecr.ap-northeast-1.amazonaws.com/agricoladb-server-lambda:latest .

.PHONY: deploy-lambda
deploy-lambda:
	(cd lambroll; AWS_PROFILE=AgricolaDevJP-admin lambroll --function ./function.jsonnet --region="ap-northeast-1" --tfstate "s3://agricoladevjp-terraform-states/AgricolaDevJP/agricoladb-server/terraform_states.tfstate" deploy)
