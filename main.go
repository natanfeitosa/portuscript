package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/natanfeitosa/portuscript/cmd"
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

func init() {
	cmd.Commit = Commit
	cmd.Datetime = Datetime
	cmd.Version = Version
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "portuscript [arquivo]",
		Short:   "Portuscript é uma linguagem programação completamente em Português",
		Long:    strings.ReplaceAll(strings.Trim(LongDescription, "\n "), "\t", "    "),
		Version: Version,
	}
	cmd.InstalarComandos(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
