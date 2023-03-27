GO               = go
M                = $(shell printf "\033[34;1m>>\033[0m")
GOBIN			 ?= $(PWD)/bin
TARGET_DIR       ?= $(PWD)/.build

.PHONY: start-server
start-server:
	$(GO) run ./cmd/server/*.go

.PHONY: start-client
start-client:
	$(GO) run ./cmd/client/*.go

.PHONY: build-server
build-server: ## Build 'server' binary
	$(info $(M) building server...)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build --ldflags "-linkmode external" -o $(TARGET_DIR)/server ./cmd/server/*.go

.PHONY: build-client
build-client: ## Build 'client' binary
	$(info $(M) building client...)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build --ldflags "-linkmode external" -o $(TARGET_DIR)/client ./cmd/client/*.go


help: ##Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
