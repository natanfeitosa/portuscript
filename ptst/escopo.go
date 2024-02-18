package ptst

type Simbolo struct {
	Nome      string
	Valor     Objeto
	Constante bool
}

type Escopo struct {
	Simbolos map[string]*Simbolo
	Pai      *Escopo
}

func NewEscopo() *Escopo {
	return &Escopo{Simbolos: make(map[string]*Simbolo)}
}

func (e *Escopo) NewEscopo() *Escopo {
	return &Escopo{Simbolos: make(map[string]*Simbolo), Pai: e}
}

func (e *Escopo) Len() int {
	return len(e.Simbolos)
}

func (e *Escopo) DefinirSimbolo(simbolo *Simbolo) error {
	exists, ok := e.Simbolos[simbolo.Nome]
	if ok && exists != nil {
		if exists.Constante {
			return NewErroF(TipagemErro, "A constante '%s' já existe", simbolo.Nome)
		}
	}

	e.Simbolos[simbolo.Nome] = simbolo
	return nil
}

func (e *Escopo) RedefinirValor(nome string, valor Objeto) error {
	simbolo, ok := e.Simbolos[nome]

	if !ok {
		if e.Pai != nil {
			return e.Pai.RedefinirValor(nome, valor)
		}

		return NewErroF(TipagemErro, "Você não pode reatribuir valor a '%s', pois a variável não existe", nome)
	}

	if simbolo.Constante {
		return NewErroF(TipagemErro, "Você não pode reatribuir valor a '%s', pois é uma constante", nome)
	}

	simbolo.Valor = valor
	return nil
}

func (e *Escopo) ObterValor(nome string) (Objeto, error) {
	simbolo, ok := e.Simbolos[nome]

	if !ok {
		if e.Pai != nil {
			return e.Pai.ObterValor(nome)
		}

		return nil, NewErroF(NomeErro, "'%s' não foi encontrado/a, talvez você não tenha definido ainda", nome)
	}

	return simbolo.Valor, nil
}

func (e *Escopo) ExcluirSimbolo(nome string) error {
	if _, ok := e.Simbolos[nome]; !ok {
		return NewErroF(NomeErro, "'%s' não foi encontrado/a, talvez você não tenha definido ainda", nome)
	}

	delete(e.Simbolos, nome)
	return nil
}

func NewVarSimbolo(nome string, valor Objeto) *Simbolo {
	return &Simbolo{nome, valor, false}
}

func NewConstSimbolo(nome string, valor Objeto) *Simbolo {
	return &Simbolo{nome, valor, true}
}