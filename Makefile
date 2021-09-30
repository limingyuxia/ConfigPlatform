SERVER_NAME = ConfigServer
GOCMD=go
GOBUILD=$(GOCMD) build -x

all: build
build:
	export GO111MODULE=on
	$(GOBUILD) -o bin/$(SERVER_NAME)
clean:
	rm -f bin/$(SERVER_NAME)
