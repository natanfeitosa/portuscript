package ptst

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

func Em(a, b Objeto) (Objeto, error) {
	if A, ok := a.(I__contem__); ok {
		res, err := A.M__contem__(b)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, NewErroF(TipagemErro, "O tipo '%s' não aceita a operação 'em'", b.Tipo().Nome)
}
