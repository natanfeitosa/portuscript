package parser_test

import (
	"testing"

	"github.com/natanfeitosa/portuscript/parser"
)

func TestSimplesAtribuicaoVariavel(t *testing.T) {
	esperada := &parser.Programa{}
	decaracao1 := &parser.DeclVar{Constante: false, Nome: "var1", Tipo: "inteiro"}
	esperada.Declaracoes = append(esperada.Declaracoes, decaracao1)

	decaracao2 := &parser.DeclVar{Constante: false, Nome: "var2", Inicializador: &parser.TextoLiteral{Valor: "\"inteiro?\""}}
	esperada.Declaracoes = append(esperada.Declaracoes, decaracao2)

	code := `
	var var1: inteiro;
	var var2 = "inteiro?";
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestSimplesAtribuicaoConstante(t *testing.T) {
	esperada := &parser.Programa{}
	decaracao1 := &parser.DeclVar{Constante: true, Nome: "const1", Tipo: "inteiro"}
	esperada.Declaracoes = append(esperada.Declaracoes, decaracao1)

	decaracao2 := &parser.DeclVar{Constante: true, Nome: "const2", Inicializador: &parser.TextoLiteral{Valor: "\"inteiro?\""}}
	esperada.Declaracoes = append(esperada.Declaracoes, decaracao2)

	code := `
	const const1: inteiro;
	const const2 = "inteiro?";
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestSimplesReatribuicao(t *testing.T) {
	esperada := &parser.Programa{}

	inteiro := &parser.InteiroLiteral{Valor: "1"}

	esperada.Declaracoes = []parser.BaseNode{
		&parser.Reatribuicao{Objeto: &parser.Identificador{Nome: "variavel"}, Operador: "+=", Expressao: inteiro},
		&parser.Reatribuicao{Objeto: &parser.Identificador{Nome: "variavel"}, Operador: "-=", Expressao: inteiro},
		&parser.Reatribuicao{Objeto: &parser.Identificador{Nome: "variavel"}, Operador: "/=", Expressao: inteiro},
		&parser.Reatribuicao{Objeto: &parser.Identificador{Nome: "variavel"}, Operador: "*=", Expressao: inteiro},
	}

	code := `
	variavel += 1;
	variavel -= 1;
	variavel /= 1;
	variavel *= 1;
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}
