.DEFAULT_GOAL := all

PROTOC := protoc -I . ./tracker.proto --go_out=plugins=grpc:.

.PHONY: protoc
protoc:
	$(PROTOC)

.PHONY: mockery
mockery:
	mockery -all

.PHONY: all
all: protoc mockery