package ptst

func Adiciona(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__adiciona__); ok {
		res, err := A.M__adiciona__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '+' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func AdicionaEAtribui(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__adiciona_e_atribui__); ok {
		if res, err := A.M__adiciona_e_atribui__(b); err != nil {
			return nil, err
		} else if res != NaoImplementado {
			return res, nil
		}
	}

	return Adiciona(a, b)
}

func Multiplica(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__multiplica__); ok {
		res, err := A.M__multiplica__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '*' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func MultiplicaEAtribui(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__multiplica_e_atribui__); ok {
		if res, err := A.M__multiplica_e_atribui__(b); err != nil {
			return nil, err
		} else if res != NaoImplementado {
			return res, nil
		}
	}

	return Multiplica(a, b)
}

func Subtrai(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__subtrai__); ok {
		res, err := A.M__subtrai__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '-' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func SubtraiEAtribui(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__subtrai_e_atribui__); ok {
		if res, err := A.M__subtrai_e_atribui__(b); err != nil {
			return nil, err
		} else if res != NaoImplementado {
			return res, nil
		}
	}

	return Subtrai(a, b)
}

func Divide(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__divide__); ok {
		res, err := A.M__divide__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '/' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func DivideEAtribui(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__divide_e_atribui__); ok {
		if res, err := A.M__divide_e_atribui__(b); err != nil {
			return nil, err
		} else if res != NaoImplementado {
			return res, nil
		}
	}

	return Divide(a, b)
}

func DivideInteiro(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__divide_inteiro__); ok {
		res, err := A.M__divide_inteiro__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '//' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func DivideInteiroEAtribui(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__divide_inteiro_e_atribui__); ok {
		if res, err := A.M__divide_inteiro_e_atribui__(b); err != nil {
			return nil, err
		} else if res != NaoImplementado {

			return res, nil
		}
	}

	return DivideInteiro(a, b)
}

func Mod(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__mod__); ok {
		res, err := A.M__mod__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '%%' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func MenorQue(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__menor_que__); ok {
		res, err := A.M__menor_que__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '<' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func MenorOuIgual(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__menor_ou_igual__); ok {
		res, err := A.M__menor_ou_igual__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '<=' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func Igual(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__igual__); ok {
		res, err := A.M__igual__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	// if B, ok := b.(I__igual__); ok {
	// 	res, err := B.O__igual__(a)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return res, nil
	// }

	return nil, NewErroF(TipagemErro, "A operação '==' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func Diferente(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__diferente__); ok {
		res, err := A.M__diferente__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '!=' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func MaiorQue(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__maior_que__); ok {
		res, err := A.M__maior_que__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '>' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func MaiorOuIgual(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__maior_ou_igual__); ok {
		res, err := A.M__maior_ou_igual__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '>=' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func Ou(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__ou__); ok {
		res, err := A.M__ou__(b)

		if err != nil {
			return nil, err
		}

		if res != NaoImplementado {
			return res, nil
		}
	}

	return nil, NewErroF(TipagemErro, "A operação '|' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func E(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__e__); ok {
		res, err := A.M__e__(b)

		if err != nil {
			return nil, err
		}

		if res != NaoImplementado {
			return res, nil
		}
	}

	return nil, NewErroF(TipagemErro, "A operação '&' não é suportada entre os tipos '%s' e '%s'", a.Tipo().Nome, b.Tipo().Nome)
}

func Neg(a Objeto) (Objeto, error) {
	if A, ok := a.(I__neg__); ok {
		res, err := A.M__neg__()
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '-' não é suportada para o tipo '%s'", a.Tipo().Nome)
}

func Pos(a Objeto) (Objeto, error) {
	if A, ok := a.(I__pos__); ok {
		res, err := A.M__pos__()
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "A operação '+' não é suportada para o tipo '%s'", a.Tipo().Nome)
}
