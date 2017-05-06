# Notes on developing Bursar

## What syntax is working right now?

- Booleans, ints, floats, strings, chars (runes)
- Equality and numeric comparisons (<, >, etc)
- Single and multiple assignments
- Nested expressions and correct parsing precedence
- If/else/elseif expressions
- Function definition
- Modules (single file)
- Comments

# What are the obvious missing pieces for an MVP?

- Importing other Bursar modules
- Importing Go packages? But we can run it through goimports
- Ability to define a main func? We could implement that anyway though
- Function calls!
- Calls to Go functions

### Function calls

Syntax: ```
let add = func a b {
    a + b;
}

add 5 6;
```

### Calls to Go functions

## What can we do with that?

## What next?

