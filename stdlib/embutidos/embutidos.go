package embutidos

import (
	"github.com/natanfeitosa/portuscript/ptst"
)

func init() {
	constantes := ptst.Mapa{
		"Verdadeiro": ptst.Verdadeiro,
		"Falso":      ptst.Falso,
		"Nulo":       ptst.Nulo,
	}

	metodos := []*ptst.Metodo{
		_emb_imprima,
		_emb_leia_,
		_emb_doc,
		_emb_int,
		_emb_texto,
		_emb_tamanho,
	}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome: "embutidos",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}
