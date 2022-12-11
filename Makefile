.PHONY: clean
clean:
	$(RM) -r ./graph
	mkdir -p ./graph

.PHONY: generate
generate:
	go generate ./

.PHONY: run
run:
	go run server.go

.PHONY: docker-build
docker-build:
	docker-compose build

.PHONY: docker-start
docker-start:
	docker-compose up -d

.PHONY: docker-stop
docker-stop:
	docker-compose down

.PHONY: docker-restart-server
docker-restart-server:
	docker-compose restart server
