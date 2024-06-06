# Makefile

# Specify the Go binary name
BINARY_NAME=loadbalance_ebpf


# Generated files
GEN_FILES=reuseportlb_*
# Default server ID (will be overridden)
SERVER_ID=

# Go generate command
GO_GENERATE_CMD=go generate

# Go build command
GO_BUILD_CMD=go build

# Default target
all: generate build

# Generate eBPF code
generate:
	$(GO_GENERATE_CMD)

# Build the Go application
build:
	$(GO_BUILD_CMD) -o $(BINARY_NAME)

# Launch the server with the specified server ID
launch: check-server-id
	sudo ./$(BINARY_NAME) $(SERVER_ID)

# Check if SERVER_ID is set
check-server-id:
ifndef SERVER_ID
	$(error SERVER_ID is not set. Use 'make launch SERVER_ID=<id>')
endif

# Clean up generated files
clean:
	sudo rm -f $(BINARY_NAME)
	sudo rm -f $(GEN_FILES)

# Usage help
help:
	@echo "Usage:"
	@echo "  make [TARGET] SERVER_ID=<id>"
	@echo "Targets:"
	@echo "  all        - Generate and build the application (default)"
	@echo "  generate   - Run go generate"
	@echo "  build      - Build the Go application"
	@echo "  launch        - Run the server with SERVER_ID=<id>"
	@echo "  clean      - Clean up generated files"
	@echo "  help       - Show this help message"

.PHONY: all generate build run check-server-id clean help