SERVER_NAME = ConfigPlatform
GOCMD = go
GOBUILD = $(GOCMD) build -x

.PHONY: all build clean
all: build
build:
	export GO111MODULE=on
	$(GOBUILD) -o bin/$(SERVER_NAME)
clean:
	rm -f bin/$(SERVER_NAME)
