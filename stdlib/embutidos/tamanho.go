package embutidos

import "github.com/natanfeitosa/portuscript/ptst"

func emb_tamanho_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("tamanho", false, args, 1, 1); err != nil {
		return nil, err
	}

	if obj, ok := args[0].(ptst.I__tamanho__); ok {
		return obj.O__tamanho__()
	}

	return nil, ptst.NewErroF(ptst.TipagemErro, "Objeto do tipo '%s' não implementa a interface '__tamanho__'.", args[0].Tipo().Nome)
}

var _emb_tamanho = ptst.NewMetodoOuPanic(
	"tamanho",
	emb_tamanho_fn,
	"tamanho(obj) -> Retorna o tamanho do objeto, mas se o objeto não implementar o método `__tamanho__`, um erro será lançado",
)