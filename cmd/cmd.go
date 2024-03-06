package cmd

import "github.com/spf13/cobra"

var (
	Commit   string = "-"
	Datetime string = "0000-00-00T00:00:00"
	Version  string = "dev"
)

func InstalarComandos(raiz *cobra.Command) {
	raiz.AddCommand(comandoAtualize())
	raiz.AddCommand(comandoExecutar())
}