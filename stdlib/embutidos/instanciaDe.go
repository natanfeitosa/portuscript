package embutidos

import "github.com/natanfeitosa/portuscript/ptst"

func met_emb_instanciaDe(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("instanciaDe", false, args, 2, 2); err != nil {
		return nil, err
	}

	return ptst.InstanciaDe(args[0], args[1])
}

var _emb_instanciaDe = ptst.NewMetodoOuPanic(
	"instanciaDe",
	met_emb_instanciaDe,
	"instanciaDe(obj, tipos) -> o parâmetro `tipos` pode ser um tipo ou uma tupla de tipos/classes. Verifica se o obj é instancia de algum dos tipos",
)
