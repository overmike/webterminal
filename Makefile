
webterminal: proto
	go build

proto:
	protoc -I ./pb/ --go_out=plugins=grpc:terminal/ pb/terminal.proto

.PHONY: proto
