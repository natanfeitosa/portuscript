package ptst

import (
	"os"
	"sync"

	"github.com/natanfeitosa/portuscript/parser"
)

type OpcsContexto struct {
	// Argumentos do terminal quando roda por exemplo: `go run ./main.go`
	Args []string
	// Caminhos possíveis para resolução de módulos e arquivos
	CaminhosPadrao []string
}

type Contexto struct {
	Modulos   *TabelaModulos
	Opcs      OpcsContexto
	fechado   bool
	waitgroup sync.WaitGroup
	once      sync.Once
	// ErroAtual *Erro
}

func NewContexto(opcs OpcsContexto) *Contexto {
	context := &Contexto{
		Modulos: NewTabelaModulos(),
		Opcs:    opcs,
		fechado: false,
	}

	MultiImporteModulo(context, "embutidos")
	return context
}

func (c *Contexto) TransformarEmAst(caminho string, useSysPaths bool, curDir string) (string, parser.BaseNode, error) {
	if err := c.adicionarTrabalho(); err != nil {
		return "", nil, err
	}
	defer c.encerrarTrabalho()

	caminhos := []string{}
	if useSysPaths {
		caminhos = c.Opcs.CaminhosPadrao
	}

	caminho, err := ResolveArquivoPtst(caminho, caminhos, curDir)
	if err != nil {
		return "", nil, err
	}

	codigo, err := os.ReadFile(caminho)
	if err != nil {
		return "", nil, NewErroF(SistemaErro, "Erro ao acessar '%s': %s", caminho, err)
	}

	ast, err := c.StringParaAst(string(codigo))
	return caminho, ast, err
}

func (c *Contexto) StringParaAst(codigo string) (parser.BaseNode, error) {
	ast, err := parser.NewParserFromString(string(codigo)).Parse()
	if err != nil {
		return nil, NewErroF(SintaxeErro, "%s", err)
	}

	return ast, nil
}

func (c *Contexto) AvaliarAst(ast parser.BaseNode, escopo *Escopo) (Objeto, error) {
	if err := c.adicionarTrabalho(); err != nil {
		return nil, err
	}
	defer c.encerrarTrabalho()

	interpret := &Interpretador{Ast: ast, Contexto: c, Escopo: escopo}
	// defer interpret.Contexto.Terminar()
	MultiImporteModulo(interpret.Contexto, "embutidos")

	return interpret.Inicializa()
}

func (c *Contexto) ObterModulo(nome string) (*Modulo, error) {
	return c.Modulos.ObterModulo(nome)
}

func (c *Contexto) InicializarModulo(implementacao *ModuloImpl) (*Modulo, error) {
	if err := c.adicionarTrabalho(); err != nil {
		return nil, err
	}
	defer c.encerrarTrabalho()
	// FIXME: adicionar a lógica para compilação e cache de módulos definidos do lado ptst da história

	modulo, err := c.Modulos.NewModulo(c, implementacao)
	if err != nil {
		return nil, err
	}

	if implementacao.Ast != nil {
		_, err := c.AvaliarAst(implementacao.Ast, modulo.Escopo)
		if err != nil {
			return nil, err
		}
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
	c.once.Do(func() {
		c.waitgroup.Wait()
		c.fechado = true
	})
}
