package parser

// FIXME: talvez os nós devessem não só guardar o valor, mas também o token inteiro nos lugares devidos?

type BaseNode interface{ isExpr() }

// Programa representa o nó raiz do programa na AST.
type Programa struct {
	Declaracoes []BaseNode // Lista de declarações no programa
}

type DeclVar struct {
	Constante     bool     // Define se é constante
	Nome          string   // Nome da variável
	Tipo          string   // Tipo da variável (opcional)
	Inicializador BaseNode // Inicializador da variável (opcional)
}

// Casos como `variavel += 1`, `variavel -= 1` e outros
type Reatribuicao struct {
	Nome      string
	Operador  string   // =, -=, +=, /=, //=, *= ou outros
	Expressao BaseNode // A expressao após o sinal de igualdade
}

type OpBinaria struct {
	Esq      BaseNode // Expressão da esquerda
	Operador string   // Operador
	Dir      BaseNode // Expressão da direita
}

type OpUnaria struct {
	Operador  string   // Operador
	Expressao BaseNode // Expressão
}

type DeclFuncao struct {
	Nome       string                 // Nome da função
	Parametros []*DeclFuncaoParametro // Parâmetros da função
	Corpo      *Bloco                 // Corpo da função
}

type DeclFuncaoParametro struct {
	Nome   string   // Nome do parametro
	Tipo   string   // Tipo do parametro (opcional)
	Padrao BaseNode // Valor padrão para o parametro (opcional)
}

type Bloco struct {
	Declaracoes []BaseNode
}

type RetorneNode struct {
	Expressao BaseNode
}

type ChamadaFuncao struct {
	Nome       string     // Nome da função a ser chamada
	Argumentos []BaseNode // Argumentos da função
}

type ExpressaoSe struct {
	Condicao    BaseNode // Condição do if
	Corpo       *Bloco   // Corpo do if
	Alternativa BaseNode // Corpo do else ou if else (opcional)
}

type ExpressaoEnquanto struct {
	Condicao BaseNode // Condição do while
	Corpo    Bloco    // Corpo do while
}

type TextoLiteral struct {
	Valor string
}

type InteiroLiteral struct {
	Valor string
}

type DecimalLiteral struct {
	Valor string
}

type ConstanteLiteral struct {
	Valor string
}

type Identificador struct {
	Nome string
}

type Anotacao struct {
	Corpo string
}

func (*Programa) isExpr()            {}
func (*DeclVar) isExpr()             {}
func (*Reatribuicao) isExpr()        {}
func (*OpBinaria) isExpr()           {}
func (*OpUnaria) isExpr()            {}
func (*DeclFuncao) isExpr()          {}
func (*DeclFuncaoParametro) isExpr() {}
func (*Bloco) isExpr()               {}
func (*RetorneNode) isExpr()         {}
func (*ChamadaFuncao) isExpr()       {}
func (*ExpressaoSe) isExpr()         {}
func (*ExpressaoEnquanto) isExpr()   {}
func (*TextoLiteral) isExpr()        {}
func (*InteiroLiteral) isExpr()      {}
func (*DecimalLiteral) isExpr()      {}
func (*ConstanteLiteral) isExpr()    {}
func (*Identificador) isExpr()       {}
func (*Anotacao) isExpr()            {}
