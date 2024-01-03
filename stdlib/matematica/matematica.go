package matematica

import (
	"math"

	"github.com/natanfeitosa/portuscript/ptst"
)

func init() {
	constantes := ptst.Mapa{
		"PI": ptst.Decimal(math.Pi),
		"E": ptst.Decimal(math.E),
	}

	metodos := []*ptst.Metodo{
		_mat_raiz,
		_mat_potencia,
		_mat_absoluto,
		_mat_piso,
		_mat_teto,
	}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome: "matematica",
				CaminhoModulo: "stdlib/matematica",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}