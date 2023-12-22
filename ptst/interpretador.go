package ptst

import (
	"strconv"

	"github.com/natanfeitosa/portuscript/parser"
)

type Interpretador struct {
	Ast      parser.BaseNode
	Contexto *Contexto
	Caminho  Texto
}

func (i *Interpretador) Inicializa() (Objeto, error) {
	var declaracoes []parser.BaseNode

	switch ast := i.Ast.(type) {
	case *parser.Programa:
		declaracoes = ast.Declaracoes
	case *parser.Bloco:
		declaracoes = ast.Declaracoes
	default:
		return nil, i.criarErroF(TipagemErro, "Quando usar o método `Inicializa()`, a ast deve ser do tipo `Programa` ou `Bloco`")
	}

	return i.Visite(declaracoes)
}

func (i *Interpretador) Visite(nodes []parser.BaseNode) (Objeto, error) {
	var resultado Objeto
	var err error

	for _, node := range nodes {
		resultado, err = i.visite(node)
		adicionaContextoSeNaoTiver(err, i.Contexto)

		if err != nil {
			return nil, err
		}

		if _, ok := node.(*parser.RetorneNode); ok {
			return resultado, nil
		}
	}

	return resultado, nil
}

func (i *Interpretador) visite(node parser.BaseNode) (Objeto, error) {
	switch node.(type) {
	case *parser.DeclVar:
		return i.visiteDeclVar(node.(*parser.DeclVar))
	case *parser.DeclFuncao:
		return i.visiteDeclFuncao(node.(*parser.DeclFuncao))
	case *parser.ChamadaFuncao:
		return i.visiteChamadaFuncao(node.(*parser.ChamadaFuncao))
	case *parser.TextoLiteral:
		return i.visiteTextoLiteral(node.(*parser.TextoLiteral))
	case *parser.InteiroLiteral:
		return i.visiteInteiroLiteral(node.(*parser.InteiroLiteral))
	case *parser.OpBinaria:
		return i.visiteOpBinaria(node.(*parser.OpBinaria))
	case *parser.Identificador:
		return i.visiteIdentificador(node.(*parser.Identificador))
	case *parser.Reatribuicao:
		return i.visiteReatribuicao(node.(*parser.Reatribuicao))
	case *parser.ExpressaoSe:
		return i.visiteExpressaoSe(node.(*parser.ExpressaoSe))
	case *parser.Bloco:
		return i.visiteBloco(node.(*parser.Bloco))
	case *parser.RetorneNode:
		return i.visiteRetorneNode(node.(*parser.RetorneNode))
	case *parser.Enquanto:
		return i.visiteEnquanto(node.(*parser.Enquanto))
	case *parser.AcessoMembro:
		return i.visiteAcessoMembro(node.(*parser.AcessoMembro))
	case *parser.PareNode:
		return i.visitePareNode(node.(*parser.PareNode))
	case *parser.ContinueNode:
		return i.visiteContinueNode(node.(*parser.ContinueNode))
	}

	return nil, nil
}

func (i *Interpretador) visiteDeclVar(node *parser.DeclVar) (Objeto, error) {
	var valor Objeto = Nulo

	// FIXME: constante sempre deve ter o inicializador
	if node.Inicializador != nil {
		val, err := i.visite(node.Inicializador)

		if err != nil {
			err.(*Erro).AdicionarContexto(i.Contexto)
			return nil, err
		}

		valor = val
	}

	// FIXME: vamos apenas ignorar o tipo?
	simbolo := NewVarSimbolo(node.Nome, valor)

	if node.Constante {
		simbolo.Flag = SimboloConstanteFlag
	}

	err := i.Contexto.DefinirSimboloLocal(simbolo)

	if err != nil {
		err.(*Erro).AdicionarContexto(i.Contexto)
		return nil, err
	}

	return nil, nil
}

