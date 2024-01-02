package playground

import (
	"fmt"
	"os"
	"strings"

	"github.com/natanfeitosa/portuscript/parser"
	"github.com/natanfeitosa/portuscript/ptst"
	"github.com/peterh/liner"
)

func criarContexto() *ptst.Contexto {
	contexto := ptst.NewContexto("<playground>")
	modulos := ptst.NewTabelaModulos()
	modulos.NewModulo(contexto, ptst.ObtemImplModulo("embutidos"))
	contexto.Modulos = modulos

	return contexto
}

func stringParaAst(codigo string) (*parser.Programa, error) {
	return parser.NewParserFromString(codigo).Parse()
}

// func obterMetodosPrincipais(ctx *ptst.Contexto) (leia *ptst.Metodo, imprima *ptst.Metodo, err error) {
// 	simbolo, err := ctx.ObterSimbolo("leia")
// 	if err != nil {
// 		return
// 	}

// 	leia = simbolo.Valor.(*ptst.Metodo)

// 	simbolo, err = ctx.ObterSimbolo("imprima")
// 	if err != nil {
// 		return
// 	}

// 	imprima = simbolo.Valor.(*ptst.Metodo)
// 	return
// }

const banner = `
Bem vindos ao Portuscript v%s.

(%s) [%s]
`

func Inicializa(version, datetime, commit string) {
	finalizado := false
	finalizar := func() {
		finalizado = true
	}

	contextoRaiz := criarContexto()
	simbolo, err := contextoRaiz.ObterSimbolo("imprima")

	if err != nil {
		finalizar()
		ptst.LancarErro(err)
	}

	imprima := simbolo.Valor.(*ptst.Metodo)

	contextoRaiz.DefinirSimboloLocal(
		ptst.NewVarSimbolo(
			"sair",
			ptst.NewMetodoOuPanic("sair", func(mod ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
				finalizar()
				return nil, nil
			}, ""),
		),
	)

	imprima.O__chame__(
		ptst.Tupla{
			ptst.Texto(
				fmt.Sprintf(strings.Trim(banner, "\n "), version, datetime, commit),
			),
		},
	)

	line := liner.NewLiner()
	defer line.Close()

	for !finalizado {
		codigo, err := line.Prompt(">>> ")

		if err != nil {
			finalizar()
			fmt.Fprintln(os.Stderr, err)
		}

		line.AppendHistory(codigo)

		if len(codigo) < 1 {
			imprima.O__chame__(ptst.Tupla{ptst.Texto("Entrada vazia")})
			continue
		}

		ast, err := stringParaAst(codigo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		
		interpretador := &ptst.Interpretador{
			Ast:      ast,
			Contexto: contextoRaiz,
			Caminho:  contextoRaiz.Caminho,
		}

		result, err := interpretador.Visite(ast.Declaracoes)
		// fmt.Printf("\n\n%t\n%t\n\n", result, err)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if result != nil {
			result, _ := ptst.NewTexto(result)
			imprima.O__chame__(ptst.Tupla{result.(ptst.Texto)})
		}
	}
}
