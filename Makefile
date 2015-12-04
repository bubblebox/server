DESTINATION = dist

SERVER_DIR = server
SERVER_BINARY = $(SERVER_DIR)/firedragon

CLIENT_DIR = client
CLIENT_DIST = $(CLIENT_DIR)/dist/*

$(DESTINATION): FORCE
	# Create directory structure
	mkdir -p $(DESTINATION)

	# Build client (Ember.js)
	$(MAKE) -C $(CLIENT_DIR)

	# Move HTML to server/client_dist
	cp -r $(CLIENT_DIST) $(SERVER_DIR)/client_dist

	# Compile server binary
	$(MAKE) -C $(SERVER_DIR)

	cp $(SERVER_BINARY) $(DESTINATION)

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
	$(MAKE) -C $(SERVER_DIR) clean
	$(MAKE) -C $(CLIENT_DIR) clean

.PHONY: clean

FORCE:
