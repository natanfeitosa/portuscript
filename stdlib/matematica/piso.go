package matematica

import (
	"math"

	"github.com/natanfeitosa/portuscript/ptst"
)

func met_mat_piso(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("piso", false, args, 1, 1); err != nil {
		return nil, err
	}

	num, err := ptst.NewDecimal(args[0])
	if err != nil {
		return nil, err
	}

	return ptst.Inteiro(math.Floor(float64(num.(ptst.Decimal)))), nil
}

var _mat_piso = ptst.NewMetodoOuPanic(
	"piso",
	met_mat_piso,
	"piso(decimal) -> Retorna o numero arredondado para baixo",
)