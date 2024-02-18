package ptst

import (
	"github.com/natanfeitosa/portuscript/parser"
)

type Funcao struct {
	Nome     string        // Disponível em `funcao.__nome__`
	Doc      Texto         // Disponível em `funcao.__doc__`
	args     Tupla         // Uma tupla com os nomes na mesma ordem. Isso pode ser melhorado, não?
	corpo    *parser.Bloco // A ast toda?
	contexto *Contexto     // Contexto pai do contexto interno (confuso, né? Arrumaremos depois)
	escopo   *Escopo
}

var TipoFuncao = NewTipo("Funcao", "Uma funcao Portuscript")

func (f *Funcao) Tipo() *Tipo {
	return TipoFuncao
}

func NewFuncao(nome string, corpo *parser.Bloco, contexto *Contexto, escopo *Escopo) *Funcao {
	return &Funcao{Nome: nome, corpo: corpo, contexto: contexto, escopo: escopo}
}

func (f *Funcao) O__chame__(args Tupla) (Objeto, error) {
	if len(f.args) != len(args) {
		// FIXME: este erro é adequado? A mensagem também não deveria se atentar ao plural/singular quando aplicável?
		return nil, NewErroF(TipagemErro, "%v() esperava receber %v argumentos, mas %v foram encontrados", f.Nome, len(f.args), len(args))
	}

	escopo := f.escopo.NewEscopo()

	for i, nome := range f.args {
		nomeStr, _ := NewTexto(nome)
		escopo.DefinirSimbolo(
			NewVarSimbolo(string(nomeStr.(Texto)), args[i]),
		)
	}

	return (&Interpretador{Ast: f.corpo, Contexto: f.contexto, Escopo: escopo}).Inicializa()
}

var _ I__chame__ = (*Funcao)(nil)
