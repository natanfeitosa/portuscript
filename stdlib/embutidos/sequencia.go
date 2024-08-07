package embutidos

import (
	"github.com/natanfeitosa/portuscript/ptst"
)

type SequenciaNumerica struct {
	Inicio, Fim, Passo, Atual ptst.Inteiro
}

var TipoSequenciaNumerica = ptst.NewTipo("SequenciaNumerica", "Gerador de numeros inteiro com ordem crescente")

func (sn *SequenciaNumerica) Tipo() *ptst.Tipo {
	return TipoSequenciaNumerica
}

func (sn *SequenciaNumerica) M__iter__() (ptst.Objeto, error) {
	return sn, nil
}

func (sn *SequenciaNumerica) M__proximo__() (ptst.Objeto, error) {
	if sn.Passo > 0 && sn.Atual >= sn.Fim {
		return nil, ptst.NewErro(ptst.FimIteracao, ptst.Nulo)
	}

	if sn.Passo < 0 && sn.Atual <= sn.Fim {
		return nil, ptst.NewErro(ptst.FimIteracao, ptst.Nulo)
	}

	sn.Atual += sn.Passo
	return sn.Atual, nil
}

var _ ptst.I_iterador = (*SequenciaNumerica)(nil)

var met_emb_sequencia_doc = `sequencia(fim) -> SequenciaNumerica
sequencia(inicio, fim, passo?) -> SequenciaNumerica

Gera uma lista de números de [inicio] a [fim] (exclusivos), com incremento de [passo]`

func met_emb_sequencia(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("sequencia", false, args, 1, 3); err != nil {
		return nil, err
	}

	var inicio, fim, passo ptst.Objeto = ptst.Inteiro(0), ptst.Inteiro(1), ptst.Inteiro(1)
	var err error

	switch len(args) {
	case 3:
		if inicio, err = ptst.NewInteiro(args[0]); err != nil {
			return nil, err
		}

		if fim, err = ptst.NewInteiro(args[1]); err != nil {
			return nil, err
		}

		if passo, err = ptst.NewInteiro(args[2]); err != nil {
			return nil, err
		} else if passo.(ptst.Inteiro) == 0 {
			return nil, ptst.NewErroF(ptst.ValorErro, "O valor de passo da sequência deve ser diferente de zero")
		}

	case 2:
		if inicio, err = ptst.NewInteiro(args[0]); err != nil {
			return nil, err
		}

		if fim, err = ptst.NewInteiro(args[1]); err != nil {
			return nil, err
		}

	default:
		if fim, err = ptst.NewInteiro(args[1]); err != nil {
			return nil, err
		}
	}

	sn := &SequenciaNumerica{
		Inicio: inicio.(ptst.Inteiro),
		Fim:    fim.(ptst.Inteiro),
		Passo:  passo.(ptst.Inteiro),
		Atual:  0,
	}
	return sn, nil
}

var _emb_sequencia = ptst.NewMetodoOuPanic(
	"sequencia",
	met_emb_sequencia,
	met_emb_sequencia_doc,
)
