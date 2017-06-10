# Notes on developing Bursar

## What's Bursar's mission statement?
All the power of the Go platform (static binaries, concurrency, fast and fast GC) with
a higher-level syntax. Bursar's higher-level abstractions may result in more code
and slower speeds than the representative native Go code, but this tradeoff is explicitly made.
 
## What sort of stuff can you do in Bursar that you can't in Go?
- Mostly ML-like syntax
- Generic functions (H-M type checking)
- Immutability by default
- Pattern matching and algebraic data types
- Typeclasses/traits/interfaces
- Should we support currying?

## What syntax is working right now?
- Booleans, ints, floats, strings, chars (runes)
- Equality and numeric comparisons (<, >, etc)
- Single and multiple assignments
- Nested expressions and correct parsing precedence
- If/else/elseif expressions
- Function definition
- Modules (single file)
- Comments
- Calls to Go functions
- Calls to Bursar functions

# What are the obvious missing pieces for an MVP?
- Importing other Bursar modules
- Importing Go packages? But we can run it through goimports
- Ability to define a main func? We could implement that anyway though

### Function calls

Syntax: 
```
let add = func a b {
    a + b;
}

add 5 6;
Mod.add 5 6;
```

### Calls to Go functions

Perhaps for MVP all calls should just map to Go funcs? That'd be easiest.
Also need to work out exporting/header files/signatures and how this will map to Go.

## What next?

[X] Rename back to Braid
[X] Make new function definition rule that includes the let statement
[ ] 'Type' rule
    - [ ] Record types
    - [ ] Variant types
[ ] Compile `let = if` rule specially
[ ] `let a = b` compiles to `b()`, use state to look up if `b` is function and if not, no parentheses 
[ ] Work out module signatures - maybe like https://realworldocaml.org/v1/en/html/files-modules-and-programs.html#nested-modules
[ ] Work out typeclasses - Elm example https://medium.com/@eeue56/why-type-classes-arent-important-in-elm-yet-dd55be125c81
[ ] 'Module' rule

Compiling currently maps straight to outputting code text, needs more passes:
- [ ] Hindley-Milner type inference, so we can predict errors and map function 
      params to types where needed
- [ ] Linking (Do functions mentioned exist? Do modules?)
- [ ] Listing and generating of required concretely-typed generic functions
- Then generating source  

