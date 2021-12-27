package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"xiyan.life/nup/parser"
)

func main() {
	// print res
	is := antlr.NewInputStream(
		`
this is a new paragraph

this is a new block \bf[id="123"]{this is bold \it[id="456"]{this is italic \href{this is href} still italic \ic{no
longer href } again italic} this is bold} this is nothing
`)

	lexer := parser.NewNupLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewNupParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&NupListener{}, p.Document())
}