// FIXME: Isso vai funcionar?
func (i *Interpretador) visiteDeclFuncao(node *parser.DeclFuncao) (Objeto, error) {
	funcao := NewFuncao(node.Nome, node.Corpo, i.Contexto)

	for _, param := range node.Parametros {
		funcao.args = append(funcao.args, Texto(param.Nome))
	}

	simbolo := NewVarSimbolo(node.Nome, funcao)
	simbolo.Flag = SimboloFuncaoFlag

	err := i.Contexto.DefinirSimboloLocal(simbolo)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Interpretador) visiteChamadaFuncao(node *parser.ChamadaFuncao) (Objeto, error) {
	objeto, err := i.visite(node.Identificador)

	if err != nil {
		return nil, err
	}
	// if !simbolo.isFuncao() {
	// 	return nil, NewErroF(TipagemErro, "'%s' não é um chamável tipo função", simbolo.Nome)
	// }

	var args Tupla

	for _, argnode := range node.Argumentos {
		arg, err := i.visite(argnode)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	return Chamar(objeto, args)
}

func (i *Interpretador) visiteTextoLiteral(node *parser.TextoLiteral) (Objeto, error) {
	return Texto(node.Valor[1 : len(node.Valor)-1]), nil
}

func (i *Interpretador) visiteInteiroLiteral(node *parser.InteiroLiteral) (Objeto, error) {
	numero, err := strconv.ParseInt(node.Valor, 10, 64)

	if err != nil {
		return nil, err
	}

	return Inteiro(numero), nil
}

func (i *Interpretador) visiteOpBinaria(node *parser.OpBinaria) (Objeto, error) {
	esquerda, err := i.visite(node.Esq)

	if err != nil {
		return nil, err
	}

	direita, err := i.visite(node.Dir)

	if err != nil {
		return nil, err
	}

	switch node.Operador {
	case "+":
		return Adiciona(esquerda, direita)
	case "*":
		return Multiplica(esquerda, direita)
	case "-":
		return Subtrai(esquerda, direita)
	case "/":
		return Divide(esquerda, direita)
	case "<":
		return MenorQue(esquerda, direita)
	case "<=":
		return MenorOuIgual(esquerda, direita)
	case "==":
		return Igual(esquerda, direita)
	case "!=":
		return Diferente(esquerda, direita)
	case ">":
		return MaiorQue(esquerda, direita)
	case ">=":
		return MenorOuIgual(esquerda, direita)
	case "ou":
		return Ou(esquerda, direita)
	case "e":
		return E(esquerda, direita)
	}

	return nil, NewErroF(TipagemErro, "A operação '%s' não é suportada entre os tipos '%s' e '%s'", node.Operador, esquerda.Tipo().Nome, direita.Tipo().Nome)
}

func (i *Interpretador) visiteIdentificador(node *parser.Identificador) (Objeto, error) {
	simbolo, err := i.Contexto.ObterSimbolo(node.Nome)

	if err != nil {
		err.(*Erro).AdicionarContexto(i.Contexto)
		return nil, err
	}

	return simbolo.Valor, nil
}

// FIXME: estamos ignorando o operador, não?
func (i *Interpretador) visiteReatribuicao(node *parser.Reatribuicao) (Objeto, error) {
	valor, err := i.visite(node.Expressao)
	if err != nil {
		return nil, err
	}

	err = i.Contexto.RedefinirSimbolo(node.Nome, valor)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Interpretador) visiteExpressaoSe(node *parser.ExpressaoSe) (Objeto, error) {
	condicao, err := i.visite(node.Condicao)
	if err != nil {
		return nil, err
	}

	if condicao, err = NewBooleano(condicao); err != nil {
		return nil, err
	}

	if condicao.(Booleano) {
		return i.visite(node.Corpo)
	}

	return i.visite(node.Alternativa)
}

func (i *Interpretador) visiteBloco(node *parser.Bloco) (Objeto, error) {
	for _, decl := range node.Declaracoes {
		if _, err := i.visite(decl); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// FIXME: adicionar erro para caso não esteja dentro de função
func (i *Interpretador) visiteRetorneNode(node *parser.RetorneNode) (Objeto, error) {
	return i.visite(node.Expressao)
}

func (i *Interpretador) visiteEnquanto(node *parser.Enquanto) (Objeto, error) {
	for {
		condicao, err := i.visite(node.Condicao)
		if err != nil {
			return nil, err
		}

		if condicao, err = NewBooleano(condicao); err != nil {
			return nil, err
		}

		if !condicao.(Booleano) {
			break
		}

		_, err = i.visite(node.Corpo)
		if err != nil {
			if objErr, ok := err.(*Erro); ok {
				switch objErr.Tipo() {
				case ErroContinue:
					// Continue para a próxima iteração do loop
					continue
				case ErroPare:
					// Pare o loop
					return nil, nil
				}
			}

			return nil, err
		}
	}

	return nil, nil
}

func (i *Interpretador) visiteAcessoMembro(node *parser.AcessoMembro) (Objeto, error) {
	dono, err := i.visite(node.Dono)
	if err != nil {
		return nil, err
	}

	// membro, err := i.visite(node.Membro)
	// if err != nil {
	// 	return nil, err
	// }

	membro := node.Membro.(*parser.Identificador).Nome
	return ObtemItemS(dono, membro)
}

func (i *Interpretador) visitePareNode(node *parser.PareNode) (Objeto, error) {
	return nil, NewErro(ErroPare, nil)
}

func (i *Interpretador) visiteContinueNode(node *parser.ContinueNode) (Objeto, error) {
	return nil, NewErro(ErroContinue, nil)
}

func (i *Interpretador) criarErro(tipo *Tipo, args Objeto) error {
	erro := NewErro(tipo, args)
	erro.Contexto = i.Contexto
	return erro
}

func (i *Interpretador) criarErroF(tipo *Tipo, format string, args ...any) error {
	erro := NewErroF(tipo, format, args...)
	erro.Contexto = i.Contexto
	return erro
}
