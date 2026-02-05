# Get a list of all directories in cmd/ (the app names)
APPS := $(shell ls cmd)

# The directory where binaries will be placed
BIN_DIR := bin

.PHONY: all build clean help

# Default target: build everything
all: build

## build: Compiles all apps into the /bin directory
build: clean
	@mkdir -p $(BIN_DIR)
	@echo "Building all apps..."
	@for app in $(APPS) ; do \
		echo "-> Compiling $$app"; \
		go build -o $(BIN_DIR)/$$app ./cmd/$$app; \
	done
	@echo "Done! Binaries are in $(BIN_DIR)/"

## clean: Removes the /bin directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

## help: Shows available commands
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
