package ptst

import (
	"strings"
	"sync"
)

type Contexto struct {
	Pai       *Contexto       // O contexto anterior (pai), se aplicavel
	Caminho   Texto           // O caminho (path) do arquivo em execução
	Locais    *TabelaSimbolos // Variaveis e Constantes Locais
	Globais   *TabelaSimbolos // Variaveis e Constantes Globais
	Modulos   *TabelaModulos
	waitgroup sync.WaitGroup
	once      sync.Once
	fechado   bool
	// ErroAtual *Erro
}

func NewContexto(caminho Texto) *Contexto {
	context := &Contexto{Caminho: caminho}
	context.Locais = NewTabelaSimbolos()
	context.Globais = NewTabelaSimbolos()
	context.Modulos = NewTabelaModulos()
	context.fechado = false
	return context
}

func NewContextoI(i *Interpretador) *Contexto {
	contexto := NewContexto(i.Caminho)
	i.Contexto = contexto
	return contexto
}

func (c *Contexto) NewContexto() *Contexto {
	context := NewContexto(c.Caminho)
	context.Pai = c
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

func (c *Contexto) ExcluirSimbolo(nome string) error {
	return c.Locais.ExcluirSimbolo(nome)
}

func (c *Contexto) ObterModulo(nome string) (Objeto, error) {
	return c.Modulos.ObterModulo(nome)
}

func (c *Contexto) InicializarModulo(implementacao *ModuloImpl) (Objeto, error) {
	if err := c.adicionarTrabalho(); err != nil {
		return nil, err
	}
	defer c.encerrarTrabalho()
	// FIXME: adicionar a lógica para compilação e cache de módulos definidos do lado ptst da história

	modulo, err := c.Modulos.NewModulo(c, implementacao)
	if err != nil {
		return nil, err
	}

	return modulo, nil
}

func (c *Contexto) adicionarTrabalho() error {
	if c.fechado {
		err := NewErro(RuntimeErro, Texto("Contexto já fechado"))
		err.Contexto = c
		return err
	}

	c.waitgroup.Add(1)
	return nil
}

func (c *Contexto) encerrarTrabalho() {
	c.waitgroup.Done()
}

func (c *Contexto) Terminar() {
	c.once.Do(func () {
		c.waitgroup.Wait()
		c.fechado = true
	})
}