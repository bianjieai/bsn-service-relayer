#!/usr/bin/make -f

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building relayer..."
	@go build -mod=readonly -o build/ relayer/cmd/relayer
else
	@echo "building relayer..."
	@go build -mod=readonly -o build/ relayer/cmd/relayer
endif

install: go.sum
	@echo "installing relayer..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/relayer-bcos relayer/cmd/relayer

test-eth:
	@echo "sending the interchain request to ethereum for testing..."
	@go test -mod=readonly -count=1 relayer/tests/ethereum
	@echo "sent successfully"
