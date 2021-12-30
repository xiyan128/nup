package main

import (
	"bufio"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"xiyan.life/nup/parser"
	"xiyan.life/nup/translator"
)

func main() {
	is := antlr.NewInputStream(
		`
\title{Nup, a markup language for the web}

\bf[id="first"]{this is the start of a \ref[to="q1"]{new chapter}}

Let's write some latex: $x^2 + y^2 = z^2$. This conclusion is drawn from \ref[to="q2"]{this equation}

\heading{this is a heading}

\quote[cite="https://google.com"]{
This is really cool, \it{isn't it?}\ref[to="q1"]{}
\deref[id="q2"]{$\frac{1}{2} \int_2 x^2 dx$}

This is a second paragraph in the quote, 
and this is not a new paragraph.

This, however, is a new paragraph.
}

\para{a para}
\deref[id="q1"]{this ` + "`function`" + ` is been referenced}

another para

\figure[src="https://www.w3schools.com/images/colorpicker2000.png"]{caption}

\box[title="this is the title"]{this is the content}

` + "```[lang=\"cpp\"]" + `
package main
import "fmt"
func main() {
	fmt.Println("Hello, 世界")
}` + "```" + `
nea \para{dsa}
`)
	lexer := parser.NewNupLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewNupParser(stream)

	// write to file
	f, _ := os.Create("output.html")
	w := bufio.NewWriter(f)

	antlr.ParseTreeWalkerDefault.Walk(translator.NewNupListener(w), p.Document())
	w.Flush()
}
