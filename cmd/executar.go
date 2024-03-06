package cmd

import (
	"os"

	"github.com/natanfeitosa/portuscript/playground"
	"github.com/natanfeitosa/portuscript/ptst"
	_ "github.com/natanfeitosa/portuscript/stdlib"
	"github.com/spf13/cobra"
)

var codigo string

func comandoExecutar() *cobra.Command {
	executar := &cobra.Command{
		Use:     "executar [arquivo]",
		Short:   "Executa um arquivo ou algum código inline",
		Aliases: []string{"exec"},
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

	executar.PersistentFlags().StringVarP(&codigo, "codigo", "c", "", "Use para rodar um código inline.")
	return executar
}
