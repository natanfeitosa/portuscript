package ptst

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/compartilhado"
)

type Inteiro int64

var TipoInteiro = TipoObjeto.NewTipo(
	"Inteiro",
	`Inteiro(obj) -> Inteiro
Cria um novo objeto de inteiro para representar o objeto recebido.
Chama obj.__inteiro__() ou se esse não for encontrado, um erro pode ser lançado.
	`,
)

func NewInteiro(obj any) (Objeto, error) {
	switch b := obj.(type) {
	case nil:
		return Inteiro(0), nil
	case int:
		return Inteiro(b), nil
	case Inteiro:
		return b, nil
	case Texto:
		num, _ := compartilhado.StringParaInt(string(b))
		return Inteiro(num), nil
	case string:
		num, _ := compartilhado.StringParaInt(string(b))
		return Inteiro(num), nil
	default:
		if O, ok := b.(I__inteiro__); ok {
			return O.M__inteiro__()
		}

		return nil, NewErroF(TipagemErro, "O argumento do construtor do tipo Inteiro deve ser uma string, Texto ou um outro tipo que implemente o método __inteiro__()")
	}
}

func init() {
	TipoInteiro.Nova = func(args Tupla) (Objeto, error) {
		return NewInteiro(args[0])
	}
}

func (i Inteiro) Tipo() *Tipo {
	return TipoInteiro
}

// func (t Inteiro) ObtemMapa() Mapa {
// 	return t.Tipo().Mapa
// }

func (i Inteiro) M__texto__() (Objeto, error) {
	return Texto(fmt.Sprintf("%d", i)), nil
}

func (i Inteiro) M__booleano__() (Objeto, error) {
	return NewBooleano(i != 0)
}

func (i Inteiro) M__inteiro__() (Objeto, error) {
	return i, nil
}

func (i Inteiro) M__decimal__() (Objeto, error) {
	return Decimal(i), nil
}

func (i Inteiro) M__adiciona__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i + bInt.(Inteiro), nil
}

func (i Inteiro) M__multiplica__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i * bInt.(Inteiro), nil
}

func (i Inteiro) M__subtrai__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i - bInt.(Inteiro), nil
}

// FIXME: adicionar erro de divisão por zero
func (i Inteiro) M__divide__(b Objeto) (Objeto, error) {
	bInt, err := NewDecimal(b)
	if err != nil {
		return nil, err
	}

	return Decimal(i) / bInt.(Decimal), nil
}

// FIXME: adicionar erro de divisão por zero
func (i Inteiro) M__divide_inteiro__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i / bInt.(Inteiro), nil
}

func (i Inteiro) M__mod__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return i % bInt.(Inteiro), nil
}

func (i Inteiro) M__neg__() (Objeto, error) {
	return -i, nil
}

func (i Inteiro) M__pos__() (Objeto, error) {
	return +i, nil
}

func (i Inteiro) M__menor_que__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '<' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i < b.(Inteiro))
}

func (i Inteiro) M__menor_ou_igual__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '<=' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i <= b.(Inteiro))
}

func (i Inteiro) M__igual__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return Falso, nil
	}

	return NewBooleano(i == b.(Inteiro))
}

func (i Inteiro) M__diferente__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return Verdadeiro, nil
	}

	return NewBooleano(i != b.(Inteiro))
}

func (i Inteiro) M__maior_que__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '>' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i > b.(Inteiro))
}

func (i Inteiro) M__maior_ou_igual__(b Objeto) (Objeto, error) {
	if !MesmoTipo(i, b) {
		return nil, NewErroF(TipagemErro, "A operação '>=' não é suportada entre os tipos '%s' e '%s'", i.Tipo().Nome, b.Tipo().Nome)
	}

	return NewBooleano(i >= b.(Inteiro))
}

func (i Inteiro) M__ou__(b Objeto) (Objeto, error) {
	if MesmoTipo(i, b) {
		return NewInteiro(i | b.(Inteiro))
	}

	return NaoImplementado, nil
}

func (i Inteiro) M__e__(b Objeto) (Objeto, error) {
	if MesmoTipo(i, b) {
		return NewInteiro(i & b.(Inteiro))
	}

	return NaoImplementado, nil
}

var _ I_conversaoEntreTipos = (*Inteiro)(nil)
var _ I_aritmeticaMatematica = (*Inteiro)(nil)
var _ I_comparacaoRica = (*Inteiro)(nil)
var _ I_aritmeticaBooleana = (*Inteiro)(nil)
