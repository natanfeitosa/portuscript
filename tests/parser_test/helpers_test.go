package parser_test

import (
	"reflect"

	"github.com/natanfeitosa/portuscript/lexer"
	"github.com/natanfeitosa/portuscript/parser"
)

func createParser(code string) *parser.Parser {
	lexer := lexer.NewLexer(code)
	return parser.NewParser(lexer)
}

func isAstEquals(recebida *parser.Programa, esperada *parser.Programa) bool {
	return reflect.DeepEqual(recebida, esperada)
}

func createParserAndCompare(code string, expected *parser.Programa) (error, bool) {
	received, err := createParser(code).Parse()

	if err != nil {
		return err, false
	}

	return nil, isAstEquals(received, expected)
}
