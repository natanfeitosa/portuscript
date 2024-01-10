package embutidos

import "github.com/natanfeitosa/portuscript/ptst"

func met_emb_mesmoTipo(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("mesmoTipo", false, args, 2, 2); err != nil {
		return nil, err
	}

	return ptst.Booleano(ptst.MesmoTipo(args[0], args[1])), nil
}

var _emb_mesmoTipo = ptst.NewMetodoOuPanic(
	"mesmoTipo",
	met_emb_mesmoTipo,
	"mesmoTipo(obj1, obj2) -> Verifica se os dois objetos são do mesmo tipo",
)
