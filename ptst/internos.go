package ptst

import (
	"fmt"
	"os"
	"reflect"
	"strings"
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
		return I.M__chame__(argsTupla)
	}

	return nil, NewErroF(TipagemErro, "O objeto '%s' não é do tipo chamável.", obj.Tipo().Nome)
}

func NomeAtributo(obj Objeto) (string, error) {
	if nome, ok := obj.(Texto); ok {
		return string(nome), nil
	}

	return "", NewErroF(AtributoErro, "O nome do atributo deve ser do tipo texto, não '%s'", obj.Tipo().Nome)
}

func ObtemAtributoS(inst Objeto, nome string) (Objeto, error) {
	if I, ok := inst.(I__obtem_attributo__); ok {
		return I.M__obtem_attributo__(nome)
	}

	// FIXME: e se o método for definido do "outro lado" do código?

	if len(nome) > 4 && (strings.HasPrefix(nome, "__") && strings.HasSuffix(nome, "__")) {
		ref := reflect.ValueOf(inst)
		m := ref.MethodByName("M" + nome)
		if m.IsValid() {
			return NewMetodoProxyDeNativo(nome, m.Interface())
		}
	}

	if I, ok := inst.(I_ObtemMapa); ok {
		mapa := I.ObtemMapa()
		
		if res, ok := mapa[nome]; ok {
			return res, nil
		}
	}

	tipo := inst.Tipo()
	if obj := tipo.G_ObtemAtributoOuNil(nome); obj != nil {
		if desc, ok := obj.(I__obtem__); ok {
			return desc.M__obtem__(inst, tipo)
		}

		return obj, nil
	}

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

func Proximo(obj Objeto) (Objeto, error) {
	if iter, ok := obj.(I__proximo__); ok {
		return iter.M__proximo__()
	}

	return nil, NewErroF(TipagemErro, "O objeto do tipo '%s' não implementa a interface do iterador", obj.Tipo().Nome)
}

func Iter(obj Objeto) (Objeto, error) {
	if iter, ok := obj.(I__iter__); ok {
		return iter.M__iter__()
	}

	return nil, NewErroF(TipagemErro, "O objeto do tipo '%s' não implementa a interface do iterador", obj.Tipo().Nome)
}

// FIXME: esta não é a melhor assinatura possível
func InstanciaDe(obj Objeto, tipos any) (Booleano, error) {
	switch tipo_tupla := tipos.(type) {
	case Tupla:
		for _, tipo := range tipo_tupla {
			if ok, err := InstanciaDe(obj, tipo); ok {
				return ok, nil
			} else if err != nil {
				return false, err
			}
		}

		return false, nil
	default:
		// FIXME: verificar se realmente é um tipo usável
		return obj.Tipo() == tipos.(*Tipo), nil
	}
}

func ObtemItem(inst, arg Objeto) (Objeto, error) {
	if I, ok := inst.(I__obtem_item__); ok {
		return I.M__obtem_item__(arg)
	}
	
	return nil, NewErroF(TipagemErro, "O tipo '%s' não suporta o uso de indices", inst.Tipo().Nome)
}

func DefineItem(inst, chave, valor Objeto) (Objeto, error) {
	if I, ok := inst.(I__define_item__); ok {
		return I.M__define_item__(chave, valor)
	}
	
	return nil, NewErroF(TipagemErro, "O tipo '%s' não suporta a atribuição por indice", inst.Tipo().Nome)
}

func NovaInstancia(obj Objeto, args Tupla) (Objeto, error) {
	if I, ok := obj.(I__nova_instancia__); ok {
		return I.M__nova_instancia__(obj.(*Tipo), args)
	}

	// t := obj.(*Tipo)
	// if t.Base != nil {
	// 	return t.M__nova_instancia__(t, args)
	// }

	return nil, NewErroF(TipagemErro, "O objeto '%s' não é instanciável", obj.Tipo().Nome)
}
