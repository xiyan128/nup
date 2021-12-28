package main

import (
	"bufio"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"xiyan.life/nup/parser"
)

func main() {
	is := antlr.NewInputStream(
		`
\bf{this is the start of a new chapter}

Let's write some latex: $x^2 + y^2 = z^2$.

\quote{This is really cool, \it{isn't it?}}
`)
	lexer := parser.NewNupLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewNupParser(stream)

	// write to file
	f, _ := os.Create("output.html")
	w := bufio.NewWriter(f)

	antlr.ParseTreeWalkerDefault.Walk(NewNupListener(w), p.Document())
	w.Flush()
}
