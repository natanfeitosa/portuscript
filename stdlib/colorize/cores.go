package colorize

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/ptst"
)

type Background struct{}

var TipoBackground = ptst.TipoObjeto.NewTipo(
	"Background",
	"",
)

func (b *Background) Tipo() *ptst.Tipo {
	return TipoBackground
}

type Foreground struct{}

var TipoForeground = ptst.TipoObjeto.NewTipo(
	"Foreground",
	"",
)

func (f *Foreground) Tipo() *ptst.Tipo {
	return TipoForeground
}

// Fonte das cores https://www.oocities.org/tutorialdhtml/a12.htm
var cores = []*cor{
	{"vermelho", "ff0000"},
	{"lima", "00ff00"},
	{"azul", "0000ff"},
	{"amarelo", "ffff00"},
	{"agua", "00ffff"},
	{"fuchsia", "ff00ff"},
	{"branco", "fff"},
	{"preto", "000"},
}

func init() {
	for _, cor := range cores {
		r, g, b, err := HexParaRgb(cor.Hex)
		if err != nil {
			panic(err)
		}

		TipoBackground.Mapa[cor.Nome] = ptst.NewMetodoOuPanic(
			cor.Nome,
			criaRenderizadorDeCores(
				r, g, b,
				true,
			),
			fmt.Sprintf("Define a cor %s ao fundo do texto", cor.Nome),
		)

		TipoForeground.Mapa[cor.Nome] = ptst.NewMetodoOuPanic(
			cor.Nome,
			criaRenderizadorDeCores(
				r, g, b,
				false,
			),
			fmt.Sprintf("Define a cor %s ao texto", cor.Nome),
		)
	}
}