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
\bf[id="first"]{this is the start of a \ref[to="q1"]{new chapter}}

Let's write some latex: $x^2 + y^2 = z^2$. This conclusion is drawn from \ref[to="q2"]{this equation}

\quote[cite="https://google.com"]{
This is really cool, \it{isn't it?}
\deref[id="q2"]{$\frac{1}{2} \int_2 x^2 dx$}

This is a second paragraph in the quote, 
and this is not a new paragraph.

This, however, is a new paragraph.
}

\para{a para}
\deref[id="q1"]{this is been referenced}

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
