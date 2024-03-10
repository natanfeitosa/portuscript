package playground

import (
	"fmt"
	"os"

	"github.com/natanfeitosa/portuscript/ptst"
)

type Executor struct {
	Contexto *ptst.Contexto
	Modulo   *ptst.Modulo
}

func NovoExecutor(ctx *ptst.Contexto) *Executor {
	exec := new(Executor)
	exec.Contexto = ctx
	exec.Modulo, _ = ctx.InicializarModulo(&ptst.ModuloImpl{
		Info: ptst.ModuloInfo{
			Arquivo: "<playground>",
		},
	})

	return exec
}

func (e *Executor) ExecutarCodigo(codigo string) {
	ast, err := e.Contexto.StringParaAst(codigo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var resultado, texto ptst.Objeto

	if resultado, err = e.Contexto.AvaliarAst(ast, e.Modulo.Escopo); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if texto, err = ptst.NewTexto(resultado); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	
	fmt.Println(texto)
}

func (e *Executor) RegistrarMetodo(metodo *ptst.Metodo) error {
	return e.Modulo.Escopo.DefinirSimbolo(
		ptst.NewVarSimbolo(
			metodo.Nome,
			metodo,
		),
	)
}

func (e *Executor) Terminar() {
	e.Contexto.Terminar()
}
