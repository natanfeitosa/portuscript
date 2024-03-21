package ptst

import "fmt"

type Metodo struct {
	Nome     string
	Doc      string
	Modulo   *Modulo
	chamavel interface{}
}

var TipoMetodo = NewTipo("Metodo", "Um metodo Portuscript")

func (o *Metodo) Tipo() *Tipo {
	return TipoMetodo
}

func (m *Metodo) Chamar(inst Objeto, args Tupla) (Objeto, error) {
	switch f := m.chamavel.(type) {
	case func(inst Objeto, args Tupla) (Objeto, error):
		return f(inst, args)
	// case func(inst Objeto, args Tupla, kwargs StringDict) (Objeto, error):
	// 	return f(inst, args, NewStringDict())
	case func(Objeto) (Objeto, error):
		if len(args) != 0 {
			return nil, NewErroF(TipagemErro, "%s() não aceita argumentos, %d foram recebidos", m.Nome, len(args))
		}
		return f(inst)
	case func(Objeto, Objeto) (Objeto, error):
		// if len(args) != 1 {
		// 	return nil, NewErroF(TipagemErro, "%s() esperava receber 1 argumento, %d foram recebidos", m.Nome, len(args))
		// }
		// return f(inst, args[0])
		return f(inst, Objeto(args))
	}

	panic(fmt.Sprintf("Tipo de método desconhecido: %T", m.chamavel))
}

func (m *Metodo) ObtemDoc() string {
	return m.Doc
}

func (m *Metodo) M__chame__(args Tupla) (Objeto, error) {
	return m.Chamar(Objeto(m.Modulo), args)
}

// FIXME: isso deve retornar um proxy
func (m *Metodo) M__obtem__(inst Objeto, dono *Tipo) (Objeto, error) {
	if inst != Nulo {
		return NewMetodoProxy(inst, m), nil
	}

	return m, nil
}

var _ I__chame__ = (*Metodo)(nil)
var _ I_Chamar = (*Metodo)(nil)
var _ I_ObtemDoc = (*Metodo)(nil)
var _ I__obtem__ = (*Metodo)(nil)

// Copiado de https://github.com/go-python/gpython/blob/main/py/method.go#L97C1-L115C2
func NewMetodo(nome string, chamavel interface{}, doc string) (*Metodo, error) {
	return &Metodo{
		Nome:     nome,
		Doc:      doc,
		chamavel: chamavel,
	}, nil
}

func NewMetodoOuPanic(nome string, chamavel interface{}, doc string) *Metodo {
	m, err := NewMetodo(nome, chamavel, doc)

	if err != nil {
		panic(err)
	}

	return m
}

func NewMetodoProxyDeNativo(nome string, chamavel interface{}) (*Metodo, error) {
	metodo := &Metodo{
		Nome: nome,
	}

	switch fn := chamavel.(type) {
	// Métodos como `__iter__`, `__tamanho__` e `__text__`
	case func() (Objeto, error):
		metodo.chamavel = func(_ Objeto) (Objeto, error) {
			return fn()
		}
	case func(Objeto) (Objeto, error):
		metodo.chamavel = func(_ Objeto, arg Objeto) (Objeto, error) {
			return fn(arg)
		}
	// `__nova_instancia__`
	case func(*Tipo, Tupla) (Objeto, error):
		metodo.chamavel = func(inst Objeto, args Tupla) (Objeto, error) {
			meta := args[0].(*Tipo)
			args = args[0:]

			return fn(meta, args)
		}
	default:
		// FIXME: adicionar um tipo de erro adequado
		return nil, fmt.Errorf("não foi possível criar um proxy para o método %T", fn)
	}

	return metodo, nil
}
