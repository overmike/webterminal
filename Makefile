
webterminal: proto vendor
	go build

proto:
	protoc -I/usr/local/include -I./pb/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:terminal/ \
		--grpc-gateway_out=logtostderr=true:terminal/ \
		pb/terminal.proto
vendor:
	dep ensure -v

.PHONY: proto vendor
