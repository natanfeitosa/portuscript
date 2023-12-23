package ptst

type MetodoProxy struct {
	Inst   Objeto // a instancia
	Metodo Objeto
}

var TipoMetodoProxy = NewTipo("MetodoProxy", "Um método que não exatamente reflete o método real")

func newMetodoProxy(inst, metodo Objeto) *MetodoProxy {
	return &MetodoProxy{inst, metodo}
}

func (mp *MetodoProxy) Tipo() *Tipo {
	return TipoMetodoProxy
}

func (mp *MetodoProxy) O__chame__(args Tupla) (Objeto, error) {
	if m, ok := mp.Metodo.(*Metodo); ok {
		return m.Chamar(mp.Inst, args)
	}

	// FIXME: o que fazer com o `inst`?
	return Chamar(mp.Metodo, args)
}

func (mp *MetodoProxy) ObtemDoc() string {
	return mp.Metodo.(*Metodo).ObtemDoc()
}

var _ I__chame__ = (*MetodoProxy)(nil)
var _ I_ObtemDoc = (*MetodoProxy)(nil)