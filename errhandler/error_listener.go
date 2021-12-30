package errhandler

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SyntaxError struct {
	line, column int
	msg          string
}

type RuntimeError struct {
	line, column int
	msg          string
}

func (r RuntimeError) Error() string {
	return fmt.Sprintf("Runtime error at line %d, column %d: %s", r.line, r.column, r.msg)
}

func (c SyntaxError) Error() string {
	return fmt.Sprintf("Syntax error at line %d, column %d: %s", c.line, c.column, c.msg)
}

type NupErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

func (c *NupErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &SyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (c *NupErrorListener) RuntimeError(line, column int, msg string) {
	c.Errors = append(c.Errors, &RuntimeError{
		line:   line,
		column: column,
		msg:    msg,
	})
}
