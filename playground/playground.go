package playground

import (
	"fmt"
	"os"
	"strings"

	"github.com/natanfeitosa/portuscript/ptst"
	"github.com/peterh/liner"
)

const banner = `
Bem vindos ao Portuscript v%s.

(%s) [%s]
`

func Inicializa(ctx *ptst.Contexto, version, datetime, commit string) {
	finalizado := false
	finalizar := func() {
		finalizado = true
	}

	mod, _ := ctx.InicializarModulo(&ptst.ModuloImpl{
		Info: ptst.ModuloInfo{
			Arquivo: "<playground>",
		},
	})

	imprima, err := mod.Contexto.Modulos.Embutidos.M__obtem_attributo__("imprima")

	if err != nil {
		finalizar()
		ptst.LancarErro(err)
	}

	mod.Escopo.DefinirSimbolo(
		ptst.NewVarSimbolo(
			"sair",
			ptst.NewMetodoOuPanic("sair", func(mod ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
				finalizar()
				return nil, nil
			}, ""),
		),
	)

	ptst.Chamar(imprima, ptst.Texto(fmt.Sprintf(strings.Trim(banner, "\n "), version, datetime, commit)))

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
			ptst.Chamar(imprima, ptst.Texto("Entrada vazia"))
			continue
		}

		ast, err := ctx.StringParaAst(codigo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		result, err := ctx.AvaliarAst(ast, mod.Escopo)
		// fmt.Printf("\n\n%t\n%t\n\n", result, err)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if result != nil {
			result, _ := ptst.NewTexto(result)
			ptst.Chamar(imprima,result.(ptst.Texto))
		}
	}
}
