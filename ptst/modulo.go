package ptst

import (
	"sync"
)

type ModuloFlags int32

const (
	ModuloCompartilhavel ModuloFlags = iota
	ModuloEntrada                    // O módulo/arquivo de entrada usado na linha de comando
)

type ModuloInfo struct {
	Nome          string // Disponível em __nome__
	Doc           string // Disponível em __doc__
	CaminhoModulo string // Caminho relativo (talvez?) referente ao módulo. Disponível em __caminho__
	Flags         ModuloFlags
}

type ModuloImpl struct {
	Info    ModuloInfo
	Metodos []*Metodo

	// Talvez os dois próximos sejam um pouco redundante
	Constantes Mapa
	Variaveis  Mapa
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
	Impl     *ModuloImpl
	Contexto *Contexto
}

var TipoModulo = NewTipo("Modulo", "Modulo doc")

func (m *Modulo) Tipo() *Tipo {
	return TipoModulo
}

type TabelaModulos struct {
	modulos   map[string]*Modulo
	Embutidos *Modulo
}

func NewTabelaModulos() *TabelaModulos {
	return &TabelaModulos{modulos: make(map[string]*Modulo)}
}

func (tabela *TabelaModulos) NewModulo(ctx *Contexto, impl *ModuloImpl) (*Modulo, error) {
	nome := impl.Info.Nome
	modulo := &Modulo{impl, NewContexto(Texto(impl.Info.CaminhoModulo))}

	if nome == "" {
		nome = "__entrada__"
		modulo.Impl.Info.Flags = ModuloEntrada
	}

	// modulo.Contexto.Caminho = Texto(impl.Info.CaminhoModulo)
	modulo.Contexto.Globais.DefinirSimbolo(NewConstSimbolo("__nome__", Texto(nome)))
	modulo.Contexto.Globais.DefinirSimbolo(NewConstSimbolo("__caminho__", Texto(impl.Info.CaminhoModulo)))
	modulo.Contexto.Globais.DefinirSimbolo(NewConstSimbolo("__doc__", Texto(impl.Info.Doc)))

	for _, metodo := range impl.Metodos {
		// metodo.Modulo = modulo
		instMetodo := new(Metodo)
		*instMetodo = *metodo
		instMetodo.Modulo = modulo
		modulo.Contexto.Globais.DefinirSimbolo(NewVarSimbolo(metodo.Nome, instMetodo))
	}

	for nome, valor := range impl.Constantes {
		modulo.Contexto.Globais.DefinirSimbolo(NewConstSimbolo(string(nome), valor))
	}

	for nome, valor := range impl.Variaveis {
		modulo.Contexto.Globais.DefinirSimbolo(NewVarSimbolo(string(nome), valor))
	}

	tabela.modulos[nome] = modulo
	if nome == "embutidos" {
		tabela.Embutidos = modulo
	}

	return modulo, nil
}

func (tabela *TabelaModulos) RecuperaModulo(nome string) (*Modulo, error) {
	if modulo, ok := tabela.modulos[nome]; ok {
		return modulo, nil
	}

	return nil, NewErroF(ImportacaoErro, "O módulo '%v' não pode ser achado", nome)
}
