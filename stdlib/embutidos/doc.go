package embutidos

import "github.com/natanfeitosa/portuscript/ptst"

func met_emb_doc(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("doc", false, args, 1, 1); err != nil {
		return nil, err
	}

	arg := args[0]
	imp, err := mod.(*ptst.Modulo).Contexto.ObterSimbolo("imprima")

	if err != nil {
		return nil, err
	}

	if obj, ok := arg.(ptst.I_ObtemDoc); ok {
		return ptst.Chamar(imp.Valor, ptst.Tupla{ptst.Texto(obj.ObtemDoc())})
	}

	return ptst.Chamar(imp.Valor, ptst.Tupla{ptst.Texto(arg.Tipo().ObtemDoc())})
}

var _emb_doc = ptst.NewMetodoOuPanic(
	"doc",
	met_emb_doc,
	"doc(objeto) -> Obtem a documentação do objeto",
)