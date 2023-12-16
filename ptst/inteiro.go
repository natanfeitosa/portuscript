package ptst

import "fmt"

type Inteiro int64

var TipoInteiro = TipoObjeto.NewTipo(
	"Inteiro",
	`Inteiro(obj) -> Inteiro
Cria um novo objeto de inteiro para representar o objeto recebido.
Chama obj.__inteiro__() ou se esse não for encontrado, um erro pode ser lançado.
	`,
)

func NewInteiro(obj Objeto) (Objeto, error) {
	switch b := obj.(type) {
	case Inteiro:
		return b, nil
	case nil:
		return Inteiro(0), nil
	default:
		if O, ok := b.(I__inteiro__); ok {
			return O.O__inteiro__()
		}
	}

	// FIXME: ?????
	return nil, nil
}

func (i Inteiro) Tipo() *Tipo {
	return TipoInteiro
}

func (i Inteiro) O__texto__() (Objeto, error) {
	return Texto(fmt.Sprintf("%d", i)), nil
}

func (i Inteiro) O__booleano__() (Objeto, error) {
	return NewBooleano(i != 0)
}

func (i Inteiro) O__inteiro__() (Objeto, error) {
	return i, nil
}

func (i Inteiro) O__decimal__() (Objeto, error) {
	return Decimal(i), nil
}

func (i Inteiro) O__adiciona__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i + bInt.(Inteiro), nil
}

func (i Inteiro) O__multiplica__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i * bInt.(Inteiro), nil
}

func (i Inteiro) O__subtrai__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i - bInt.(Inteiro), nil
}

// FIXME: adicionar erro de divisão por zero
func (i Inteiro) O__divide__(b Objeto) (Objeto, error) {
	bInt, err := NewDecimal(b)
	if err != nil {
		return nil, err
	}

	return Decimal(i) / bInt.(Decimal), nil
}

func (i Inteiro) O__menor_que__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '<' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i < b.(Inteiro))
}

func (i Inteiro) O__menor_ou_igual__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '<=' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i <= b.(Inteiro))
}

func (i Inteiro) O__igual__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return Falso, nil
	}

	return NewBooleano(i == b.(Inteiro))
}

func (i Inteiro) O__diferente__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return Verdadeiro, nil
	}

	return NewBooleano(i != b.(Inteiro))
}

func (i Inteiro) O__maior_que__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '>' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i > b.(Inteiro))
}

func (i Inteiro) O__maior_ou_igual__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '>=' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i >= b.(Inteiro))
}

func (i Inteiro) O__ou__(b Objeto) (Objeto, error) {
	booleano, err := i.O__booleano__()
	if err != nil {
		return nil, err
	}

	if booleano.(Booleano) {
		return i, nil
	}

	return b, nil
}

func (i Inteiro) O__e__(b Objeto) (Objeto, error) {
	booleano, err := i.O__booleano__()
	if err != nil {
		return nil, err
	}

	if booleano.(Booleano) {
		return b, nil
	}

	return i, nil
}

var _ conversaoEntreTipos = (*Inteiro)(nil)
var _ aritmeticaMatematica = (*Inteiro)(nil)
var _ comparacaoRica = (*Inteiro)(nil)
var _ aritmeticaBooleana = (*Inteiro)(nil)
