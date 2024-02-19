package ptst

import (
	"sync"

	"github.com/natanfeitosa/portuscript/parser"
)

type ModuloInfo struct {
	Nome          string // Disponível em __nome__
	Doc           string // Disponível em __doc__
	Arquivo string // Caminho relativo (talvez?) referente ao módulo. Disponível em __caminho__
}

type ModuloImpl struct {
	Info    ModuloInfo
	Metodos []*Metodo

	// Talvez os dois próximos sejam um pouco redundante
	Constantes Mapa
	Variaveis  Mapa

	Ast parser.BaseNode
}

type GerenciadorModulos struct {
	mu    sync.RWMutex
	Impls map[string]*ModuloImpl
}

func (g *GerenciadorModulos) RegistraModuloImpl(impl *ModuloImpl) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Impls[impl.Info.Nome] = impl
}

func (g *GerenciadorModulos) ObtemImplModulo(nome string) *ModuloImpl {
	g.mu.RLock()
	defer g.mu.RUnlock()
	impl := g.Impls[nome]
	return impl
}

var gerenciador = GerenciadorModulos{
	Impls: make(map[string]*ModuloImpl),
}

func RegistraModuloImpl(impl *ModuloImpl) {
	gerenciador.RegistraModuloImpl(impl)
}

func ObtemImplModulo(nome string) *ModuloImpl {
	return gerenciador.ObtemImplModulo(nome)
}

type Modulo struct {
	Impl         *ModuloImpl
	Contexto     *Contexto
	Escopo       *Escopo
	acessoRapido Mapa
}

var TipoModulo = NewTipo("Modulo", "Modulo doc")

func (m *Modulo) Tipo() *Tipo {
	return TipoModulo
}

func (m *Modulo) M__obtem_attributo__(nome string) (objeto Objeto, err error) {
	ok := false
	if objeto, ok = m.acessoRapido[nome]; ok {
		return
	}

	objeto, err = m.Escopo.ObterValor(nome)
	if err != nil {
		return
	}

	m.acessoRapido[nome] = objeto
	return
}

var _ I__obtem_attributo__ = (*Modulo)(nil)

type TabelaModulos struct {
	modulos   map[string]*Modulo
	Embutidos *Modulo
}

func NewTabelaModulos() *TabelaModulos {
	return &TabelaModulos{modulos: make(map[string]*Modulo)}
}

func (tabela *TabelaModulos) NewModulo(ctx *Contexto, impl *ModuloImpl) (*Modulo, error) {
	nome := impl.Info.Nome
	modulo := &Modulo{
		Impl:         impl,
		Contexto:     ctx,
		Escopo:       NewEscopo(),
		acessoRapido: NewMapaVazio(),
	}

	if nome == "" {
		nome = "__entrada__"
	}

	modulo.Escopo.DefinirSimbolo(NewConstSimbolo("__nome__", Texto(nome)))
	modulo.Escopo.DefinirSimbolo(NewConstSimbolo("__arquivo__", Texto(impl.Info.Arquivo)))
	modulo.Escopo.DefinirSimbolo(NewConstSimbolo("__doc__", Texto(impl.Info.Doc)))

	for _, metodo := range impl.Metodos {
		// metodo.Modulo = modulo
		instMetodo := new(Metodo)
		*instMetodo = *metodo
		instMetodo.Modulo = modulo
		modulo.Escopo.DefinirSimbolo(NewVarSimbolo(metodo.Nome, instMetodo))
	}

	for nome, valor := range impl.Constantes {
		modulo.Escopo.DefinirSimbolo(NewConstSimbolo(string(nome), valor))
	}

	for nome, valor := range impl.Variaveis {
		modulo.Escopo.DefinirSimbolo(NewVarSimbolo(string(nome), valor))
	}

	tabela.modulos[nome] = modulo
	if nome == "embutidos" {
		tabela.Embutidos = modulo
	}

	return modulo, nil
}

func (tabela *TabelaModulos) ObterModulo(nome string) (*Modulo, error) {
	m, ok := tabela.modulos[nome]
	if !ok {
		return nil, NewErroF(ImportacaoErro, "O módulo '%v' não pode ser achado", nome)
	}

	return m, nil
}
