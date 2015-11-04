DESTINATION = dist

EMBER_BUILD_CMD = ember build
EMBER_BUILD_OPTIONS = -prod

EMBER_TEST_CMD = ember test
EMBER_TEST_OPTIONS =

$(DESTINATION): FORCE
	$(EMBER_BUILD_CMD) $(EMBER_BUILD_OPTIONS) --output-path $(DESTINATION)/

setup:
	npm install
	bower install

.PHONY: setup

test: FORCE
	$(EMBER_TEST_CMD) $(EMBER_TEST_OPTIONS)

.PHONY: test

clean:
	rm -rf $(DESTINATION)

.PHONY: clean

FORCE:
