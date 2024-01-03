package matematica

import (
	"math"

	"github.com/natanfeitosa/portuscript/ptst"
)

func met_mat_teto(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("teto", false, args, 1, 1); err != nil {
		return nil, err
	}

	num, err := ptst.NewDecimal(args[0])
	if err != nil {
		return nil, err
	}

	return ptst.Inteiro(math.Ceil(float64(num.(ptst.Decimal)))), nil
}

var _mat_teto = ptst.NewMetodoOuPanic(
	"teto",
	met_mat_teto,
	"teto(decimal) -> Retorna o numero arredondado para cima",
)