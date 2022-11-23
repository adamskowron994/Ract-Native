all:proto
PRODADDR=joshcarp-grpc-vs-connect-ogaheemccq-uc.a.run.app
PORT=443
.PHONY: proto proto-docker


proto:              ## Remake the proto generation
	docker run --rm -v $$(pwd):/grpc-vs-connect:rw joshcarp/protoc -I./grpc-vs-connect/proto/ --go_out=paths=source_relative:/grpc-vs-connect/backend/pkg/proto/elizav1 eliza.proto
	docker run --rm -v $$(pwd):/grpc-vs-connect:rw joshcarp/protoc -I././grpc-vs-connect/proto/ --go-grpc_out=paths=source_relative:/grpc-vs-connect/backend/pkg/proto/elizav1 eliza.proto
	docker run --rm -v $$(pwd):/grpc-vs-connect:rw joshcarp/protoc -I././grpc-vs-connect/proto/ --js_out=import_style=commonjs:/grpc-vs-connect/frontend/src/proto eliza.proto
	docker run --rm -v $$(pwd):/grpc-vs-connect:rw joshcarp/protoc -I././grpc-vs-connect/proto/ --grpc-web_out=import_style=commonjs,,mode=grpcwebtext:/grpc-vs-connect/frontend/src/proto eliza.proto

docker:             ## Build the authentication service
	docker build .  --build-arg SERVICE=$(SERVICE) -t joshcarp/grpc-vs-connect

run:                ## Run docker
	docker run --rm -p 443:443 joshcarp/grpc-vs-connect

ping:               ## Ping the authentication service
	docker run --rm joshcarp/grpcurl -d '{"email": "Hello", "password": "123", "userid": "123" }' --plaintext host.docker.internal:$(PORT) grpc-vs-connect.authenticate/Register

client:             ## Make the demo client
	docker build . -f build/client.Dockerfile -t joshcarp/grpc-vs-connect-client

client.run:         ## Run the demo client in docker
	docker run --rm -e ADDR=host.docker.internal:443 joshcarp/grpc-vs-connect-client

secret:             ## Remake a jwt secret
	openssl rand -hex 64  | pbcopy

docker-compose:     ## Run all the services in build/docker-compose.yaml
	docker-compose -f docker-compose.yaml up

help:               ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: fontend
fontend:
	cd frontend && npm install && npm start

website:
	rm -rf docs/* && echo "docs.epicportfol.io" > docs/CNAME && cd hugo && hugo -t reveal-hugo && mv  public/* ../docs && cd .. && make docs -B
