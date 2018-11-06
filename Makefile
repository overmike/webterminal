
webterminal: proto vendor
	go build

proto:
	protoc -I ./pb/ --go_out=plugins=grpc:terminal/ pb/terminal.proto

vendor:
	dep ensure -v

.PHONY: proto vendor
