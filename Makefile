protoc-gen-validate=vendor/github.com/envoyproxy/protoc-gen-validate

all: build

env:
	go mod vendor
	git submodule init
	git submodule update

build:
	go build

proto:
	protoc \
		--proto_path=. \
		--proto_path=/usr/local/include \
		--proto_path=googleapis \
		--proto_path=${protoc-gen-validate} \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--validate_out="lang=go:." \
		definition/*.proto
