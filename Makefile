PROTO_DIR := proto
GEN_DIR := gen/go

generate:
	@mkdir -p ${GEN_DIR}/auth
	@protoc -I ${PROTO_DIR}/ ${PROTO_DIR}/auth/auth.proto \
		--go_out=${GEN_DIR} --go_opt=paths=source_relative \
		--go-grpc_out=${GEN_DIR} --go-grpc_opt=paths=source_relative
	@echo "Код успешно сгенерирован в ${GEN_DIR}"
