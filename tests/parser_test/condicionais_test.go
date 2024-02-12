package parser_test

import (
	"testing"

	"github.com/natanfeitosa/portuscript/parser"
)

func TestExpresaoSeSemCorpo(t *testing.T) {
	esperada := &parser.Programa{}
	esperada.Declaracoes = []parser.BaseNode{
		&parser.ExpressaoSe{
			Condicao: &parser.ConstanteLiteral{Valor: "Verdadeiro"},
			Corpo: &parser.Bloco{},
		},
	}

	code := `
	se(Verdadeiro) {}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestExpresaoSe(t *testing.T) {
	esperada := &parser.Programa{}
	esperada.Declaracoes = []parser.BaseNode{
		&parser.ExpressaoSe{
			Condicao: &parser.ConstanteLiteral{Valor: "Verdadeiro"},
			Corpo: &parser.Bloco{
				Declaracoes: []parser.BaseNode{
					&parser.ChamadaFuncao{
						Identificador: &parser.Identificador{Nome: "imprima"},
						Argumentos: []parser.BaseNode{
							&parser.TextoLiteral{
								Valor: "\"Verdadeiro é definitivamente verdadeiro\"",
							},
						},
					},
				},
			},
		},
	}

	code := `
	se (Verdadeiro) {
		imprima("Verdadeiro é definitivamente verdadeiro")
	}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}
