package ptst

type _Nulo struct {}

var Nulo = _Nulo(struct{}{})

var TipoNulo = TipoObjeto.NewTipo("Nulo", "Tipo que referencia a algo sem valor definido")

func (n _Nulo) Tipo() *Tipo {
	return TipoNulo
}
