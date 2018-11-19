
VERSION := $(shell git describe --always --long --dirty --tag)

webterminal:
	go build -v -ldflags="-X github.com/overmike/webterminal/cmd.version=${VERSION}"

clean:
	rm webterminal

proto_gen:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

proto:
	protoc -I/usr/local/include -I./pb/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:terminal/ \
		--grpc-gateway_out=logtostderr=true:terminal/ \
		--swagger_out=logtostderr=true:terminal/ \
		pb/terminal.proto

vendor:
	dep ensure -v


web:
	yarn --cwd js
	yarn --cwd js build

rice:
	go get -u github.com/GeertJohan/go.rice/rice

asset: web
	rice -i github.com/overmike/webterminal/cmd embed-go

docker:
	docker build -t overmike/webterminal .

.PHONY: proto_gen proto vendor
