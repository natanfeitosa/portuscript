package ptst

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Texto string

// FIXME: adicionar construtor
var TipoTexto = TipoObjeto.NewTipo(
	"Texto",
	`Texto(obj) -> Texto
Cria um novo objeto de texto para representar o objeto recebido.
Chama obj.__texto__() ou obj.__repr__(), se nenhum dos dois for encontrado, um erro pode ser lançado.
	`,
)

func NewTexto(arg any) (Objeto, error) {
	switch obj := arg.(type) {
	case nil:
		return Texto(""), nil
	case string:
		obj, err := strconv.Unquote(`"` + obj + `"`)
		if err != nil {
			return nil, err
		}

		return Texto(obj), nil
	case Texto:
		return obj, nil
	}

	if met, _ := ObtemAtributoS(arg.(Objeto), "__texto__"); met != nil {
		return met.(I__chame__).M__chame__(Tupla{})
	}

	if O, ok := arg.(I__texto__); ok {
		return O.M__texto__()
	}

	return nil, nil
}

func init() {
	TipoTexto.Nova = func(args Tupla) (Objeto, error) {
		return NewTexto(args[0])
	}
}

func (t Texto) Tipo() *Tipo {
	return TipoTexto
}

func (t Texto) M__texto__() (Objeto, error) {
	return t, nil
}

func (t Texto) M__bytes__() (Objeto, error) {
	return NewBytes(string(t))
}

func (t Texto) M__booleano__() (Objeto, error) {
	return NewBooleano(len(t) != 0)
}

func (t Texto) M__igual__(outro Objeto) (Objeto, error) {
	if !MesmoTipo(t, outro) {
		return Falso, nil
	}

	return NewBooleano(t == outro.(Texto))
}

// func (t Texto) M__ou__(outro Objeto) (Objeto, error) {}

func (t Texto) M__adiciona__(outro Objeto) (Objeto, error) {
	if !MesmoTipo(t, outro) {
		return nil, NewErroF(TipagemErro, "Não é possível concatenar o tipo '%s' com '%s'", t.Tipo().Nome, outro.Tipo().Nome)
	}

	outroTexto, err := NewTexto(outro)

	if err != nil {
		return nil, err
	}

	return Texto(fmt.Sprintf("%s%s", t, outroTexto.(Texto))), nil
}

func (t Texto) M__multiplica__(outro Objeto) (Objeto, error) {
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

func (t Texto) M__tamanho__() (Objeto, error) {
	return Inteiro(utf8.RuneCountInString(string(t))), nil
}

// func (t Texto) M__subtrai__(outro Objeto) (Objeto, error) {}

// func (t Texto) M__divide__(outro Objeto) (Objeto, error) {}

func (t Texto) String() string {
	return string(t)
}

// func (t Texto) ObtemMapa() Mapa {
// 	return t.Tipo().Mapa
// }

var _ I__texto__ = (*Texto)(nil)
var _ I__bytes__ = (*Texto)(nil)
var _ I__booleano__ = (*Texto)(nil)
var _ I__igual__ = (*Texto)(nil)

// var _ I__ou__ = (*Texto)(nil)
// var _ I__inteiro__ = (*Texto)(nil)
// var _ I__decimal__ = (*Texto)(nil)
var _ I__adiciona__ = (*Texto)(nil)
var _ I__multiplica__ = (*Texto)(nil)

// var _ I__subtrai__ = (*Texto)(nil)
// var _ I__divide__ = (*Texto)(nil)
// var _ I_Mapa = (*Texto)(nil)
var _ I__tamanho__ = (*Texto)(nil)

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

	TipoTexto.Mapa["titulo"] = NewMetodoOuPanic("titulo", func(inst Objeto) (Objeto, error) {
		titularizado := strings.Title(strings.ToLower(string(inst.(Texto))))
		return Texto(titularizado), nil
	}, "retorna uma cópia do texto com a primeira letra da frase em maiúsculo")

	TipoTexto.Mapa["maiusculas"] = NewMetodoOuPanic("maiusculas", func(inst Objeto) (Objeto, error) {
		return Texto(strings.ToUpper(string(inst.(Texto)))), nil
	}, "retorna uma cópia do texto com todas as letras em maiúsculas")

	TipoTexto.Mapa["minusculas"] = NewMetodoOuPanic("minusculas", func(inst Objeto) (Objeto, error) {
		return Texto(strings.ToLower(string(inst.(Texto)))), nil
	}, "retorna uma cópia do texto com todas as letras em minúsculas")
}
