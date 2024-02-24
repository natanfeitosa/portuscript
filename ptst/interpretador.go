package ptst

import (
	"github.com/natanfeitosa/portuscript/parser"
)

type Interpretador struct {
	Ast          parser.BaseNode
	Contexto     *Contexto
	Escopo       *Escopo
	ValorRetorno Objeto
}

func (i *Interpretador) entrarNoEscopo(escopo *Escopo) {
	if escopo == nil {
		escopo = i.Escopo.NewEscopo()
	}

	if escopo.Pai == nil && i.Escopo != nil {
		escopo.Pai = i.Escopo
	}

	i.Escopo = escopo
}

func (i *Interpretador) sairDoEscopo() {
	i.Escopo = i.Escopo.Pai
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

		// Verifica se há um valor de retorno
		if i.ValorRetorno != nil {
			return i.ValorRetorno, nil
		}
	}

	return resultado, nil
}

func (i *Interpretador) visite(astNode parser.BaseNode) (Objeto, error) {
	switch node := astNode.(type) {
	case *parser.DeclVar:
		return i.visiteDeclVar(node)
	case *parser.DeclFuncao:
		return i.visiteDeclFuncao(node)
	case *parser.ChamadaFuncao:
		return i.visiteChamadaFuncao(node)
	case *parser.TextoLiteral:
		return i.visiteTextoLiteral(node)
	case *parser.InteiroLiteral:
		return i.visiteInteiroLiteral(node)
	case *parser.DecimalLiteral:
		return i.visiteDecimalLiteral(node)
	case *parser.TuplaLiteral:
		return i.visiteTuplaLiteral(node)
	case *parser.ListaLiteral:
		return i.visiteListaLiteral(node)
	case *parser.OpBinaria:
		return i.visiteOpBinaria(node)
	case *parser.OpUnaria:
		return i.visiteOpUnaria(node)
	case *parser.Identificador:
		return i.visiteIdentificador(node)
	case *parser.Reatribuicao:
		return i.visiteReatribuicao(node)
	case *parser.ExpressaoSe:
		return i.visiteExpressaoSe(node)
	case *parser.Bloco:
		return i.visiteBloco(node)
	case *parser.RetorneNode:
		return i.visiteRetorneNode(node)
	case *parser.Enquanto:
		return i.visiteEnquanto(node)
	case *parser.AcessoMembro:
		return i.visiteAcessoMembro(node)
	case *parser.BlocoPara:
		return i.visiteBlocoPara(node)
	case *parser.PareNode:
		return i.visitePareNode(node)
	case *parser.ContinueNode:
		return i.visiteContinueNode(node)
	case *parser.ImporteDe:
		return i.visiteImporteDe(node)
	case *parser.Indexacao:
		return i.visiteIndexacao(node)
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
		simbolo.Constante = true
	}

	if err := i.Escopo.DefinirSimbolo(simbolo); err != nil {
		return nil, err
	}

	return nil, nil
}

// FIXME: Isso vai funcionar?
func (i *Interpretador) visiteDeclFuncao(node *parser.DeclFuncao) (Objeto, error) {
	funcao := NewFuncao(node.Nome, node.Corpo, i.Contexto, i.Escopo)

	for _, param := range node.Parametros {
		funcao.args = append(funcao.args, Texto(param.Nome))
	}

	err := i.Escopo.DefinirSimbolo(NewVarSimbolo(node.Nome, funcao))
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
	return NewTexto(node.Valor[1 : len(node.Valor)-1])
}

func (i *Interpretador) visiteInteiroLiteral(node *parser.InteiroLiteral) (Objeto, error) {
	return NewInteiro(node.Valor)
}

func (i *Interpretador) visiteDecimalLiteral(node *parser.DecimalLiteral) (Objeto, error) {
	return NewDecimal(node.Valor)
}

func (i *Interpretador) visiteTuplaLiteral(node *parser.TuplaLiteral) (Objeto, error) {
	var tupla Tupla

	for _, elemento := range node.Elementos {
		item, err := i.visite(elemento)
		if err != nil {
			return nil, err
		}

		tupla = append(tupla, item)
	}
	return tupla, nil
}

func (i *Interpretador) visiteListaLiteral(node *parser.ListaLiteral) (Objeto, error) {
	lista := &Lista{}

	for _, elemento := range node.Elementos {
		item, err := i.visite(elemento)
		if err != nil {
			return nil, err
		}

		lista.Adiciona(item)
	}
	return lista, nil
}

func (i *Interpretador) visiteOpUnaria(node *parser.OpUnaria) (Objeto, error) {
	operando, err := i.visite(node.Expressao)

	if err != nil {
		return nil, err
	}

	switch node.Operador {
	case "-":
		return Neg(operando)
	case "+":
		return Pos(operando)
	}

	return nil, NewErroF(TipagemErro, "A operação '%s' não é suportada para o tipo '%s'", node.Operador, operando.Tipo().Nome)
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
	case "//":
		return DivideInteiro(esquerda, direita)
	case "%":
		return Mod(esquerda, direita)
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
	objeto, err := i.Escopo.ObterValor(node.Nome)

	if err != nil {
		if objeto, err = i.Contexto.Modulos.Embutidos.M__obtem_attributo__(node.Nome); err == nil {
			return objeto, nil
		}

		return nil, err
	}

	return objeto, nil
}

func (i *Interpretador) visiteReatribuicao(node *parser.Reatribuicao) (Objeto, error) {
	var direita, esquerda, valor Objeto
	var err error

	if esquerda, err = i.visite(node.Objeto); err != nil {
		return nil, err
	}

	if direita, err = i.visite(node.Expressao); err != nil {
		return nil, err
	}

	// FIXME: provavelmente nao seria melhor criar métodos específicos para operações implace e usar estes como fallback??
	switch node.Operador {
	case "+=":
		valor, err = Adiciona(esquerda, direita)
	case "*=":
		valor, err = Multiplica(esquerda, direita)
	case "-=":
		valor, err = Subtrai(esquerda, direita)
	case "/=":
		valor, err = Divide(esquerda, direita)
	case "//=":
		valor, err = DivideInteiro(esquerda, direita)
	default:
		valor = direita
	}

	if err != nil {
		return nil, err
	}

	switch obj := node.Objeto.(type) {
	case *parser.Indexacao:
		if esquerda, err = i.visite(obj.Objeto); err != nil {
			return nil, err
		}

		chave, err := i.visite(obj.Argumento)
		if err != nil {
			return nil, err
		}

		return DefineItem(esquerda, chave, valor)
	case *parser.Identificador:
		return nil, i.Escopo.RedefinirValor(obj.Nome, valor)
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
	i.entrarNoEscopo(nil)
	defer i.sairDoEscopo()

	for _, decl := range node.Declaracoes {
		if _, err := i.visite(decl); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// FIXME: adicionar erro para caso não esteja dentro de função
func (i *Interpretador) visiteRetorneNode(node *parser.RetorneNode) (Objeto, error) {
	// Se encontrarmos um retorne, definimos o valor de retorno no interpretador
	valor, err := i.visite(node.Expressao)
	if err != nil {
		return nil, err
	}
	i.ValorRetorno = valor
	return valor, nil
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
	return ObtemAtributoS(dono, membro)
}

func (i *Interpretador) visiteBlocoPara(node *parser.BlocoPara) (Objeto, error) {
	// FIXME: isso provavelmente não é muito eficiente e correto
	i.Escopo.DefinirSimbolo(NewVarSimbolo(node.Identificador, Nulo))
	defer func() {
		i.Escopo.ExcluirSimbolo(node.Identificador)
	}()

	var item, iterador Objeto
	var err error

	if iterador, err = i.visite(node.Iterador); err != nil {
		return nil, err
	}

	if iterador, err = Iter(iterador); err != nil {
		return nil, err
	}

	for {
		if item, err = Proximo(iterador); err != nil {
			if objErr, ok := err.(*Erro); ok {
				if objErr.Tipo() == FimIteracao {
					return nil, nil
				}
			}

			return nil, err
		}

		i.Escopo.RedefinirValor(node.Identificador, item)

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
}

func (i *Interpretador) visitePareNode(node *parser.PareNode) (Objeto, error) {
	return nil, NewErro(ErroPare, Nulo)
}

func (i *Interpretador) visiteContinueNode(node *parser.ContinueNode) (Objeto, error) {
	return nil, NewErro(ErroContinue, Nulo)
}

func (i *Interpretador) visiteImporteDe(node *parser.ImporteDe) (Objeto, error) {
	caminho, err := i.visiteTextoLiteral(node.Caminho)
	if err != nil {
		return nil, err
	}

	modulo, err := MaquinarioImporteModulo(i.Contexto, string(caminho.(Texto)), i.Escopo)
	if err != nil {
		return nil, err
	}

	for _, nome := range node.Nomes {
		obj, err := ObtemAtributoS(modulo, nome)
		if err != nil {
			return nil, err
		}

		i.Escopo.DefinirSimbolo(NewVarSimbolo(nome, obj))
	}

	return nil, nil
}

func (i *Interpretador) visiteIndexacao(node *parser.Indexacao) (Objeto, error) {
	var obj, arg Objeto
	var err error

	if obj, err = i.visite(node.Objeto); err != nil {
		return nil, err
	}

	if arg, err = i.visite(node.Argumento); err != nil {
		return nil, err
	}

	return ObtemItem(obj, arg)
}

func (i *Interpretador) criarErroF(tipo *Tipo, format string, args ...any) error {
	erro := NewErroF(tipo, format, args...)
	erro.Contexto = i.Contexto
	return erro
}
