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
