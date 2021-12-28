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
\bf[id="first"]{this is the start of a new chapter}

Let's write some latex: $x^2 + y^2 = z^2$.

\quote[cite="https://google.com"]{
This is really cool, \it{isn't it?}

This is a second paragraph in the quote, 
and this is not a new paragraph.

This, however, is a new paragraph.
}

\para{a para}

\para{another para}

nea \para{dsa}
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
