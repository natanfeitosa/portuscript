package ptst

import (
	"fmt"
	"os"
)

func LancarErro(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func adicionaContextoSeNaoTiver(err error, context *Contexto) {
	if err == nil {
		return
	}

	if err.(*Erro).Contexto == nil {
		err.(*Erro).Contexto = context
	}
}

func Chamar(obj Objeto, args Objeto) (Objeto, error) {
	// if I, ok := obj.(I_Chamar); ok {
	// 	return I.Chamar(obj, args.(Tupla))
	// }
	var argsTupla Tupla

	if t, ok := args.(Tupla); ok {
		argsTupla = t
	} else {
		argsTupla = Tupla{args}
	}

	if I, ok := obj.(I__chame__); ok {
		return I.O__chame__(argsTupla)
	}

	return nil, NewErroF(TipagemErro, "O objeto '%s' não é do tipo chamável.", obj.Tipo().Nome)
}

func NomeAtributo(obj Objeto) (string, error) {
	if nome, ok := obj.(Texto); ok {
		return string(nome), nil
	}

	return "", NewErroF(AtributoErro, "O nome do atributo deve ser do tipo texto, não '%s'", obj.Tipo().Nome)
}

func ObtemItemS(inst Objeto, nome string) (Objeto, error) {
	if I, ok := inst.(I__obtem_attributo__); ok {
		return I.O__obtem_attributo__(nome)
	}

	// FIXME: e se o método for definido do "outro lado" do código?
	// FIXME: adicionar um proxy para métodos nativos

	if I, ok := inst.(I_Mapa); ok {
		mapa := I.ObtemMapa()
		
		if res, ok := mapa[Texto(nome)]; ok {
			return res, nil
		}
	}

	// FIXME: adicionar a capacidade de chamar o método __obtem_item__ e __obtem__
	return nil, NewErroF(AtributoErro, "O atributo '%s' não existe no tipo '%s'", nome, inst.Tipo().Nome)
}

func Nao(obj Objeto) (Objeto, error) {
	booleano, err := NewBooleano(obj)
	if err != nil {
		return nil, err
	}

	switch booleano.(Booleano) {
	case Falso:
		return Verdadeiro, nil
	case Verdadeiro:
		return Falso, nil
	}

	// FIXME
	return nil, nil
}