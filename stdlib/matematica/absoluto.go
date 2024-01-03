package matematica

import (
	"math"

	"github.com/natanfeitosa/portuscript/ptst"
)

func met_mat_absoluto(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("absoluto", false, args, 1, 1); err != nil {
		return nil, err
	}

	numero, err := ptst.NewDecimal(args[0])
	if err != nil {
		return nil, err
	}

	return ptst.Decimal(math.Abs(float64(numero.(ptst.Decimal)))), nil
}

var _mat_absoluto = ptst.NewMetodoOuPanic(
	"absoluto",
	met_mat_absoluto,
	"absoluto(numero) -> Retorna o valor absoluto de um número, isso é, sem sinal caso houver",
)
