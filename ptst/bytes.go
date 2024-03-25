package ptst

import (
	"bytes"
)

type Bytes struct {
	Itens     []byte
	Congelado bool
}

func NewBytes(arg any) (Objeto, error) {
	switch obj := arg.(type) {
	case nil:
		return &Bytes{make([]byte, 0), false}, nil
	case string:
		return &Bytes{[]byte(obj), false}, nil
	case *Bytes:
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

func (b *Bytes) Tipo() *Tipo {
	return TipoBytes
}

func (b *Bytes) M__diferente__(outro Objeto) (Objeto, error) {
	res, err := b.M__igual__(outro)
	if err != nil {
		return nil, err
	}

	return Booleano(!res.(Booleano)), nil
}

func (b *Bytes) M__igual__(outro Objeto) (Objeto, error) {
	if !MesmoTipo(b, outro) {
		return Falso, nil
	}

	return NewBooleano(bytes.Equal(b.Itens, outro.(*Bytes).Itens))
}

func (b *Bytes) M__maior_ou_igual__(outro Objeto) (Objeto, error) {
	outroT, err := Tamanho(outro)
	if err != nil {
		return nil, err
	}

	return NewBooleano(len(b.Itens) >= int(outroT.(Inteiro)))
}

func (b *Bytes) M__maior_que__(outro Objeto) (Objeto, error) {
	outroT, err := Tamanho(outro)
	if err != nil {
		return nil, err
	}

	return NewBooleano(len(b.Itens) > int(outroT.(Inteiro)))
}

func (b *Bytes) M__menor_ou_igual__(outro Objeto) (Objeto, error) {
	outroT, err := Tamanho(outro)
	if err != nil {
		return nil, err
	}

	return NewBooleano(len(b.Itens) <= int(outroT.(Inteiro)))
}

func (b *Bytes) M__menor_que__(outro Objeto) (Objeto, error) {
	outroT, err := Tamanho(outro)
	if err != nil {
		return nil, err
	}

	return NewBooleano(len(b.Itens) < int(outroT.(Inteiro)))
}

func (b *Bytes) M__tamanho__() (Objeto, error) {
	if b.Itens == nil {
		return Inteiro(0), nil
	}

	return NewInteiro(len(b.Itens))
}

func (b *Bytes) M__booleano__() (Objeto, error) {
	return NewBooleano(b.Itens != nil && len(b.Itens) > 0)
}

func (b *Bytes) M__decimal__() (Objeto, error) {
	panic("unimplemented")
}

func (b *Bytes) M__inteiro__() (Objeto, error) {
	panic("unimplemented")
}

func (b *Bytes) M__texto__() (Objeto, error) {
	return NewTexto(string(b.Itens))
}

var TipoBytes = TipoObjeto.NewTipo(
	"Bytes",
	"Bytes(obj) -> Bytes",
)

var _ I_comparacaoRica = (*Bytes)(nil)
var _ I__tamanho__ = (*Bytes)(nil)
var _ I_conversaoEntreTipos = (*Bytes)(nil)
