DESTINATION = firedragon

GO_BUILD_CMD = go build
GO_LDFLAGS =
GO_TEST_CMD = go test -v ./...

$(DESTINATION): *.go
	$(GO_BUILD_CMD) $(GO_LDFLAGS) -o $(DESTINATION) server.go

setup:
	echo "Nothing to do"

.PHONY: setup

test: FORCE
	$(GO_TEST_CMD)

.PHONY: test

clean:
	rm -f $(DESTINATION)

.PHONY: clean

FORCE:
