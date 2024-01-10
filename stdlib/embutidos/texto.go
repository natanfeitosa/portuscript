package embutidos

import "github.com/natanfeitosa/portuscript/ptst"

func met_emb_texto(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("texto", false, args, 0, 1); err != nil {
		return nil, err
	}

	return ptst.NewTexto(args[0])
}

var _emb_texto = ptst.NewMetodoOuPanic(
	"texto",
	met_emb_texto,
	"texto(objeto) -> Recebe um objeto e retorna uma representação no tipo texto, se possível",
)