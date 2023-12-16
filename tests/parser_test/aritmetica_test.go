package parser_test

import (
	"testing"

	"github.com/natanfeitosa/portuscript/parser"
)

func TestMultiplicacao(t *testing.T) {
	literal := &parser.InteiroLiteral{Valor: "1"}
	identif := &parser.Identificador{Nome: "a"}

	esperada := &parser.Programa{}
	esperada.Declaracoes = []parser.BaseNode{
		&parser.OpBinaria{Esq: literal, Operador: "*", Dir: literal},
		&parser.OpBinaria{Esq: literal, Operador: "*", Dir: identif},
		&parser.OpBinaria{Esq: identif, Operador: "*", Dir: identif},
		&parser.OpBinaria{Esq: identif, Operador: "*", Dir: literal},
	}

	code := `
	1 * 1
	1 * a
	a * a
	a * 1
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}
