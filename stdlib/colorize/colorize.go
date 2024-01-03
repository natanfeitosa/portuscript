package colorize

import (
	"os"

	"github.com/natanfeitosa/portuscript/ptst"
)

const (
	InicioCodigo = "\x1b["
	FimCodigo    = "m"
	ResetCodigo  = "\x1b[0m"
	TplFgRGB     = "38;2;%d;%d;%d"
	TplBgRGB     = "48;2;%d;%d;%d"
)

var SuportaCores = os.Getenv("NO_COLOR") == ""

func init() {
	constantes := ptst.Mapa{
		"FUNDO": &Background{},
		"TEXTO": &Foreground{},
		"SUPORTA": ptst.Booleano(SuportaCores),
	}

	metodos := []*ptst.Metodo{
		_color_converteRGB,
		_color_imprimac,
	}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome:          "colorize",
				CaminhoModulo: "stdlib/colorize",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}
