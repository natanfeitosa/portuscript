package ptst

type Lista struct {
	Itens Tupla
}

var TipoLista = TipoObjeto.NewTipo(
	"Lista",
	"Lista(obj) -> Lista",
)

func (l *Lista) Tipo() *Tipo {
	return TipoLista
}

func (l *Lista) M__iter__() (Objeto, error) {
	return NewIterador(l.Itens)
}

func (l *Lista) M__texto__() (Objeto, error) {
	return l.Itens.GRepr("[", "]")
}

func (l *Lista) M__tamanho__() (Objeto, error) {
	return l.Itens.M__tamanho__()
}

func (l *Lista) M__obtem_item__(obj Objeto) (Objeto, error) {
	return l.Itens.ObtemItem(obj, "Lista")
}

func (l *Lista) M__define_item__(chave, valor Objeto) (Objeto, error) {
	if _, err := l.Itens.DefineItem(chave, valor, l.Tipo().Nome); err != nil {
		return nil, err
	}

	return l, nil
}

var _ I__iter__ = (*Lista)(nil)
var _ I__texto__ = (*Lista)(nil)
var _ I__tamanho__ = (*Lista)(nil)
var _ I__obtem_item__ = (*Lista)(nil)
var _ I__define_item__ = (*Lista)(nil)

func (l *Lista) Adiciona(item Objeto) (Objeto, error) {
	l.Itens = append(l.Itens, item)
	return nil, nil
}

func (l *Lista) Indice(obj Objeto) (Objeto, error) {
	for indice, item := range l.Itens {
		if ok, _ := Igual(item, obj); ok.(Booleano) {
			return Inteiro(indice), nil
		}
	}

	objTexto, err := NewTexto(obj)
	if err != nil {
		return nil, err
	}

	return nil, NewErroF(ValorErro, "O item '%s' não está na lista", objTexto)
}

func (l *Lista) Pop(indice Inteiro) (Objeto, error) {
	tamanho, err := l.M__tamanho__()
	if err != nil {
		return nil, err
	}

	if indice > tamanho.(Inteiro) || indice < 0 {
		return nil, NewErroF(IndiceErro, "O range é de %d indice(s), %d está fora dele", tamanho.(Inteiro), indice)
	}

	var removido Objeto
	var novaTupla Tupla

	for idx, item := range l.Itens {
		if idx == int(indice) {
			removido = item
			continue
		}

		novaTupla = append(novaTupla, item)
	}
	
	l.Itens = novaTupla
	return removido, nil
}

func init() {
	TipoLista.Mapa["adiciona"] = NewMetodoOuPanic("adiciona", func(inst Objeto, args Tupla) (Objeto, error) {
		if err := VerificaNumeroArgumentos("adiciona", true, args, 1, 1); err != nil {
			return nil, err
		}

		inst.(*Lista).Adiciona(args[0])
		return nil, nil
	}, "O método recebe um objeto e adiciona ao fim da lista")

	// TipoLista.Mapa["insere"] = NewMetodoOuPanic("insere", func(inst Objeto, args Tupla) (Objeto, error) {
	// 	if len(args) < 2 {
	// 		return nil, NewErroF(TipagemErro, "O método insere() esperava receber no mínimo 2 argumentos, mas recebeu um total de %v", len(args))
	// 	}

	// 	indice, objeto := args[0], args[1]

	// 	return nil, nil
	// }, "")

	TipoLista.Mapa["extende"] = NewMetodoOuPanic("extende", func(inst Objeto, args Tupla) (Objeto, error) {
		if err := VerificaNumeroArgumentos("extende", true, args, 1, 1); err != nil {
			return nil, err
		}

		inst.(*Lista).Itens = append(inst.(*Lista).Itens, (args[0].(Tupla))...)
		return nil, nil
	}, "Adiciona os elementos da lista recebida ao fim da lista atual")

	TipoLista.Mapa["remove"] = NewMetodoOuPanic("remove", func(inst Objeto, args Tupla) (Objeto, error) {
		if err := VerificaNumeroArgumentos("remove", true, args, 1, 1); err != nil {
			return nil, err
		}

		instancia := inst.(*Lista)
		idx, err := instancia.Indice(args[0])
		if err != nil {
			return nil, err
		}

		return instancia.Pop(idx.(Inteiro))
	}, "Remove um elemento da lista e o retorna, se existir")

	TipoLista.Mapa["pop"] = NewMetodoOuPanic("pop", func(inst Objeto, args Tupla) (Objeto, error) {
		if err := VerificaNumeroArgumentos("pop", true, args, 0, 1); err != nil {
			return nil, err
		}

		idx := Inteiro(0)

		if len(args) == 1 {
			idx = args[0].(Inteiro)
		}

		return inst.(*Lista).Pop(idx)
	}, "Remove um item da lista com base no seu índice")

	TipoLista.Mapa["indice"] = NewMetodoOuPanic("indice", func(inst Objeto, args Tupla) (Objeto, error) {
		if err := VerificaNumeroArgumentos("indice", true, args, 1, 1); err != nil {
			return nil, err
		}

		return inst.(*Lista).Indice(args[0])
	}, "Retorna o índice de um elemento se ele existir na lista")

	TipoLista.Mapa["limpa"] = NewMetodoOuPanic("limpa", func(inst Objeto) (Objeto, error) {
		inst.(*Lista).Itens = Tupla(nil)
		return nil, nil
	}, "Limpa completamente a lista")
}