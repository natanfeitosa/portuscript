package ptst

type Tupla []Objeto

var TipoTupla = TipoObjeto.NewTipo(
	"Tupla",
	"Tupla(obj) -> Tupla",
)

func (t Tupla) Tipo() *Tipo {
	return TipoTupla
}

func (t Tupla) GRepr(inicio, fim string) (Objeto, error) {
	junta, err := ObtemAtributoS(Texto(","), "junta")
	if err != nil {
		return nil, err
	}

	res, err := Chamar(junta, t)
	if err != nil {
		return nil, err
	}

	return (Texto(inicio) + res.(Texto) + Texto(fim)), nil
}

func (t Tupla) M__iter__() (Objeto, error) {
	return NewIterador(t)
}

func (t Tupla) M__texto__() (Objeto, error) {
	return t.GRepr("(", ")")
}

func (t Tupla) M__tamanho__() (Objeto, error) {
	return Inteiro(len(t)), nil
}

func (t Tupla) ObtemItem(i Objeto, nomeTipo string) (Objeto, error) {
	if I, ok := i.(Inteiro); ok {
		return t[I], nil
	}

	return nil, NewErroF(TipagemErro, "O tipo '%s' não é aceito para indexação no tipo '%s'. Use um 'Inteiro'.", i.Tipo().Nome, nomeTipo)
}

func (t Tupla) DefineItem(chave, valor Objeto, nomeTipo string) (Objeto, error) {
	if I, ok := chave.(Inteiro); ok {
		t[I] = valor
		return t, nil
	}

	return nil, NewErroF(TipagemErro, "O tipo '%s' não é aceito para indexação no tipo '%s'. Use um 'Inteiro'.", chave.Tipo().Nome, nomeTipo)
}

func (t Tupla) M__obtem_item__(obj Objeto) (Objeto, error) {
	return t.ObtemItem(obj, t.Tipo().Nome)
}

func (t Tupla) M__define_item__(chave, valor Objeto) (Objeto, error) {
	return t.DefineItem(chave, valor, t.Tipo().Nome)
}

var _ I__iter__ = Tupla(nil)
var _ I__texto__ = Tupla(nil)
var _ I__tamanho__ = Tupla(nil)
var _ I__obtem_item__ = Tupla(nil)
var _ I__define_item__ = Tupla(nil)
