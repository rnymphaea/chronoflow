PROTO_DIR := proto
GEN_DIR := gen/go

MODE ?= local

LOCAL_DEPLOY_DIR := deploy/local

ifeq ($(MODE), local)
	BUILD = local_build 
	RUN = local_up
	STOP = local_stop
endif

.PHONY: all
all: $(BUILD) $(RUN)

.PHONY:
stop: $(STOP)

.PHONY: local_build
local_build:
	docker compose -f $(LOCAL_DEPLOY_DIR)/docker-compose.yml build

.PHONY: local_up
local_up:
	docker compose -f $(LOCAL_DEPLOY_DIR)/docker-compose.yml up -d

.PHONY: local_stop
local_stop:
	docker compose -f $(LOCAL_DEPLOY_DIR)/docker-compose.yml down

generate:
	@mkdir -p ${GEN_DIR}/auth
	@protoc -I ${PROTO_DIR}/ ${PROTO_DIR}/auth/auth.proto \
		--go_out=${GEN_DIR} --go_opt=paths=source_relative \
		--go-grpc_out=${GEN_DIR} --go-grpc_opt=paths=source_relative
	@echo "Код успешно сгенерирован в ${GEN_DIR}"
