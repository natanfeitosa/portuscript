package ptst

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/compartilhado"
)

type Decimal float64

var TipoDecimal = TipoObjeto.NewTipo(
	"Decimal",
	`Decimal(obj) -> Decimal
Cria um novo objeto de decimal para representar o objeto recebido.
Chama obj.__decimal__() ou se esse não for encontrado, um erro pode ser lançado.
	`,
)

func NewDecimal(obj any) (Objeto, error) {
	switch b := obj.(type) {
	case nil:
		return Decimal(0), nil
	case Decimal:
		return b, nil
	case Texto:
		num, _ := compartilhado.StringParaDec(string(b))
		return Decimal(num), nil
	case string:
		num, _ := compartilhado.StringParaDec(string(b))
		return Decimal(num), nil
	default:
		if O, ok := b.(I__decimal__); ok {
			return O.O__decimal__()
		}

		// FIXME: isso está certo?
		return nil, nil
	}
}

func (d Decimal) Tipo() *Tipo {
	return TipoDecimal
}

func (d Decimal) O__texto__() (Objeto, error) {
	if i := int64(d); Decimal(i) == d {
		return Texto(fmt.Sprintf("%d.0", i)), nil
	}

	return Texto(fmt.Sprintf("%g", d)), nil
}

func (d Decimal) O__booleano__() (Objeto, error) {
	return NewBooleano(d != 0)
}

func (d Decimal) O__inteiro__() (Objeto, error) {
	return Inteiro(d), nil
}

func (d Decimal) O__decimal__() (Objeto, error) {
	return d, nil
}

func (d Decimal) O__adiciona__(outro Objeto) (Objeto, error) {
	outroInt, err := NewDecimal(outro)
	if err != nil {
		return nil, err
	}

	return d + outroInt.(Decimal), nil
}

func (d Decimal) O__multiplica__(outro Objeto) (Objeto, error) {
	outroInt, err := NewDecimal(outro)
	if err != nil {
		return nil, err
	}

	return d * outroInt.(Decimal), nil
}

func (d Decimal) O__subtrai__(outro Objeto) (Objeto, error) {
	outroInt, err := NewDecimal(outro)
	if err != nil {
		return nil, err
	}

	return outroInt.(Decimal) - d, nil
}

// FIXME: adicionar erro de divisão por zero
func (d Decimal) O__divide__(outro Objeto) (Objeto, error) {
	outroInt, err := NewDecimal(outro)
	if err != nil {
		return nil, err
	}

	return outroInt.(Decimal) - d, nil
}

// FIXME: adicionar erro de divisão por zero
func (d Decimal) O__divide_inteiro__(b Objeto) (Objeto, error) {
	bInt, err := NewInteiro(b)
	if err != nil {
		return nil, err
	}

	return Inteiro(d) / bInt.(Inteiro), nil
}

var _ conversaoEntreTipos = (*Decimal)(nil)
var _ aritmeticaMatematica = (*Decimal)(nil)