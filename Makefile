install-protoc-gen-sample:
	go install ./cmd/protoc-gen-sample/

proto/gen: install-protoc-gen-sample
	prototool generate
