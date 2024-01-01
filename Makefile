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

.PHONY: ecr-login
ecr-login:
	aws --profile=AgricolaDevJP-admin ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 131551699606.dkr.ecr.ap-northeast-1.amazonaws.com

.PHONY: ecr-push-lambda
ecr-push-lambda:
	docker buildx build --platform linux/arm64/v8 -f ./docker/production/server-lambda/Dockerfile --push -t 131551699606.dkr.ecr.ap-northeast-1.amazonaws.com/agricoladb-server-lambda:latest .

.PHONY: deploy-lambda
deploy-lambda:
	(cd lambroll; AWS_PROFILE=AgricolaDevJP-admin lambroll --function ./function.jsonnet --region="ap-northeast-1" --tfstate "s3://agricoladevjp-terraform-states/AgricolaDevJP/agricoladb-server/terraform_states.tfstate" deploy)
