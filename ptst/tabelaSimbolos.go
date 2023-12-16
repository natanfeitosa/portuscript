package ptst

type SimboloFlag int

const (
	SimboloVariavelFlag SimboloFlag = iota
	SimboloConstanteFlag
	SimboloFuncaoFlag
	SimboloClasseFlag
)

type Simbolo struct {
	Nome  string
	Valor Objeto
	Flag  SimboloFlag
}

func NewVarSimbolo(nome string, valor Objeto) *Simbolo {
	return &Simbolo{nome, valor, SimboloVariavelFlag}
}

func NewConstSimbolo(nome string, valor Objeto) *Simbolo {
	return &Simbolo{nome, valor, SimboloConstanteFlag}
}

// func NewFuncSimbolo(nome string, valor Objeto) *Simbolo {
// 	return &Simbolo{nome, valor, SimboloFuncaoFlag}
// }

func (s *Simbolo) isVariavel() bool {
	return s.Flag == SimboloVariavelFlag
}

func (s *Simbolo) isConstante() bool {
	return s.Flag == SimboloConstanteFlag
}

// func (s *Simbolo) isFuncao() bool {
// 	return s.Flag == SimboloFuncaoFlag
// }

// func (s *Simbolo) isClasse() bool {
// 	return s.Flag == SimboloClasseFlag
// }

type TabelaSimbolos struct {
	Simbolos map[string]*Simbolo
}

func NewTabelaSimbolos() *TabelaSimbolos {
	return &TabelaSimbolos{Simbolos: make(map[string]*Simbolo)}
}

func (t *TabelaSimbolos) Len() int {
	return len(t.Simbolos)
}

func (t *TabelaSimbolos) DefinirSimbolo(simbolo *Simbolo) error {
	exists, ok := t.Simbolos[simbolo.Nome]
	if ok && exists != nil {
		if exists.isConstante() {
			return NewErroF(TipagemErro, "A constante '%s' já existe", simbolo.Nome)
		}
	}

	t.Simbolos[simbolo.Nome] = simbolo
	return nil
}

func (t *TabelaSimbolos) RedefinirSimbolo(nome string, valor Objeto) error {
	simbolo, ok := t.Simbolos[nome]

	if !ok {
		return NewErroF(TipagemErro, "Você não pode reatribuir valor a '%s', pois é ela não existe", nome)
	}

	if simbolo.isConstante() {
		return NewErroF(TipagemErro, "Você não pode reatribuir valor a '%s', pois é uma constante", nome)
	}

	simbolo.Valor = valor
	return nil
}

func (t *TabelaSimbolos) ObterSimbolo(nome string) (*Simbolo, error) {
	simbolo, ok := t.Simbolos[nome]

	if !ok {
		return nil, NewErroF(NomeErro, "'%s' não foi encontrado/a, talvez você não tenha definido ainda", nome)
	}

	return simbolo, nil
}

func (t *TabelaSimbolos) ExcluirSimbolo(nome string) error {
	if _, ok := t.Simbolos[nome]; !ok {
		return NewErroF(NomeErro, "'%s' não foi encontrado/a, talvez você não tenha definido ainda", nome)
	}

	delete(t.Simbolos, nome)
	return nil
}
