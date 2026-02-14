BINARY   := bwm
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS  := -ldflags "-s -w -X github.com/Softorize/bwm/cmd.Version=$(VERSION)"

.PHONY: build clean install

build:
	go build $(LDFLAGS) -o $(BINARY) .

clean:
	rm -f $(BINARY)

install:
	go install $(LDFLAGS) .
