OWNER := rnymphaea
PROJECT := chronoflow

PROTO_DIR := proto
GEN_DIR := gen/go
SERVICES := auth
LINTER := golangci-lint

VERSION ?= v1

MODE ?= local

LOCAL_DEPLOY_DIR := deploy/local

ifeq ($(MODE), local)
	BUILD = local_build 
	RUN = local_up
	DOWN = local_down
endif

.PHONY: all
all: $(BUILD) $(RUN)

.PHONY:
down: $(DOWN)

.PHONY: local_build
local_build:
	docker compose -f $(LOCAL_DEPLOY_DIR)/docker-compose.yml build

.PHONY: local_up
local_up:
	docker compose -f $(LOCAL_DEPLOY_DIR)/docker-compose.yml up -d

.PHONY: local_down
local_down:
	docker compose -f $(LOCAL_DEPLOY_DIR)/docker-compose.yml down

.PHONY: test 
test:
	@for service in $(SERVICES); do \
		echo "Running tests for $$service..."; \
		cd $$service && go test ./... -v; \
		cd ..; \
	done

.PHONY: prepare 
prepare:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@export PATH="$PATH:$(go env GOPATH)/bin"

.PHONY: lint
lint: prepare
	@for service in $(SERVICES); do \
		echo "Running linters for $$service..."; \
		cd $$service && $(LINTER) run ./...; \
		cd ..; \
	done

generate: prepare
	@for service in $(SERVICES); do \
		mkdir -p ${GEN_DIR}/$$service/$(VERSION); \
		protoc -I ${PROTO_DIR}/ ${PROTO_DIR}/$$service/$(VERSION)/*.proto \
			--go_out=${GEN_DIR} --go_opt=paths=source_relative \
			--go-grpc_out=${GEN_DIR} --go-grpc_opt=paths=source_relative; \
	done
	@cd $(GEN_DIR) && go mod init github.com/$(OWNER)/$(PROJECT)/$(GEN_DIR)
	@cd $(GEN_DIR) && go mod tidy
	@echo "Код успешно сгенерирован в ${GEN_DIR}"

.PHONY: clean
clean:
	@rm -rf $(GEN_DIR)
	@echo "Удалено: $(GEN_DIR)"
