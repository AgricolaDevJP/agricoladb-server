.PHONY: clean
clean:
	rm -rf ./graph
	mkdir -p ./graph

.PHONY: generate
generate:
	go generate ./