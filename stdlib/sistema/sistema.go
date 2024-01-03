package sistema

import (
	"runtime"

	"github.com/natanfeitosa/portuscript/ptst"
)

func init() {
	constantes := ptst.Mapa{
		"ARQUITETURA": ptst.Texto(runtime.GOARCH),
		"NOME": ptst.Texto(runtime.GOOS),
	}

	metodos := []*ptst.Metodo{}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome:          "sistema",
				CaminhoModulo: "stdlib/sistema",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}
