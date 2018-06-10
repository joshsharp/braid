# Braid
A functional language with Reason-like syntax that compiles to Go.

I’m working on a language I’m calling Braid, an ML-like language that compiles to Go. Braid’s syntax is heavily inspired by [Reason](https://reasonml.github.io/), itself a more C-like syntax on top of OCaml. So really I’m writing a language that aims to be fairly similar to OCaml in what it can do, but visually a bit closer to Go. I’m not trying to reimplement OCaml or Reason 1:1 on top of Go, but build something sharing many of the same concepts.

I've written some more about it on my [Braid dev blog](https://braid.joshsharp.com.au/).

## Status

Very, very alpha.

## Goals
- Pair an OCaml-like language with the benefits of the Go platform (speed, concurrency, static binaries, a healthy ecosystem)
- Bring powerful FP concepts to Go
- Get around Go's lack of generics
- Interop with Go code
- Ability to use Go stdlib

## Non-goals
- Performance matching idiomatic Go
- Just reimplementing Reason on top of Go

## Language overview

Consider anything ticked off to exist in the language, but be barely usable.

- [X] Record types
- [X] Variant types
- [X] If-expressions
- [X] Importing Go functions and types
- [X] Immutability by default
- [X] Hindley-Milner type inference
- [X] Type annotations
- [X] Implicit return
- [ ] Modules
- [ ] Pattern matching
- [ ] Currying
- [ ] Typeclasses/traits
- [ ] Concurrency
- [ ] Infix operators

Braid supports records and variants:

```
type Person = {
  name: string,
  age: int64,
}

type Fruit = 
  | Peach
  | Plum
  | Pear

type Option ('a) =
  | Some ('a)
  | None
  
let result = Option("it worked")
```

Braid attempts to support significant newlines, meaning no `;` required &mdash; however this is probably broken in a lot of cases right now.

A full example:

```
module Main

// record type
type Payload = {
  name: string,
  data: string,
}

// go interop - external functions must be annotated
extern func println = "fmt.Println" (s: string) -> ()
extern func printf1 = "fmt.Printf" (s: string, arg1:string) -> ()

/* func to add cheesiness to any two items */
let cheesy = (item, item2) {
  item ++ " and " ++ item2 ++ " with cheese please"
}

let main = {
  // nested functions
  let something = {
    4 + 9
  }
  let a = something()
  let yumPizza = cheesy("pineapple", "bbq sauce")
  println(yumPizza)
  // calling a go function
  printf1("Woo I can print %s\n", "6")
  let b = Payload{name: "greeting", data: "hi"}
  println(b.name)
}
```

## Trying it out

Grab the correct Braid package for your platform from the [releases](https://github.com/joshsharp/braid/releases), extract the `braid` binary, and run it.

```sh
./braid filename.bd
```

This will compile your Braid file to Go and print the resulting Go source code to stdout.

```sh
./braid filename.bd > main.go
```

You can redirect this into a file if you like.

## Developing

### Requirements
- [Go 1.10](https://golang.org/dl/)
- [GB](https://getgb.io/)
- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)

### Building

Making sure Go and GB are in your path, clone the Braid repository into a new directory:

```sh
git clone https://github.com/joshsharp/braid.git
```

Enter the new `braid` directory and fetch the requirements:

```sh
cd braid
gb vendor restore
```

Make sure the vendored dependencies are built (you'll only need to do this once):

```sh
cd vendor
gb build all
```

Use the makefile at `src/braid/Makefile` to build and run Braid:

```sh
cd ../src/braid
make run file=examples/example.bd
```

## FAQ
### Will Braid support X?

I don't know yet. I'm open to proposals, provided you help me do the work.

### Do you even know what you're doing?

Nope, not at all. I have no formal background in this stuff. Really I'm doing it for fun. I'd love to see it reach maturity, because I want to use it myself. But I'll need a lot of help if it's to get that far.

## Contributing
### Contribution guidelines
Your help makes Braid better! I welcome pull requests, bug fixes, and issue reports.

Before proposing a change, please first create an issue to discuss your proposal.

## License

Licensed under the [MIT License](https://choosealicense.com/licenses/mit/).
