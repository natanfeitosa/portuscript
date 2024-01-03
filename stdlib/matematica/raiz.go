package matematica

import "github.com/natanfeitosa/portuscript/ptst"

func met_mat_raiz(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("raiz", false, args, 1, 2); err != nil {
		return nil, err
	}

	indice := ptst.Decimal(2.0)

	if len(args) > 1 {
		dec, err := ptst.NewDecimal(args[1])
		if err != nil {
			return nil, err
		}

		indice = dec.(ptst.Decimal)
	}

	return met_mat_potencia(inst, ptst.Tupla{args[0], 1.0/indice})
}

var _mat_raiz = ptst.NewMetodoOuPanic(
	"raiz",
	met_mat_raiz,
	"raiz(radicando, indice?) -> Retorna a raiz de radicando por indice. Se indice não for definido, o padrão é 2 (raiz quadrada do radicando)",
)
