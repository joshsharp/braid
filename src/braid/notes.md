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

### Type declarations

```
type result 'a 'b = 
| OK 'a
| Error 'b

type myPayload = {data: string}
type person = {age: int, name: string}
type vaguePerson 'a = {name: string, extra: 'a}

type people = list person
type vaguePeople 'a = list vaguePerson 'a
type stringPeople string = list vaguePerson string
```

### Calls to Go functions

Perhaps for MVP all calls should just map to Go funcs? That'd be easiest.
Also need to work out exporting/header files/signatures and how this will map to Go.

## What next?

[X] Rename back to Braid
[X] Make new function definition rule that includes the let statement
[X] Array type literals
[X] Calls need to be compiled properly (add semicolons back in?)
[X] BinOpParens needs to be compiled with parentheses 
[X] 'Type' rule
    - [X] Record types
    - [X] Variant types
    - [X] Alias types
[ ] Record type literals
[ ] Variant type literals
[ ] 'Module' rule
[ ] Compile `let = if` rule specially - this means if expr branches need to be unified
    Kotlin has special let if expression form that's unified https://kotlinlang.org/docs/reference/control-flow.html 
[ ] Ifs as expressions might need to be compiled to anonymous functions like so:
    `a := []string{"one","two", func() string{ if true { return "yes" } else { return "no" } }() };`
[ ] `match` rule
[ ] `let a = b` compiles to `b()`, use state to look up if `b` is function and if not, no parentheses 
[ ] Work out module signatures. Maybe like Elm: `module Main exposing (func1, func2)`
[ ] Work out typeclasses - Elm example https://medium.com/@eeue56/why-type-classes-arent-important-in-elm-yet-dd55be125c81
[ ] Exposed functions need to be uppercased
[ ] Calls to external functions need to be uppercased
[ ] Look at standard typeclasses in Haskell, see which we could use

Compiling currently maps straight to outputting code text, needs more passes:
- [ ] Hindley-Milner type inference, so we can predict errors and map function 
      params to types where needed
- [ ] Linking (Do functions mentioned exist? Do modules?)
- [ ] Listing and generating of required concretely-typed generic functions
- Then generating source  
 
[ ] Automated building, maybe fork something like gb
