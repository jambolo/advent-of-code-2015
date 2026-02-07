# Get a list of all directories in cmd/ (the app names)
APPS := $(shell ls cmd)

# The directory where binaries will be placed
BIN_DIR := bin

# Common source files that all apps depend on (but also including tests)
COMMON_SRC := $(wildcard internal/load/*.go) \
			  $(wildcard internal/setup/*.go) \
			  $(wildcard internal/utils/*.go)

.PHONY: all build build-all clean help test

# Default target: build only changed apps
build: bin/day01 bin/day02 bin/day03 bin/day04 bin/day05 bin/day06 bin/day07 bin/day08 bin/day09 bin/day10 \
	   bin/day11 bin/day12 bin/day13 # bin/day14 bin/day15 bin/day16 bin/day17 bin/day18 bin/day19 bin/day20 \
	   bin/day21 bin/day22 bin/day23 bin/day24 bin/day25

## build-all: Compiles all apps into the /bin directory (forces rebuild)
build-all: clean
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

## test: Run all tests
test: $(COMMON_SRC)
	go test ./internal/...

# Pattern rule: build bin/dayNN from cmd/dayNN/
bin/%: cmd/%/*.go $(COMMON_SRC)
	go build -o $@ ./cmd/$*
