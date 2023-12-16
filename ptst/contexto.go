package ptst

import (
	"strings"
)

type Contexto struct {
	Pai     *Contexto       // O contexto anterior (pai), se aplicavel
	Caminho Texto           // O caminho (path) do arquivo em execução
	Locais  *TabelaSimbolos // Variaveis e Constantes Locais
	Globais *TabelaSimbolos // Variaveis e Constantes Globais
	Modulos *TabelaModulos
	// ErroAtual *Erro
}

func NewContexto(caminho Texto) *Contexto {
	context := &Contexto{Caminho: caminho}
	context.Globais = NewTabelaSimbolos()
	context.Locais = NewTabelaSimbolos()
	return context
}

func (c *Contexto) NewContexto() *Contexto {
	context := &Contexto{
		Pai:     c,
		Caminho: c.Caminho,
		Locais:  NewTabelaSimbolos(),
		Globais: NewTabelaSimbolos(),
	}
	return context
}

func (c *Contexto) RedefinirSimbolo(nome string, valor Objeto) error {
	ignorar := func(er error) bool {
		return strings.HasSuffix(string(er.(*Erro).Mensagem.(Texto)), "ela não existe")
	}

	err := c.Locais.RedefinirSimbolo(nome, valor)
	if err != nil && !ignorar(err) {
		return err
	}

	if c.Globais != nil && c.Globais.Len() > 0 {
		err = c.Globais.RedefinirSimbolo(nome, valor)
		if err != nil && !ignorar(err) {
			return err
		}
	}

	if c.Pai != nil {
		err = c.Pai.RedefinirSimbolo(nome, valor)
		if err != nil && !ignorar(err) {
			return err
		}
	}

	// return c.Locais.DefinirSimbolo(simbolo)
	return nil
}

func (c *Contexto) DefinirSimboloLocal(simbolo *Simbolo) error {
	return c.Locais.DefinirSimbolo(simbolo)
}

func (c *Contexto) ObterSimbolo(nome string) (simbolo *Simbolo, err error) {
	simbolo, err = c.Locais.ObterSimbolo(nome)

	if simbolo != nil {
		return
	}

	simbolo, err = c.Globais.ObterSimbolo(nome)

	if simbolo != nil {
		return
	}

	// FIXME: isso realmente tá certo?
	if c.Modulos != nil && c.Modulos.Embutidos != nil {
		return c.Modulos.Embutidos.Contexto.ObterSimbolo(nome)
	}

	if c.Pai != nil && err != nil {
		return c.Pai.ObterSimbolo(nome)
	}

	return
}
