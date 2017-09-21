package main

import (
	"fmt"
	"strings"
	"braid/ast"
	"os"
	"io/ioutil"
)

func Compile(input string) (string, error) {
	lines := strings.Split(input, "\n")

	// first we make a reader from the input, which is a string
	r := strings.NewReader(input)
	// then we parse the input into ast
	result, err := ast.ParseReader("", r)

	if err != nil {

		fmt.Println("ERROR:")

		list := err.(ast.ErrorLister).Errors()
		for _, err := range list {
			// for each error, get the internal error
			pe := err.(ast.ParserError)
			printError(pe, lines)
		}
		return "", err
	} else {
		// print the input
		for i, el := range lines {
			fmt.Printf("%03d|%s\n", i+1, el)
		}

		// print the ast
		a := result.(ast.Ast)
		//fmt.Println("=", a.Print(0))

		env := make(ast.State)

		// infer types for the ast
		_, err := ast.Infer(a.(ast.Module), &env, nil)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}

		fmt.Println(env)

		// print the compiled Go
		result := a.Compile(env)
		fmt.Println(result)
		return result, nil
	}

}


func printError(pe ast.ParserError, lines []string) {

	// determine how many past lines to render
	start := pe.Pos()[0] - 1
	if pe.Pos()[0] >= 5 {
		start = pe.Pos()[0] - 5
	}

	// print those past lines up until the line of the error
	for i, el := range lines[start:pe.Pos()[0]] {
		offset := start
		fmt.Printf("%03d|%s\n", i+1+offset, el)
		//i += 1
	}

	// print the caret pointing to the position
	line := lines[pe.Pos()[0]-1]
	fmt.Printf("    ")
	for _, el := range line[:pe.Pos()[1]-1] {
		if el == '\t' {
			fmt.Printf("----")
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Printf("^\n\n")

	// print the actual error
	fmt.Printf("Line %d, character %d: ", pe.Pos()[0], pe.Pos()[1])
	fmt.Println(pe.InnerError())
}


func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Must supply an argument of a file to compile, eg. `$ braid example.bd`\n")
		return
	}

	result, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", args[0], err.Error())
		return
	}

	file := string(result)
	compiled, err := Compile(file)
	if err != nil {
		return
	}
	fmt.Print(compiled)

}
