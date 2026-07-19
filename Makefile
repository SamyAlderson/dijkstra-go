# Build script for dijkstra-go project

# Project-specific variables
NAME := dijkstra-go

# Target directories
BINDIR := bin
SRCDIR := src
TESTDIR := test

# Dependencies
PKG_ERRORS := github.com/pkg/errors

# Build rules
.PHONY: all
all: $(BINDIR)/dijkstra-go

$(BINDIR)/dijkstra-go: $(BINDIR)/main $(BINDIR)/dijkstra $(BINDIR)/graph
	go build -o $(BINDIR)/dijkstra-go $(SRCDIR)/main.go

$(BINDIR)/main $(BINDIR)/dijkstra $(BINDIR)/graph:
	go build -o $(BINDIR)/$@ $<

# Test rules
.PHONY: test
test: $(TESTDIR)/test_dijkstra $(TESTDIR)/test_graph
	go test -v $(TESTDIR)/test_dijkstra.go
	go test -v $(TESTDIR)/test_graph.go

$(TESTDIR)/test_dijkstra $(TESTDIR)/test_graph:
	go build -o $(TESTDIR)/$@ $<

# Clean rule
.PHONY: clean
clean:
	rm -f $(BINDIR)/* $(TESTDIR)/*
	rm -rf $(BINDIR) $(TESTDIR)

# Go mod management
.PHONY: mod
mod:
	go mod tidy
	go mod verify

# Print help message
.PHONY: help
help:
	@echo "Makefile for dijkstra-go project"
	@echo "Targets:"
	@echo "  all        Build the project"
	@echo "  test       Run unit tests"
	@echo "  clean      Clean up generated files"
	@echo "  mod        Manage Go modules"

# Default target
.DEFAULT_GOAL := all