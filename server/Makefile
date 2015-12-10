DESTINATION = firedragon

GO_BUILD_CMD = go build
GO_LDFLAGS =
GO_TEST_CMD = godep go test -v ./...
GO_FILES := $(shell find . -name '*.go')

$(DESTINATION): $(GO_FILES)
	$(GO_BUILD_CMD) $(GO_LDFLAGS) -o $(DESTINATION) *.go
	chmod +x $(DESTINATION)

setup:
	go get github.com/tools/godep
.PHONY: setup

test: FORCE
	$(GO_TEST_CMD)
.PHONY: test

clean:
	rm -f $(DESTINATION)
.PHONY: clean

FORCE:
