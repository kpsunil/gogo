package main

import (
	"log"
	"os"

	"asgn3/src/gentoken"
	"asgn3/src/parser"
)

// GenToken generates the tokens returned by lexer for the input program.
func GenToken(file string) {
	gentoken.PrintTokens(file)
}

// GenRightmostDerivation generates the rightmost derivations used in the bottom-up
// parsing and pretty-prints them in an HTML format.
func GenRightmostDerivation(file string) {
	parser.GenHTML(file)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Usage: gogo <filename>")
	}
	GenRightmostDerivation(args[1])
}
