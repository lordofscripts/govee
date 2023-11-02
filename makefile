GO=~/go/bin/colorgo
BIN=~/go/bin
BIN_OUT=$(BIN)/govee
CMD_CLI=./cmd

all:	govee

govee:
	$(GO) build -v -o $(BIN_OUT) -tags=none $(CMD_CLI)/*.go

update:
	go get -u all

testall:
	go test ./...

testfull:
	go test -v test/*_test.go
