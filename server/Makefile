DESTINATION = firedragon

GO_BUILD_CMD = go build
GO_LDFLAGS =
GO_TEST_CMD = go test -v ./...
GO_FILES := $(shell find . -name '*.go')

$(DESTINATION): $(GO_FILES)
	go-bindata-assetfs client_dist/...
	$(GO_BUILD_CMD) $(GO_LDFLAGS) -o $(DESTINATION) *.go
	chmod +x $(DESTINATION)

setup:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...
	mkdir -p client_dist
	echo "Ember.js not enabled in debug build." > client_dist/index.html
.PHONY: setup

test: FORCE
	go-bindata-assetfs client_dist/...
	$(GO_TEST_CMD)
.PHONY: test

clean:
	rm -f $(DESTINATION)
	mkdir -p client_dist
	rm -rf client_dist/*
	echo "Ember.js not enabled in debug build." > client_dist/index.html
.PHONY: clean

FORCE:
