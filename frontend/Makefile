all:proto
PRODADDR=joshcarp-it-project-ogaheemccq-uc.a.run.app
PORT=443
INCLUDE=-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
.PHONY: proto proto-docker

proto:              ## Remake the proto generation
	docker run --rm -v $$(pwd):/itproject:rw joshcarp/protoc $(INCLUDE) -I./itproject/proto/ --go_out=paths=source_relative:/itproject/backend/pkg/proto/itproject api.proto
	docker run --rm -v $$(pwd):/itproject:rw joshcarp/protoc $(INCLUDE) -I././itproject/proto/ --go-grpc_out=paths=source_relative:/itproject/backend/pkg/proto/itproject api.proto
	docker run --rm -v $$(pwd):/itproject:rw joshcarp/protoc $(INCLUDE) -I././itproject/proto/ --js_out=import_style=commonjs:/itproject/frontend/src/proto api.proto
	docker run --rm -v $$(pwd):/itproject:rw joshcarp/protoc $(INCLUDE) -I././itproject/proto/ --grpc-web_out=import_style=commonjs+dts,,mode=grpcwebtext:/itproject/frontend/src/proto api.proto

fontend:
	npm install && npm start
