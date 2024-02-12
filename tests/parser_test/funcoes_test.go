package parser_test

import (
	"testing"

	"github.com/natanfeitosa/portuscript/parser"
)

func TestDeclareFuncaoVazia(t *testing.T) {
	esperada := &parser.Programa{}
	funcao := &parser.DeclFuncao{Nome: "nenhumaOperacao", Corpo: &parser.Bloco{}}
	esperada.Declaracoes = append(esperada.Declaracoes, funcao)

	code := `
	func nenhumaOperacao() {}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestDeclareFuncaoComParametrosSemCorpo(t *testing.T) {
	esperada := &parser.Programa{}
	funcao := &parser.DeclFuncao{
		Nome: "nenhumaOperacao",
		Parametros: []*parser.DeclFuncaoParametro{
			{Nome: "nome", Tipo: "texto"},
			{Nome: "idade", Padrao: &parser.InteiroLiteral{Valor: "21"}},
		},
		Corpo: &parser.Bloco{},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, funcao)

	code := `
	func nenhumaOperacao(nome: texto, idade = 21) {}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestDeclareFuncaoComParametrosECorpo(t *testing.T) {
	esperada := &parser.Programa{}
	funcao := &parser.DeclFuncao{
		Nome: "nenhumaOperacao",
		Parametros: []*parser.DeclFuncaoParametro{
			{Nome: "nome", Tipo: "texto"},
			{Nome: "idade", Padrao: &parser.InteiroLiteral{Valor: "21"}},
		},
		Corpo: &parser.Bloco{
			Declaracoes: []parser.BaseNode{
				&parser.DeclVar{
					Constante:     false,
					Nome:          "ano",
					Tipo:          "inteiro",
					Inicializador: &parser.InteiroLiteral{Valor: "2023"},
				},
			},
		},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, funcao)

	code := `
	func nenhumaOperacao(nome: texto, idade = 21) {
		var ano: inteiro = 2023;
	}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestDeclareFuncaoSemParametrosComCorpo(t *testing.T) {
	esperada := &parser.Programa{}
	funcao := &parser.DeclFuncao{
		Nome: "nenhumaOperacao",
		Corpo: &parser.Bloco{
			Declaracoes: []parser.BaseNode{
				&parser.DeclVar{
					Constante:     false,
					Nome:          "ano",
					Tipo:          "inteiro",
					Inicializador: &parser.InteiroLiteral{Valor: "2023"},
				},
			},
		},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, funcao)

	code := `
	func nenhumaOperacao() {
		var ano: inteiro = 2023;
	}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestDeclareFuncaoComRetornoInteiro(t *testing.T) {
	esperada := &parser.Programa{}
	funcao := &parser.DeclFuncao{
		Nome: "nenhumaOperacao",
		Corpo: &parser.Bloco{
			Declaracoes: []parser.BaseNode{
				&parser.RetorneNode{
					Expressao: &parser.InteiroLiteral{Valor: "2023"},
				},
			},
		},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, funcao)

	code := `
	func nenhumaOperacao() {
		retorne 2023;
	}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestDeclareFuncaoComRetornoVazio(t *testing.T) {
	esperada := &parser.Programa{}
	funcao := &parser.DeclFuncao{
		Nome: "nenhumaOperacao",
		Corpo: &parser.Bloco{
			Declaracoes: []parser.BaseNode{
				&parser.RetorneNode{},
			},
		},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, funcao)

	code := `
	func nenhumaOperacao() {
		retorne;
	}
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestChamadaFuncaoComArgumentos(t *testing.T) {
	esperada := &parser.Programa{}
	chamada := &parser.ChamadaFuncao{
		Identificador: &parser.Identificador{Nome: "nenhumaOperacao"},
		Argumentos: []parser.BaseNode{
			&parser.InteiroLiteral{Valor: "2023"},
			&parser.TextoLiteral{Valor: "\"portuscript\""},
		},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, chamada)

	code := `nenhumaOperacao(2023, "portuscript")`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestChamadaFuncaoSemArgumentos(t *testing.T) {
	esperada := &parser.Programa{}
	chamada := &parser.ChamadaFuncao{
		Identificador: &parser.Identificador{Nome: "nenhumaOperacao"},
	}
	esperada.Declaracoes = append(esperada.Declaracoes, chamada)

	code := `nenhumaOperacao()`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}

func TestDeclareChameFuncaoComArgumentos(t *testing.T) {
	esperada := &parser.Programa{
		Declaracoes: []parser.BaseNode{
			&parser.DeclFuncao{
				Nome: "soma",
				Parametros: []*parser.DeclFuncaoParametro{
					{Nome: "a"},
					{Nome: "b"},
				},
				Corpo: &parser.Bloco{
					Declaracoes: []parser.BaseNode{
						&parser.RetorneNode{
							Expressao: &parser.OpBinaria{
								Esq: &parser.Identificador{Nome: "a"},
								Operador: "+",
								Dir: &parser.Identificador{Nome: "b"},
							},
						},
					},
				},
			},
			&parser.ChamadaFuncao{
				Identificador: &parser.Identificador{Nome: "imprima"},
				Argumentos: []parser.BaseNode{
					&parser.TextoLiteral{Valor: "\"2 + 2 =\""},
					&parser.ChamadaFuncao{
						Identificador: &parser.Identificador{Nome: "soma"},
						Argumentos: []parser.BaseNode{
							&parser.InteiroLiteral{Valor: "2"},
							&parser.InteiroLiteral{Valor: "2"},
						},
					},
				},
			},
		},
	}

	code := `
	func soma(a, b) {
		retorne a + b;
	}
	
	imprima("2 + 2 =", soma(2, 2))
	`

	err, ok := createParserAndCompare(code, esperada)

	if !ok {
		t.Error(err)
	}
}
