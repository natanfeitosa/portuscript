package ptst

import "bytes"

type Mapa map[string]Objeto

var TipoMapa = NewTipo(
	"Mapa",
	"Objeto chave/valor",
)

func NewMapaVazio() Mapa {
	return make(Mapa)
}

func (m Mapa) Tipo() *Tipo {
	return TipoMapa
}

func (m Mapa) M__texto__() (Objeto, error) {
	var out bytes.Buffer
	out.WriteString("{ ")
	separar := false

	for chave, valor := range m {
		if separar {
			out.WriteString(", ")
		}

		var chaveT, valorT Objeto
		var err error

		if chaveT, err = NewTexto(chave); err != nil {
			return nil, err
		}

		if valorT, err = NewTexto(valor); err != nil {
			return nil, err
		}

		out.WriteString(string(chaveT.(Texto)))
		out.WriteString(": ")
		out.WriteString(string(valorT.(Texto)))
		separar = true
	}

	out.WriteString(" }")
	return NewTexto(out.String())
}

func (m Mapa) M__iter__() (Objeto, error) {
	entradas := make(Tupla, 0, len(m))

	for chave, valor := range m {
		entradas = append(entradas, Tupla{Texto(chave), valor})
	}

	return NewIterador(entradas)
}

func (m Mapa) M__tamanho__() (Objeto, error) {
	return NewInteiro(len(m))
}

func (m Mapa) M__obtem_item__(obj Objeto) (Objeto, error) {
	chave, ok := obj.(Texto)

	if !ok {
		return nil, NewErroF(ChaveErro, "A chave para um '%s' deve ser do tipo '%s' e não '%s'", TipoMapa.Nome, TipoTexto.Nome, obj.Tipo().Nome)
	}

	if valor, ok := m[string(chave)]; ok {
		return valor, nil
	}

	return nil, NewErroF(ChaveErro, "O Mapa não tem um elemento com a chave '%s'", chave)
}

func (m Mapa) M__define_item__(obj, valor Objeto) (Objeto, error) {
	chave, ok := obj.(Texto)

	if !ok {
		return nil, NewErroF(ChaveErro, "A chave para um '%s' deve ser do tipo '%s' e não '%s'", TipoMapa.Nome, TipoTexto.Nome, obj.Tipo().Nome)
	}

	m[string(chave)] = valor
	return nil, nil
}

var _ I__iter__ = (*Mapa)(nil)
var _ I__texto__ = (*Mapa)(nil)
var _ I__tamanho__ = (*Mapa)(nil)
var _ I__obtem_item__ = (*Mapa)(nil)
var _ I__define_item__ = (*Mapa)(nil)

func (m Mapa) Chaves() (Tupla, error) {
	if len(m) == 0 {
		return Tupla(nil), nil
	}

	chaves := make(Tupla, 0)

	for chave := range m {
		chaves = append(chaves, Texto(chave))
	}

	return chaves, nil
}

func (m Mapa) Valores() (Tupla, error) {
	if len(m) == 0 {
		return Tupla(nil), nil
	}

	valores := make(Tupla, 0)

	for _, valor := range m {
		valores = append(valores, valor)
	}

	return valores, nil
}

func (m Mapa) Atualizar(outro Mapa, ignoreExistentes Booleano) (Mapa, error) {
	for c, v := range outro {
		if ignoreExistentes {
			if _, existe := m[c]; existe {
				continue
			}
		}

		m[c] = v
	}

	return m, nil
}

func init() {
	TipoMapa.Mapa["chaves"] = NewMetodoOuPanic("chaves", func(inst Objeto) (Objeto, error) {
		return inst.(Mapa).Chaves()
	}, `Retorna uma tupla contendo todos as chaves do mapa`)

	TipoMapa.Mapa["valores"] = NewMetodoOuPanic("valores", func(inst Objeto) (Objeto, error) {
		return inst.(Mapa).Valores()
	}, `Retorna uma tupla contendo todos os valores do mapa`)

	TipoMapa.Mapa["atualizar"] = NewMetodoOuPanic("atualizar", func(inst Objeto, args Tupla) (Objeto, error) {
		if err := VerificaNumeroArgumentos("atualizar", true, args, 1, 2); err != nil {
			return nil, err
		}

		ignoreExistentes := Falso
		novoMapa := args[0]

		if _, ok := novoMapa.(Mapa); !ok {
			return nil, NewErroF(TipagemErro, "Era esperado o tipo 'Mapa', mas ao invés disso foi recebido o tipo '%s'", novoMapa.Tipo().Nome)
		}

		if len(args) == 2 {
			if ignore, ok := args[1].(Booleano); ok {
				ignoreExistentes = ignore
			} else if !ok {
				return nil, NewErroF(TipagemErro, "Era esperado o tipo 'Booleano', mas ao invés disso foi recebido o tipo '%s'", args[1].Tipo().Nome)
			}
		}

		return inst.(Mapa).Atualizar(novoMapa.(Mapa), ignoreExistentes)
	}, `mapa.atualizar(outroMapa, ignoreExistentes?) -> Mapa
Atualiza o mapa atual com as chaves/valores do outro, se o parâmetro ignoreExistentes for Verdadeiro,
as chaves que se repetem serão mantidas com o valor atual do mapa que chamou o método.`)
}
