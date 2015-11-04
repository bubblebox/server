DESTINATION = dist

SERVER_DIR = server
SERVER_BINARY = $(SERVER_DIR)/firedragon

CLIENT_DIR = client
CLIENT_DIST = $(CLIENT_DIR)/dist/*

$(DESTINATION): FORCE
	# Create directory structure
	mkdir -p $(DESTINATION)
	mkdir -p $(DESTINATION)/public

	# Compile and copy server binary
	$(MAKE) -C $(SERVER_DIR)
	cp $(SERVER_BINARY) $(DESTINATION)

	# Compile and copy client assets
	$(MAKE) -C $(CLIENT_DIR)
	cp -r $(CLIENT_DIST) $(DESTINATION)/public

setup:
	$(MAKE) -C $(SERVER_DIR) setup
	$(MAKE) -C $(CLIENT_DIR) setup

.PHONY: setup

test:
	$(MAKE) -C $(SERVER_DIR) test
	$(MAKE) -C $(CLIENT_DIR) test

.PHONY: test

clean:
	rm -rf $(DESTINATION)

.PHONY: clean

FORCE:
