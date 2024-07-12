
# Set Compiler to either of: go, gocolor OR gopretty
MODE=gopretty

# Compilers
GO=go
GOCOLOR=~/go/bin/colorgo
GOPRETTY="$(HOME)/go/bin/gofilter"
# Sources
CMD_CLI=./cmd
# Outputs
BIN=~/go/bin
BIN_OUT=$(BIN)/goveelux

ifeq ($(MODE),gocolor)
        GO=$(GOCOLOR)
endif


all:	govee

govee:
ifeq ($(MODE),gopretty)
	$(GO) build -v -o $(BIN_OUT) -tags=none $(CMD_CLI)/*.go 2>&1 | $(GOPRETTY) -color -width 75 -version
else
	$(GO) build -v -o $(BIN_OUT) -tags=none $(CMD_CLI)/*.go
endif

update:
	go get -u all

testall:
	go test ./...

testfull:
	go test -v test/*_test.go
