package embutidos

import (
	"github.com/natanfeitosa/portuscript/ptst"
)

func registrarTipos(tipos []*ptst.Tipo, mapa ptst.Mapa) {
	for _, tipo := range tipos {
		mapa[tipo.Nome] = tipo
	}
}

func init() {
	constantes := ptst.Mapa{
		"Verdadeiro": ptst.Verdadeiro,
		"Falso":      ptst.Falso,
		"Nulo":       ptst.Nulo,
	}

	registrarTipos(
		[]*ptst.Tipo{
			ptst.TipoInteiro,
			ptst.TipoDecimal,
			ptst.TipoTexto,
			// ptst.TipoLista,
			// ptst.TipoTupla,
			// ptst.TipoMapa,
			ptst.TipoBooleano,
			ptst.TipoBytes,

			// Erros
			ptst.TipoErro,
			ptst.SintaxeErro,
			ptst.AtributoErro,
			ptst.TipagemErro,
			ptst.NomeErro,
			ptst.ImportacaoErro,
			ptst.ValorErro,
			ptst.IndiceErro,
			ptst.RuntimeErro,
			ptst.FimIteracao,
			ptst.ErroDeAsseguracao,
			ptst.ConsultaErro,
			ptst.ChaveErro,
			ptst.SistemaErro,
			ptst.ArquivoNaoEncontradoErro,
		},
		constantes,
	)

	metodos := []*ptst.Metodo{
		_emb_imprima,
		_emb_leia,
		_emb_doc,
		_emb_int,
		_emb_texto,
		_emb_tamanho,
		_emb_instanciaDe,
		_emb_sequencia,
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
