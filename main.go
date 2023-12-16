package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/natanfeitosa/portuscript/parser"
	"github.com/natanfeitosa/portuscript/playground"
	"github.com/natanfeitosa/portuscript/ptst"
	_ "github.com/natanfeitosa/portuscript/stdlib/embutidos"
	"github.com/spf13/cobra"
)

const Version = "0.0.1"
const LongDescription = `
	Uma linguagem orientada a objetos e eventos completamente em português que visa
facilitar os estudos por parte de novos aventureiros no mundo da programação
mas sem ficar apenas criando códigos sem uso prático ou que não refletem o mundo real.

	A documentação completa pode ser encontrada em https://github.com/natanfeitosa/portuscript
`

var codigo string

var verast = &cobra.Command{
	Use:   "ast [arquivo.ptst]",
	Short: "Imprime a ast de um código",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Lookup("codigo").Value.String() != "" {
			ast, err := parser.NewParserFromString(codigo).Parse()

			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}

			astbytes, err := parser.Ast2string(ast)

			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}

			fmt.Println(string(astbytes))
			return
		}
		fmt.Println(args)
	},
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "portuscript [arquivo]",
		Short:   "Portuscript é uma linguagem programação completamente em Português",
		Long:    strings.ReplaceAll(strings.Trim(LongDescription, "\n "), "\t", "    "),
		Version: Version,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				ptst.InicializaDeArquivo(args[0])
				return
			}

			if codigo != "" {
				// fmt.Printf("codigo: %v\n", codigo)
				ptst.InicializaDeString(codigo)
				return
			}

			playground.Inicializa()
		},
	}

	// rootCmd.AddCommand(verast)
	rootCmd.PersistentFlags().StringVarP(&codigo, "codigo", "c", "", "Use para rodar um código inline.")

	// rootCmd.SetVersionTemplate("portuscipt versão " + Version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
