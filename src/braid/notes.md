# Notes on developing Braid

## What's Braid's mission statement?
All the power of the Go platform (static binaries, concurrency, fast and fast GC) with
a higher-level syntax. Braid's higher-level abstractions may result in more code
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
- Calls to Braid functions

# What are the obvious missing pieces for an MVP?
- Importing other modules
- Importing Go packages? But we can run it through goimports
- Ability to define a main func? We could implement that anyway though

### Function calls

Syntax: 
```
let add = (a, b) {
    a + b;
}

add(5, 6);
Mod.add(5, 6);
```

### Type declarations

```
type Result ('a, 'b) = 
| OK 'a
| Error 'b

type MyPayload = {data: string}
type Person = {age: int, name: string}
type VaguePerson 'a = {name: string, extra: 'a}

type people = list person
type vaguePeople 'a = list vaguePerson 'a
type stringPeople string = list vaguePerson string
```

### Parsing
[X] Rename back to Braid
[X] Make new function definition rule that includes the let statement
[X] Array type literals
[X] Calls need to be compiled properly (add semicolons back in?)
[X] BinOpParens needs to be compiled with parentheses 
[X] 'Type' rule
    - [X] Record types
    - [X] Variant types
    - [X] Alias types
    - [X] Variant constructors need to support record constructor types
[X] Record type literals
[X] Variant type literals
[X] Function application should use parentheses 
[X] Type construction also
[X] Separate AST structs out so not so many multi-use types
[X] New `extern` rule
[X] `extern func` 
[X] `extern type` 
[X] `extern trait`
[X] Parse record field lookups eg `person.name`
[ ] Record and sum types need to handle `('a, 'b)` parentheses syntax
[X] `func` type
[ ] Make concrete types for func args if annotated but not used
[ ] Make annotation type that means extern pointer type `*`

### Compiling
[X] Hindley-Milner type inference, so we can predict errors and map function 
      params to types where needed
  [X] Add the inferred type to all Ast objects once inferred
  
  [X] Ifs as return types need to be unified
  [X] Ifs need Assignments added to assign the result to their temp var
  [X] Immutability means we cannot assign to a variable that already exists
  [X] Track used variables and do not compile (remove AST?) of assignments where var is not used,
      or change to `_`
  [X] Last expression in a func needs `return` AST inserted with correct variable name etc
  [X] Allow comments at the end of lines
  [X] Handle type annotations in func defns
  [X] Create stand-in Braid funcs for `extern func` imported funcs
  [X] Create stand-in types for `extern type` external records
  [ ] Create stand-in traits for `extern trait` external interfaces
  [X] Handle `package/package` paths in `extern` strings
  [X] Handle looking up complex non-base types in annotations
  [ ] Make sure external func calls are called with correct package names
  [X] Unify function call args with the function
  [X] Infer record types
  [ ] Infer variant types
  [ ] Make sure type variables get updated properly (prune function not working entirely?)
  [X] Make sure type variables get replaced properly (BinOps at least) or not compiled if not
  [ ] Ifs as expressions might need to be compiled to anonymous functions like so:
      `a := []string{"one","two", func() string{ if true { return "yes" } else { return "no" } }() }`
  [ ] `List thing` type implementation
  [X] If return is nil, omit it (Go can't return nil)
  [ ] External types in annotations, etc., need their package prefix
  [ ] Compile extern pointer type to `*Thing`
  [X] Make sure imports always come before everything else
  
[ ] Linking (Look up modules - Do they exist? Do functions mentioned exist?)
  

### Generating source
  [X] Functions need to be literals if defined inside a function (use State to change compilation behaviour)
  [ ] Generate concrete types etc (monomorphise) based on args when called
  [X] `main` needs to either not have a return type, or be renamed and wrapped in another `main`
  
### Later
[X] 'Module' rule
[X] Compile `let = if` rule specially - this means if expr branches need to be unified
    Kotlin has special let if expression form that's unified https://kotlinlang.org/docs/reference/control-flow.html 

[ ] `match` rule 
[X] Function annotations as a way of both typing a function and of specifying an external function?
[ ] Work out module signatures. Maybe like Elm: `module Main exposing (func1, func2)`
[ ] Work out typeclasses - Elm example https://medium.com/@eeue56/why-type-classes-arent-important-in-elm-yet-dd55be125c81
[ ] We need a way of defining function signatures. OCaml has interface files, Rust has inline types, 
    Haskell/Elm define on the line above. Can't leave it to H-M, need option of explicit typing. Annotations?
[ ] Exposed functions need to be uppercased
[X] Calls to external functions need to be uppercased
[ ] Look at standard typeclasses in Haskell, see which we could use

### Much later
[ ] Automated building, maybe fork something like gb
