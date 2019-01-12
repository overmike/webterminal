
VERSION := $(shell git describe --always --dirty --tag)
COMMIT := $(shell git describe --always --long)
REPO := overmike/webterminal

#EXTRA_LDFLAGS := -s -w
EXTRA_LDFLAGS := 
LDFLAGS := -X github.com/${REPO}/cmd.version=${VERSION} -X github.com/${REPO}/cmd.commit=${COMMIT} ${EXTRA_LDFLAGS}

webterminal:
	go build -v -ldflags="${LDFLAGS}"

clean:
	rm webterminal

proto_gen:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway

proto:
	protoc -I/usr/local/include -I./pb/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:terminal/ \
		--grpc-gateway_out=logtostderr=true:terminal/ \
		--swagger_out=logtostderr=true:terminal/ \
		pb/terminal.proto



web:
	yarn --cwd js
	yarn --cwd js build

rice:
	go get -u github.com/GeertJohan/go.rice/rice

asset: web
	rice -i github.com/${REPO}/cmd embed-go

docker:
	docker build --build-arg LDFLAGS="${LDFLAGS}" --target builder -t ${REPO}:builder .
	docker build --build-arg LDFLAGS="${LDFLAGS}" -t ${REPO} .
	docker tag ${REPO} ${REPO}:${VERSION}

.PHONY: proto_gen proto
