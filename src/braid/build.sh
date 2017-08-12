../../bin/pigeon ast/braid.peg | goimports > ast/grammar.go
gb build all
../../bin/braid


