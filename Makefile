
webterminal: vendor
	go build

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

.PHONY: proto_gen proto vendor
