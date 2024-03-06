package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/natanfeitosa/portuscript/playground"
	"github.com/natanfeitosa/portuscript/ptst"
	_ "github.com/natanfeitosa/portuscript/stdlib"
	"github.com/spf13/cobra"
)

var (
	Commit   string = "-"
	Datetime string = "0000-00-00T00:00:00"
	Version  string = "dev"
)

const LongDescription = `
	Uma linguagem orientada a objetos e eventos completamente em português que visa
facilitar os estudos por parte de novos aventureiros no mundo da programação
mas sem ficar apenas criando códigos sem uso prático ou que não refletem o mundo real.

	A documentação completa pode ser encontrada em https://github.com/natanfeitosa/portuscript
`

var codigo string

func main() {
	rootCmd := &cobra.Command{
		Use:     "portuscript [arquivo]",
		Short:   "Portuscript é uma linguagem programação completamente em Português",
		Long:    strings.ReplaceAll(strings.Trim(LongDescription, "\n "), "\t", "    "),
		Version: Version,
		Run: func(cmd *cobra.Command, args []string) {
			cur, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			ctx := ptst.NewContexto(ptst.OpcsContexto{CaminhosPadrao: []string{cur}})
			defer ctx.Terminar()

			// Se não passar um caminho de arquivo nem usar código inline com `-c`
			if codigo == "" && len(args) == 0 {
				playground.Inicializa(ctx, Version, Datetime, Commit)
				return
			}

			// Passou o caminho de um arquio
			if len(args) > 0 {
				_, err = ptst.ExecutarArquivo(ctx, "", args[0], cur, false)
			}

			// Usou código inline
			if codigo != "" {
				_, err = ptst.ExecutarString(ctx, codigo)
			}

			ptst.LancarErro(err)
		},
	}
	
	rootCmd.PersistentFlags().StringVarP(&codigo, "codigo", "c", "", "Use para rodar um código inline.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
