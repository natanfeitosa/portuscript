package embutidos

import "github.com/natanfeitosa/portuscript/ptst"

func met_emb_int(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("int", false, args, 0, 1); err != nil {
		return nil, err
	}

	return ptst.NewInteiro(args[0])
}

var _emb_int = ptst.NewMetodoOuPanic(
	"int",
	met_emb_int,
	"int(objeto) -> Recebe um objeto e retorna uma representação numérica do tipo inteiro, se possível",
)
