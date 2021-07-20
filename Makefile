.PHONY: full
full: build_docker

.PHONY: build_docker
build_docker:
	@./scripts/build.sh

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := full

