GOCMD=gb
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=braid

all: build run

run: grammar build
	../../bin/braid $(file)
.PHONY: run

runv: grammar build
	../../bin/braid $(file) -v

build:
	gb build all
.PHONY: build

test: grammar build
	$(GOTEST) -v ./...
.PHONY: test

grammar:
	../../vendor/bin/pigeon ast/braid.peg | goimports > ast/grammar.go



