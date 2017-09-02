run: build
	../../bin/braid
.PHONY: run

build: grammar.go gb
.PHONY: build

grammar.go:
	../../bin/pigeon ast/braid.peg | goimports > ast/grammar.go

gb:
	gb build all
.PHONY: gb


