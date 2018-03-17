GOCMD=gb
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=braid

all: build run

run: grammar build
	../../bin/braid $(ARG)
.PHONY: run

runv: grammar build
	../../bin/braid $(ARG) -v

build:
	gb build all
.PHONY: build

test: grammar build
	$(GOTEST) -v ./...
.PHONY: test

grammar:
	../../bin/pigeon ast/braid.peg | goimports > ast/grammar.go


