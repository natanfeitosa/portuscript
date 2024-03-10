package playground

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/natanfeitosa/portuscript/ptst"
	"github.com/peterh/liner"
)

const banner = `
Bem vindos ao Portuscript v%s.

(%s) [%s]
`

func homeDirectory() string {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	}
	return os.Getenv("HOME")
}

func ArquivoHistorico(escrita bool) (arquivo *os.File) {
	caminho := path.Join(homeDirectory(), ".historico_portuscript")

	if escrita {
		arquivo, _ = os.OpenFile(caminho, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		// if err != nil {
		// 	return err
		// }
		return
	}

	arquivo, _ = os.Open(caminho)
	defer arquivo.Close()
	return
}

func Inicializa(ctx *ptst.Contexto, version, datetime, commit string) {
	caminho := path.Join(homeDirectory(), ".historico_portuscript")
	arquivo, _ := os.OpenFile(caminho, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	finalizou := false
	finalizar := func() {
		fmt.Printf("Saindo...")
		finalizou = true
	}

	exec := NovoExecutor(ctx)
	exec.RegistrarMetodo(ptst.NewMetodoOuPanic("sair", func(_ ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
		finalizar()
		return nil, nil
	}, ""))

	fmt.Println(fmt.Sprintf(strings.Trim(banner, " \n"), version, datetime, commit))

	line := liner.NewLiner()
	line.ReadHistory(arquivo)

	defer func() {
		line.Close()
		arquivo.Close()
		// exec.Terminar()
	}()

	estado := NewEstado()

	for !finalizou {
		codigo, err := line.Prompt(string(estado.Indicador))
		if err != nil {
			finalizar()
			fmt.Fprintln(os.Stderr, err)
		}

		if len(codigo) < 1 {
			fmt.Println("Entrada vazia")
			continue
		}

		line.AppendHistory(codigo)
		estado.RecalcularEstado(codigo)

		if !estado.Continua {
			exec.ExecutarCodigo(estado.Codigo)
			estado.Codigo = ""
		}
	}

	line.WriteHistory(arquivo)
}
