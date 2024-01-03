package matematica

import (
	"math"

	"github.com/natanfeitosa/portuscript/ptst"
)

func met_mat_potencia(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("potencia", false, args, 2, 2); err != nil {
		return nil, err
	}

	var base, expoente ptst.Objeto
	expoente = ptst.Decimal(2.0)
	base = args[0]

	if len(args) > 1 {
		expoente = args[1]
	}

	var err error
	if base, err = ptst.NewDecimal(base); err != nil {
		return nil, err
	}

	if expoente, err = ptst.NewDecimal(expoente); err != nil {
		return nil, err
	}

	potencia := math.Pow(float64(base.(ptst.Decimal)), float64(expoente.(ptst.Decimal)))
	return ptst.Decimal(potencia), nil
}

var _mat_potencia = ptst.NewMetodoOuPanic(
	"potencia",
	met_mat_potencia,
	"potencia(base, expoente) -> Retorna a potencia de base ^ expoente",
)
