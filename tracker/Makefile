.DEFAULT_GOAL := all
VERSION := 0.1
NAME := tracker
PROTOC := protoc -I . ./tracker.proto --go_out=plugins=grpc:.

.PHONY: protoc
protoc:
	$(PROTOC)

.PHONY: mockery
mockery:
	mockery -all

build:
	pushd server; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../bin/tracker-server .; popd;

docker:
	docker build -t darren-west/$(NAME):$(VERSION) .

.PHONY: all
all: protoc mockery build docker