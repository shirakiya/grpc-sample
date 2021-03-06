bash:
	docker-compose run --rm go /bin/bash

run: proto/gen
	go run cmd/grpc_sample_main/main.go

install:
	go install ./...

install-protoc-gen-sample:
	go install ./cmd/protoc-gen-sample/

proto/gen: install-protoc-gen-sample
	buf generate
