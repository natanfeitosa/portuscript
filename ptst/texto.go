package ptst

import "fmt"

type Texto string

// FIXME: adicionar construtor
var TipoTexto = TipoObjeto.NewTipo(
	"Texto",
	`Texto(obj) -> Texto
Cria um novo objeto de texto para representar o objeto recebido.
Chama obj.__texto__() ou obj.__repr__(), se nenhum dos dois for encontrado, um erro pode ser lançado.
	`,
)

func NewTexto(obj Objeto) (Objeto, error) {
	switch obj.(type) {
	case Texto:
		return obj, nil
	case nil:
		return Texto(""), nil
	default:
		if O, ok := obj.(I__texto__); ok {
			return O.O__texto__()
		}
	}

	// FIXME: ?????
	return nil, nil
}

func (t Texto) Tipo() *Tipo {
	return TipoTexto
}

func (t Texto) O__texto__() (Objeto, error) {
	return t, nil
}

func (t Texto) O__booleano__() (Objeto, error) {
	return NewBooleano(len(t) != 0)
}

func (t Texto) O__igual__(outro Objeto) (Objeto, error) {
	if !MesmoTipo(t, outro) {
		return Falso, nil
	}

	return NewBooleano(t == outro.(Texto))
}

// func (t Texto) O__ou__(outro Objeto) (Objeto, error) {}

func (t Texto) O__adiciona__(outro Objeto) (Objeto, error) {
	if !MesmoTipo(t, outro) {
		return nil, NewErroF(TipagemErro, "Não é possível concatenar o tipo '%s' com '%s'", t.Tipo().Nome, outro.Tipo().Nome)
	}

	outroTexto, err := NewTexto(outro)

	if err != nil {
		return nil, err
	}

	return Texto(fmt.Sprintf("%s%s", t, outroTexto.(Texto))), nil
}

func (t Texto) O__multiplica__(outro Objeto) (Objeto, error) {
	switch obj := outro.(type) {
	case Inteiro:
		resultado := Texto(t)

		for i := 0; i < int(obj); i++ {
			resultado += t
		}

		return resultado, nil
	default:
		return nil, NewErroF(TipagemErro, "A operação '*' não é suportada entre os tipos '%s' e '%s'", t.Tipo().Nome, obj.Tipo().Nome)
	}
}

// func (t Texto) O__subtrai__(outro Objeto) (Objeto, error) {}

// func (t Texto) O__divide__(outro Objeto) (Objeto, error) {}

func (t Texto) String() string {
	return string(t)
}

func (t Texto) Mapa() Mapa {
	return t.Tipo().Mapa
}

var _ I__texto__ = (*Texto)(nil)
var _ I__booleano__ = (*Texto)(nil)
var _ I__igual__ = (*Texto)(nil)
// var _ I__ou__ = (*Texto)(nil)
// var _ I__inteiro__ = (*Texto)(nil)
// var _ I__decimal__ = (*Texto)(nil)
var _ I__adiciona__ = (*Texto)(nil)
var _ I__multiplica__ = (*Texto)(nil)
// var _ I__subtrai__ = (*Texto)(nil)
// var _ I__divide__ = (*Texto)(nil)

func init() {
	TipoTexto.Mapa["junta"] = NewMetodoOuPanic("junta", func(inst Objeto, iter Objeto) (Objeto, error) {
		saida := ""

		for i, arg := range iter.(Tupla) {
			texto, err := NewTexto(arg)
			if err != nil {
				return nil, err
			}

			saida += string(texto.(Texto))
			if i != len(iter.(Tupla))-1 {
				saida += string(inst.(Texto))
			}
		}

		return Texto(saida), nil
	}, `concatena o iterável recebido com o texto da instancia`)
}