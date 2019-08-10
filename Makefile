.PHONY: all

REPO= github.com/efortin/networklistener
IMAGE= kubi
TAG= 1.2.4
DOCKER_REPO= cagip



build:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -v -o ./build/networklistener -i $(GOPATH)/src/$(REPO)/cmd/main.go

test:
	 go test ./internal/services ./internal/types ./internal/utils

dep:
	glide install

