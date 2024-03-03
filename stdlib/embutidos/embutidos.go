package embutidos

import (
	"github.com/natanfeitosa/portuscript/ptst"
)

func init() {
	constantes := ptst.Mapa{
		"Verdadeiro": ptst.Verdadeiro,
		"Falso":      ptst.Falso,
		"Nulo":       ptst.Nulo,
		"Inteiro":    ptst.TipoInteiro,
		"Decimal":    ptst.TipoDecimal,
		"Texto":      ptst.TipoTexto,
		// "Lista":      ptst.TipoLista,
		// "Tupla":      ptst.TipoTupla,
		// "Mapa":       ptst.TipoMapa,
		"Booleano": ptst.TipoBooleano,
	}

	metodos := []*ptst.Metodo{
		_emb_imprima,
		_emb_leia,
		_emb_doc,
		_emb_int,
		_emb_texto,
		_emb_tamanho,
		_emb_instanciaDe,
		_emb_mesmoTipo,
		ptst.NewMetodoOuPanic(
			"tipo",
			func(_ ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
				if err := ptst.VerificaNumeroArgumentos("tipo", false, args, 1, 1); err != nil {
					return nil, err
				}

				return args[0].Tipo(), nil
			},
			"Obtem o tipo de um objeto",
		),
	}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome: "embutidos",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}
