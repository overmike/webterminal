
VERSION := $(shell git describe --always --long --dirty --tag)

#EXTRA_LDFLAGS := -s -w
EXTRA_LDFLAGS := 
LDFLAGS := -X github.com/overmike/webterminal/cmd.version=${VERSION} ${EXTRA_LDFLAGS}

webterminal: vendor
	go build -v -ldflags="${LDFLAGS}"

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

dep:
	go get -u github.com/golang/dep/cmd/dep

vendor:
	@echo "Assume you have dep installed, if not please run 'make dep'"
	dep ensure -v --vendor-only


web:
	yarn --cwd js
	yarn --cwd js build

rice:
	go get -u github.com/GeertJohan/go.rice/rice

asset: web
	rice -i github.com/overmike/webterminal/cmd embed-go

docker:
	docker build --build-arg LDFLAGS="${LDFLAGS}" -t overmike/webterminal .

.PHONY: proto_gen proto vendor
