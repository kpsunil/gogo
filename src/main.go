package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	"gogo/src/asm"
	"gogo/src/parser"
	"gogo/src/scanner"
	"gogo/src/tac"
)

// GenToken generates the tokens returned by lexer from the input program.
func GenToken(file string) {
	scanner.PrintTokens(file)
}

// GenIR generates the IR instructions from the input program.
func GenIR(file string) {
	parser.GenProductions(file)
}

// GenAsmFromIR generates the assembly code using IR generated from the input program.
func GenAsmFromIR(file string) {
	// A journey of a thousand miles begins with a single step. This is that
	// step.
	asm.CodeGen(tac.GenTAC(file))
}

// GenAsm generates the assembly code from the input go program.
func GenAsm(file string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outChan := make(chan string)

	// Buffer the output generated by GenIR into a pipe. The output is
	// copied in a separate goroutine so that printing doesn't block.
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChan <- buf.String()
	}()

	GenIR(file)
	w.Close()
	// Restore the original state (before pipe was created).
	os.Stdout = old
	ir := <-outChan

	// Create a temporary IR file which will be used to generate assembly.
	// TODO: Update function signatures to avoid I/O.
	irFile := "tmp.ir"
	f, err := os.Create(irFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	if _, err = writer.WriteString(ir); err != nil {
		log.Fatal(err)
	}
	writer.Flush()

	GenAsmFromIR(irFile)
	if err = os.Remove(irFile); err != nil {
		log.Fatal(err)
	}
}

// GenHTML generates the rightmost derivations used in the bottom-up parsing and
// pretty-prints them in an HTML format.
func GenHTML(file string) {
	parser.RightmostDerivation(file)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Usage: gogo <filename>")
	}
	GenIR(args[1])
}
