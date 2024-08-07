package ptst

type Booleano bool

var (
	TipoBooleano = NewTipo(
		"Booleano",
		"Verdadeiro ou Falso",
	)
	Verdadeiro = Booleano(true)
	Falso      = Booleano(false)
)

func boolParaBooleano(obj bool) Booleano {
	if obj {
		return Verdadeiro
	}

	return Falso
}

func NewBooleano(obj any) (Objeto, error) {
	switch b := obj.(type) {
	case I__booleano__:
		return b.M__booleano__()
	case bool:
		return boolParaBooleano(b), nil
	}

	return Falso, nil
}

func init() {
	TipoBooleano.Nova = func(args Tupla) (Objeto, error) {
		return NewBooleano(args[0])
	}
}

func (b Booleano) Tipo() *Tipo {
	return TipoBooleano
}

func (b Booleano) M__texto__() (Objeto, error) {
	if b {
		return Texto("Verdadeiro"), nil
	}

	return Texto("Falso"), nil
}

func (b Booleano) M__booleano__() (Objeto, error) {
	return b, nil
}

func (b Booleano) M__igual__(a Objeto) (Objeto, error) {
	if !MesmoTipo(b, a) {
		return Falso, nil
	}

	return NewBooleano(b == a.(Booleano))
}

func (b Booleano) M__diferente__(a Objeto) (Objeto, error) {
	if !MesmoTipo(b, a) {
		// FIXME: não deveria retornar Verdadeiro?
		return Falso, nil
	}

	igual, err := b.M__igual__(a)
	if err != nil {
		return nil, err
	}

	return Nao(igual)
}

func (b Booleano) M__decimal__() (Objeto, error) {
	if b == Verdadeiro {
		return Decimal(1.0), nil
	}

	return Decimal(0), nil
}

func (b Booleano) M__inteiro__() (Objeto, error) {
	if b == Verdadeiro {
		return Inteiro(1), nil
	}

	return Inteiro(0), nil
}

func (b Booleano) M__ou__(a Objeto) (Objeto, error) {
	if MesmoTipo(b, a) {
		var bi, ai Objeto
		var err error

		if bi, err = NewInteiro(b); err != nil {
			return nil, err
		}

		if ai, err = NewInteiro(a); err != nil {
			return nil, err
		}

		return NewBooleano(bi.(Inteiro) | ai.(Inteiro))
	}

	return NaoImplementado, nil
}

func (b Booleano) M__e__(a Objeto) (Objeto, error) {
	if MesmoTipo(b, a) {
		var bi, ai Objeto
		var err error

		if bi, err = NewInteiro(b); err != nil {
			return nil, err
		}

		if ai, err = NewInteiro(a); err != nil {
			return nil, err
		}

		return NewBooleano(bi.(Inteiro) & ai.(Inteiro))
	}

	return NaoImplementado, nil
}

var _ I_conversaoEntreTipos = (*Booleano)(nil)
var _ I__igual__ = (*Booleano)(nil)
var _ I__diferente__ = (*Booleano)(nil)
var _ I_aritmeticaBooleana = (*Booleano)(nil)
