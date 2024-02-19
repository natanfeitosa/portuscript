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
	junta, err := ObtemItemS(Texto(","), "junta")
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

var _ I__iter__ = Tupla(nil)
var _ I__texto__ = Tupla(nil)
var _ I__tamanho__ = Tupla(nil)