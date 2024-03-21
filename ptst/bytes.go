package ptst

type Bytes struct {
	Itens []byte
	Congelado bool
}

var TipoBytes = TipoObjeto.NewTipo(
	"Bytes",
	"Bytes(obj) -> Bytes",
)

func NewBytes(arg any) (Objeto, error) {
	switch obj := arg.(type) {
	case nil:
		return &Bytes{make([]byte, 0), false}, nil
	case string:
		return &Bytes{[]byte(obj), false}, nil
	case Bytes:
		return obj, nil
	}

	if met, _ := ObtemAtributoS(arg.(Objeto), "__bytes__"); met != nil {
		return Chamar(met, Tupla{})
	}

	if O, ok := arg.(I__bytes__); ok {
		return O.M__bytes__()
	}

	return nil, nil
}

func init() {
	TipoBytes.Nova = func(args Tupla) (Objeto, error) {
		return NewBytes(args[0])
	}
}

func (t Bytes) Tipo() *Tipo {
	return TipoBytes
}