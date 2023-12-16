package ptst

import (
	"fmt"
	"os"
	"strings"

	"github.com/natanfeitosa/portuscript/parser"
)

func inicializa(parser *parser.Parser, caminho string) {
	ast, err := parser.Parse()

	if err != nil {
		LancarErro(err)
		return
	}

	interpret := &Interpretador{Ast: ast, Caminho: Texto(caminho)}

	// FIXME: isso pode ser melhorado, não?
	interpret.Contexto = NewContexto(interpret.Caminho)
	modulos := NewTabelaModulos()
	modulos.NewModulo(interpret.Contexto, ObtemImplModulo("embutidos"))
	interpret.Contexto.Modulos = modulos

	if _, err := interpret.Inicializa(); err != nil {
		LancarErro(err)
	}
}

func InicializaDeString(codigo string) {
	inicializa(parser.NewParserFromString(codigo), "<string>")
}

func InicializaDeArquivo(caminho string) {
	if !strings.HasSuffix(caminho, ".ptst") {
		LancarErro(fmt.Errorf("o arquivo '%s' pode não ser um arquivo Portuscript válido. Você errou a extensão '.ptst'?", caminho))
	}

	codigo, err := os.ReadFile(caminho)
	if err != nil {
		LancarErro(err)
		return
	}

	// decodificador := charmap.ISO8859_1.NewDecoder()
	// codigo, _, err := transform.Bytes(decodificador, fonte)

	if err != nil {
		LancarErro(err)
		return
	}
	
	inicializa(parser.NewParserFromString(string(codigo)), caminho)
}
