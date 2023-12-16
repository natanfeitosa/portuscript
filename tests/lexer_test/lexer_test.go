package lexer_test

import (
	"reflect"
	"testing"

	"github.com/natanfeitosa/portuscript/lexer"
)

func TestLexer(t *testing.T) {
	input := `
        var x = 10;
        const PI = 3.1415;
        funcao soma(a, b) {
            retorne a + b;
        }

        se (x == 10) {
            texto = "Hello, World!";
        } senao {
            texto = "Não é 10";
        }

        enquanto (x > 0) {
            x = x - 1;
        }

        para (i = 0; i < 5; i = i + 1) {
            imprima(i);
        }
    `

	lex := lexer.NewLexer(input)
	tokens := []lexer.Token{
		{Tipo: lexer.TokenIdentificador, Valor: "var"},
		{Tipo: lexer.TokenIdentificador, Valor: "x"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenInteiro, Valor: "10"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenIdentificador, Valor: "const"},
		{Tipo: lexer.TokenIdentificador, Valor: "PI"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenDecimal, Valor: "3.1415"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenIdentificador, Valor: "funcao"},
		{Tipo: lexer.TokenIdentificador, Valor: "soma"},
		{Tipo: lexer.TokenAbreParenteses, Valor: "("},
		{Tipo: lexer.TokenIdentificador, Valor: "a"},
		{Tipo: lexer.TokenVirgula, Valor: ","},
		{Tipo: lexer.TokenIdentificador, Valor: "b"},
		{Tipo: lexer.TokenFechaParenteses, Valor: ")"},
		{Tipo: lexer.TokenAbreChaves, Valor: "{"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenRetorne, Valor: "retorne"},
		{Tipo: lexer.TokenIdentificador, Valor: "a"},
		{Tipo: lexer.TokenMais, Valor: "+"},
		{Tipo: lexer.TokenIdentificador, Valor: "b"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenFechaChaves, Valor: "}"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenSe, Valor: "se"},
		{Tipo: lexer.TokenAbreParenteses, Valor: "("},
		{Tipo: lexer.TokenIdentificador, Valor: "x"},
		{Tipo: lexer.TokenIgualIgual, Valor: "=="},
		{Tipo: lexer.TokenInteiro, Valor: "10"},
		{Tipo: lexer.TokenFechaParenteses, Valor: ")"},
		{Tipo: lexer.TokenAbreChaves, Valor: "{"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenIdentificador, Valor: "texto"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenTexto, Valor: "\"Hello, World!\""},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenFechaChaves, Valor: "}"},
		{Tipo: lexer.TokenSenao, Valor: "senao"},
		{Tipo: lexer.TokenAbreChaves, Valor: "{"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenIdentificador, Valor: "texto"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenTexto, Valor: "\"Não é 10\""},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenFechaChaves, Valor: "}"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenEnquanto, Valor: "enquanto"},
		{Tipo: lexer.TokenAbreParenteses, Valor: "("},
		{Tipo: lexer.TokenIdentificador, Valor: "x"},
		{Tipo: lexer.TokenMaiorQue, Valor: ">"},
		{Tipo: lexer.TokenInteiro, Valor: "0"},
		{Tipo: lexer.TokenFechaParenteses, Valor: ")"},
		{Tipo: lexer.TokenAbreChaves, Valor: "{"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenIdentificador, Valor: "x"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenIdentificador, Valor: "x"},
		{Tipo: lexer.TokenMenos, Valor: "-"},
		{Tipo: lexer.TokenInteiro, Valor: "1"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenFechaChaves, Valor: "}"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},

		{Tipo: lexer.TokenPara, Valor: "para"},
		{Tipo: lexer.TokenAbreParenteses, Valor: "("},
		{Tipo: lexer.TokenIdentificador, Valor: "i"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenInteiro, Valor: "0"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenIdentificador, Valor: "i"},
		{Tipo: lexer.TokenMenorQue, Valor: "<"},
		{Tipo: lexer.TokenInteiro, Valor: "5"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenIdentificador, Valor: "i"},
		{Tipo: lexer.TokenIgual, Valor: "="},
		{Tipo: lexer.TokenIdentificador, Valor: "i"},
		{Tipo: lexer.TokenMais, Valor: "+"},
		{Tipo: lexer.TokenInteiro, Valor: "1"},
		{Tipo: lexer.TokenFechaParenteses, Valor: ")"},
		{Tipo: lexer.TokenAbreChaves, Valor: "{"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenIdentificador, Valor: "imprima"},
		{Tipo: lexer.TokenAbreParenteses, Valor: "("},
		{Tipo: lexer.TokenIdentificador, Valor: "i"},
		{Tipo: lexer.TokenFechaParenteses, Valor: ")"},
		{Tipo: lexer.TokenPontoEVirgula, Valor: ";"},
		{Tipo: lexer.TokenNovaLinha, Valor: "\n"},
		{Tipo: lexer.TokenFechaChaves, Valor: "}"},
	}

	for _, esperado := range tokens {
		token := lex.ProximoToken()
		if !reflect.DeepEqual(token, esperado) {
			t.Errorf("Esperado: %v, Obtido: %v", esperado, token)
		}
	}

	// Verifica se o último token é EOF.
	ultimoToken := lex.ProximoToken()
	if ultimoToken.Tipo != lexer.TokenEOF {
		t.Errorf("Esperado: TokenEOF, Obtido: %v", ultimoToken)
	}
}
