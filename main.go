package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kriticalflare/go-brainfuck-yourself/interpreter"
	"github.com/kriticalflare/go-brainfuck-yourself/lexer"
	"github.com/kriticalflare/go-brainfuck-yourself/parser"
)

func main() {
	fileContents, err := os.ReadFile("./testfiles/helloworld.b")
	// fileContents, err := os.ReadFile("./testfiles/krithik.b")
	// fileContents, err := os.ReadFile("./testfiles/bsort.b")
	// fileContents, err := os.ReadFile("./testfiles/test2.b")

	source := string(fileContents)

	if err != nil {
		log.Fatalf("failed to read source file: %v\n", err)
	}
	// fmt.Printf("source: %s \n", source)

	l := lexer.New(source)
	// fmt.Printf("%s \n", l.Tokens)
	p := parser.New(l)
	p.ParseProgram()
	if p.CheckParserErrors() {
		fmt.Printf("Error parsing the file\n")
		for _, err := range p.Errors {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
	}
	// fmt.Printf("Parser: \n")
	// fmt.Printf("%s\n", p.Tokens)
	i := interpreter.New(p.Tokens, 30_000)
	i.Run()
}
