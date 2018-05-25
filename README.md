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

## Installation

Instructions coming soon.

## FAQ
### Do you even know what you're doing?

Nope, not at all. I have no formal background in this stuff. Really I'm doing it for fun. I'd love to see it reach maturity, because I want to use it myself. But I'll need a lot of help if it's to get that far.
