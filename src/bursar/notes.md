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
- Calls to Go functions

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

Perhaps for MVP all calls should just map to Go funcs? That'd be easiest

### Incorrect position for parser errors

The line in `parser.parse` for adding an error needs to look like this:

```go
// make sure this doesn't go out silently
p.addErrAt(errNoMatch, p.cur.pos)
```

Maybe do a pull request to pigeon to get this fixed?

## What can we do with that?

## What next?

