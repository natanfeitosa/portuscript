package ptst

type Iterador struct {
	Posicao int
	Conteiner Objeto
}

var TipoIterador = NewTipo("Iterador", "Objeto abstrato que representa um iterador nativo")

func NewIterador(seq Objeto) (*Iterador, error) {
	return &Iterador{Posicao: 0, Conteiner: seq}, nil
}

func (it *Iterador) Tipo() *Tipo {
	return TipoIterador
}

func (it *Iterador) M__iter__() (Objeto, error) {
	return it, nil
}

func (it *Iterador) M__proximo__() (Objeto, error) {
	if tupla, ok := it.Conteiner.(Tupla); ok {
		if it.Posicao >= len(tupla) {
			return nil, NewErro(FimIteracao, Nulo)
		}
		item := tupla[it.Posicao]
		it.Posicao += 1

		return item, nil
	}

// FIXME: implementar outros casos
	return nil, nil
}

var _ I_iterador = (*Iterador)(nil)