package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"xiyan.life/nup/errhandler"
	"xiyan.life/nup/parser"
	"xiyan.life/nup/translator"
)

func main() {

	// parse command line arguments
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: nup <input file> <output file>")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Setup the input
	is, err := antlr.NewFileStream(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err.Error())
		os.Exit(1)
	}

	// global error handler
	errorHandler := &errhandler.NupErrorListener{}

	// set up the lexer
	lexer := parser.NewNupLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorHandler)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// set up the parser
	p := parser.NewNupParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorHandler)

	// write to file
	output, _ := os.Create(outputFile)
	w := bufio.NewWriter(output)

	defer func(output *os.File) {
		err := w.Flush()
		err = output.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error flushing to the output file: %s\n", err.Error())
		}
	}(output)

	antlr.ParseTreeWalkerDefault.Walk(translator.NewNupListener(w, errorHandler), p.Document())

	// print all the errors
	for _, err := range errorHandler.Errors {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
