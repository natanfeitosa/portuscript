package parser

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/lexer"
)

// FIXME: adicionar suporte a erros

type Parser struct {
	lex          *lexer.Lexer // O lexer original
	token        lexer.Token
	proximoToken lexer.Token
}

func NewParser(lex *lexer.Lexer) *Parser {
	parse := &Parser{lex: lex}
	parse.avancar()
	return parse
}

func NewParserFromString(code string) *Parser {
	return NewParser(lexer.NewLexer(code))
}

func (p *Parser) avancar() {
	if p.proximoToken.Tipo == 0 && p.token.Tipo == 0 {
		p.token = p.lex.ProximoToken()
	} else {
		p.token = p.proximoToken
	}

	// fmt.Printf("\ntoken: %#v\nlex: %#v\n", p.token, p.lex)

	p.proximoToken = p.lex.ProximoToken()
}

func (p *Parser) consume(token string) error {
	if p.token.Valor != token {
		return fmt.Errorf("Era esperado o token '%v', mas no lugar foi encontrado '%v'.", token, p.token.Valor)
	}

	p.avancar()
	return nil
}

func (p *Parser) isEof() bool {
	return p.token.Tipo == lexer.TokenEOF
}

func (p *Parser) Parse() (*Programa, error) {
	declaracoes, err := p.parseDeclaracoes()

	if err != nil {
		return nil, err
	}

	return &Programa{Declaracoes: declaracoes}, nil
}

func (p *Parser) parseDeclaracoes() ([]BaseNode, error) {
	var declaracoes []BaseNode

	for !p.isEof() && p.token.Tipo != lexer.TokenFechaChaves {
		if p.token.Tipo != lexer.TokenNovaLinha {
			declaracao, err := p.parseDeclaracao()

			if err != nil {
				return nil, err
			}

			declaracoes = append(declaracoes, declaracao)
		}

		p.avancar()
	}

	return declaracoes, nil
}

func (p *Parser) parseDeclaracao() (BaseNode, error) {
	val := p.token.Valor

	if val == "const" || val == "var" {
		return p.parseVariavel()
	}

	if !IsKeyword(val) {
		proximo := p.proximoToken.Tipo
		if proximo >= lexer.TokenMaisIgual && proximo <= lexer.TokenBarraIgual || proximo == lexer.TokenIgual {
			return p.parseReatribuicao()
		}
	}

	if val == "retorne" {
		return p.parseRetorne()
	}

	// FIXME: adicionar importaçao

	if val == "func" {
		return p.parseFuncao()
	}

	if val == "se" {
		return p.parseExpressaoSe()
	}

	if val == "enquanto" {
		return p.parseEnquanto()
	}

	if val == "pare" {
		p.avancar()
		return &PareNode{}, nil
	}

	if val == "continue" {
		p.avancar()
		return &ContinueNode{}, nil
	}

	if val == "para" {
		return p.parseBlocoPara()
	}

	return p.parseExpressao()
}

func (p *Parser) parseBlocoPara() (*BlocoPara, error) {
	p.consume("para")
	if err := p.consume("("); err != nil {
		return nil, err
	}

	// FIXME: isso pode gerar erros indesejados
	id := p.token.Valor
	p.avancar()

	if err := p.consume("em"); err != nil {
		return nil, err
	}

	iter, err := p.parsePrimario()
	if err != nil {
		return nil, err
	}
	if err := p.consume(")"); err != nil {
		return nil, err
	}

	corpo, err := p.parseBloco()
	if err != nil {
		return nil, err
	}
	
	return &BlocoPara{Identificador: id, Iterador: iter, Corpo: corpo}, nil
}

func (p *Parser) parseExpressaoSe() (*ExpressaoSe, error) {
	p.consume("se")
	if err := p.consume("("); err != nil {
		return nil, err
	}

	condicao, err := p.parseExpressao()
	if err != nil {
		return nil, err
	}

	expressaoSe := &ExpressaoSe{Condicao: condicao}
	if err := p.consume(")"); err != nil {
		return nil, err
	}

	corpo, err := p.parseBloco()
	if err != nil {
		return nil, err
	}

	expressaoSe.Corpo = corpo

	if p.token.Tipo == lexer.TokenSenao {
		p.avancar()
		var alternativa BaseNode

		switch p.token.Tipo {
		case lexer.TokenSe:
			if alternativa, err = p.parseExpressaoSe(); err != nil {
				return nil, err
			}
		case lexer.TokenAbreChaves:
			if alternativa, err = p.parseBloco(); err != nil {
				return nil, err
			}
		}

		expressaoSe.Alternativa = alternativa
	}

	return expressaoSe, nil
}

func (p *Parser) parseRetorne() (*RetorneNode, error) {
	if err := p.consume("retorne"); err != nil {
		return nil, err
	}

	retorne := &RetorneNode{}

	if p.token.Tipo != lexer.TokenPontoEVirgula {
		expressao, err := p.parseExpressao()

		if err != nil {
			return nil, err
		}

		retorne.Expressao = expressao
		if err := p.consume(";"); err != nil {
			return nil, err
		}
	}
	return retorne, nil
}

func (p *Parser) parseReatribuicao() (*Reatribuicao, error) {
	reatribuicao := &Reatribuicao{}
	reatribuicao.Nome = p.token.Valor

	p.avancar()
	// FIXME: isso devia gerar um erro
	// if !(p.token.Tipo >= lexer.TokenMaisIgual && p.token.Tipo <= lexer.TokenBarraIgual) {
	// }

	reatribuicao.Operador = p.token.Valor
	p.avancar()

	expressao, err := p.parseExpressao()

	if err != nil {
		return nil, err
	}

	reatribuicao.Expressao = expressao

	if err := p.consume(";"); err != nil {
		return nil, err
	}

	return reatribuicao, nil
}

func (p *Parser) parseFuncao() (*DeclFuncao, error) {
	if err := p.consume("func"); err != nil {
		return nil, err
	}

	funcao := &DeclFuncao{}

	// FIXME: pegamos o valor do token, mas e se ele não for um indentificador válido ou nem seja id?
	funcao.Nome = p.token.Valor
	p.avancar()

	if err := p.consume("("); err != nil {
		return nil, err
	}

	for {
		if p.token.Tipo == lexer.TokenFechaParenteses {
			break
		}

		params, err := p.parseDeclFuncaoParametro()

		if err != nil {
			return nil, err
		}

		funcao.Parametros = append(funcao.Parametros, params)

		if p.token.Tipo == lexer.TokenVirgula {
			p.avancar()
		}
	}

	if err := p.consume(")"); err != nil {
		return nil, err
	}

	corpo, err := p.parseBloco()

	if err != nil {
		return nil, err
	}

	funcao.Corpo = corpo

	return funcao, nil
}

// FIXME: adicionar novo contexto ao entrar em um bloco e voltar ao anterior ao fim do bloco
func (p *Parser) parseBloco() (*Bloco, error) {
	bloco := &Bloco{}

	if err := p.consume("{"); err != nil {
		return nil, err
	}

	decl, err := p.parseDeclaracoes()

	if err != nil {
		return nil, err
	}

	bloco.Declaracoes = decl

	if err := p.consume("}"); err != nil {
		return nil, err
	}

	return bloco, nil
}

// FIXME: talvez isso pudesse ser reaproveitado na declaraçao de variáveis
func (p *Parser) parseDeclFuncaoParametro() (*DeclFuncaoParametro, error) {
	parametro := &DeclFuncaoParametro{}

	// FIXME: espera que seja um identificador (nao palavra chave), mas o que fazer caso nao seja?
	parametro.Nome = p.token.Valor
	p.avancar()

	if p.token.Tipo == lexer.TokenDoisPontos {
		if err := p.consume(":"); err != nil {
			return nil, err
		}

		parametro.Tipo = p.token.Valor
		p.avancar()
	}

	if p.token.Tipo == lexer.TokenIgual {
		if err := p.consume("="); err != nil {
			return nil, err
		}

		expressao, err := p.parseExpressao()

		if err != nil {
			return nil, err
		}

		parametro.Padrao = expressao
	}

	return parametro, nil
}

func (p *Parser) parseVariavel() (*DeclVar, error) {
	/*
	 * FIXME:
	 * 1 - Constantes sempre devem ter um valor inicializador, o tipo é opcional
	 * 2 - Variáveis devem ter um dos dois
	 * 2.1 - Se o valor inicialiador estiver presente, mostre um alerta (opcional) quando o tipo for declarado
	 * 2.2 - Se o tipo for definido, o valor é opcional e pode ser ignorado (talvez não seja certo?)
	 */

	decl := &DeclVar{}
	decl.Constante = p.token.Valor == "const"

	p.avancar()

	// FIXME: adicionar verificação para ver se o nome da variável é válido e se tem nome

	decl.Nome = p.token.Valor
	p.avancar()

	if p.token.Tipo == lexer.TokenDoisPontos {
		if err := p.consume(":"); err != nil {
			return nil, err
		}

		decl.Tipo = p.token.Valor
		p.avancar()
	}

	// FIXME: lança um erro se não tiver tipo definido e/ou for uma constante

	if p.token.Tipo == lexer.TokenIgual {
		if err := p.consume("="); err != nil {
			return nil, err
		}

		expressao, err := p.parseExpressao()

		if err != nil {
			return nil, err
		}

		decl.Inicializador = expressao
	}

	if err := p.consume(";"); err != nil {
		return nil, err
	}

	return decl, nil
}

func (p *Parser) parseExpressao() (BaseNode, error) {
	return p.parseDisjuncao()
}

func (p *Parser) parseDisjuncao() (BaseNode, error) {
	esquerda, err := p.parseConjuncao()

	if err != nil {
		return nil, err
	}

	if p.token.Tipo == lexer.TokenBoolOu {
		p.consume("ou")
		direita, err := p.parseConjuncao()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, "ou", direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseConjuncao() (BaseNode, error) {
	esquerda, err := p.parseInversao()

	if err != nil {
		return nil, err
	}

	if p.token.Tipo == lexer.TokenBoolE {
		p.consume("e")
		direita, err := p.parseInversao()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, "e", direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseInversao() (BaseNode, error) {
	if p.token.Tipo == lexer.TokenBoolOu {
		p.consume("ou")
		operacao, err := p.parseInversao()

		if err != nil {
			return nil, err
		}

		return &OpUnaria{"ou", operacao}, nil
	}

	return p.parseComparacao()
}

func (p *Parser) parseComparacao() (BaseNode, error) {
	esquerda, err := p.parseBitABitOu()
	if err != nil {
		return nil, err
	}

	token := p.token
	switch token.Tipo {
	case lexer.TokenIgualIgual,
		lexer.TokenDiferente,
		lexer.TokenMenorQue,
		lexer.TokenMenorOuIgual,
		lexer.TokenMaiorQue,
		lexer.TokenMaiorOuIgual:
		p.avancar()
		direita, err := p.parseComparacao()
		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, token.Valor, direita}, nil
	}

	return esquerda, err
}

func (p *Parser) parseBitABitOu() (BaseNode, error) {
	esquerda, err := p.parseBitABitExOu()

	if err != nil {
		return nil, err
	}

	if p.token.Tipo == lexer.TokenBitABitOu {
		p.consume("|")
		direita, err := p.parseBitABitExOu()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, "|", direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseBitABitExOu() (BaseNode, error) {
	esquerda, err := p.parseBitABitE()

	if err != nil {
		return nil, err
	}

	if p.token.Tipo == lexer.TokenBitABitExOu {
		p.consume("^")
		direita, err := p.parseBitABitE()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, "^", direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseBitABitE() (BaseNode, error) {
	esquerda, err := p.parseDeslocamento()

	if err != nil {
		return nil, err
	}

	if p.token.Tipo == lexer.TokenBitABitE {
		p.consume("&")
		direita, err := p.parseDeslocamento()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, "&", direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseDeslocamento() (BaseNode, error) {
	esquerda, err := p.parseAritBasica()

	if err != nil {
		return nil, err
	}

	token := p.token
	if token.Tipo == lexer.TokenDeslocEsquerda || token.Tipo == lexer.TokenDeslocDireita {
		p.avancar()

		direita, err := p.parseAritBasica()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, token.Valor, direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseAritBasica() (BaseNode, error) {
	esquerda, err := p.parseTermo()

	if err != nil {
		return nil, err
	}

	token := p.token
	if token.Tipo == lexer.TokenMais || token.Tipo == lexer.TokenMenos {
		p.avancar()

		direita, err := p.parseAritBasica()

		// fmt.Println()
		// fmt.Printf("esq: %#v\ntoken: %#v\ndir: %#v", esquerda, token, direita)
		// fmt.Println()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, token.Valor, direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseTermo() (BaseNode, error) {
	esquerda, err := p.parseFator()

	if err != nil {
		return nil, err
	}

	token := p.token
	switch token.Tipo {
	case lexer.TokenAsterisco, lexer.TokenDivisao, lexer.TokenDivisaoInteira, lexer.TokenModulo:
		// FIXME: adicionar também a operaçao de multiplicaçao de matrizes? @
		p.avancar()

		direita, err := p.parseTermo()
		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, token.Valor, direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parseFator() (BaseNode, error) {
	token := p.token

	switch token.Tipo {
	case lexer.TokenMais, lexer.TokenMenos, lexer.TokenBitABitNao:
		expressao, err := p.parseFator()

		if err != nil {
			return nil, err
		}

		return &OpUnaria{token.Valor, expressao}, nil
	}

	return p.parsePotencia()
}

func (p *Parser) parsePotencia() (BaseNode, error) {
	esquerda, err := p.parsePrimario()

	if err != nil {
		return nil, err
	}

	if p.token.Tipo == lexer.TokenPotencia {
		p.avancar()
		direita, err := p.parseFator()

		if err != nil {
			return nil, err
		}

		return &OpBinaria{esquerda, "**", direita}, nil
	}

	return esquerda, nil
}

func (p *Parser) parsePrimario() (BaseNode, error) {
	// FIXME
	atom, err := p.parseAtomo()
	if err != nil {
		return nil, err
	}

	for p.token.Tipo == lexer.TokenPonto {
		p.avancar()
		membro, err := p.parseAtomo()
		if err != nil {
			return nil, err
		}

		atom = &AcessoMembro{atom, membro}
	}

	if p.token.Tipo == lexer.TokenAbreParenteses {
		chamada := &ChamadaFuncao{Identificador: atom}

		if err := p.consume("("); err != nil {
			return nil, err
		}

		for p.token.Tipo != lexer.TokenFechaParenteses {
			expressao, err := p.parseExpressao()

			if err != nil {
				return nil, err
			}

			chamada.Argumentos = append(chamada.Argumentos, expressao)

			if p.token.Tipo == lexer.TokenVirgula {
				p.avancar()
			}

			if p.token.Tipo != lexer.TokenFechaParenteses && p.proximoToken.Tipo != lexer.TokenFechaParenteses {
				// FIXME: lança um erro se o token atual não for virgula e o
				// próximo também não for o fechamento dos parentess
			}
		}

		if err := p.consume(")"); err != nil {
			return nil, err
		}

		return chamada, nil
	}

	return atom, nil
}

func (p *Parser) parseAtomo() (BaseNode, error) {
	token := p.token
	switch token.Tipo {
	// case lexer.TokenVerdadeiro, lexer.TokenFalso, lexer.TokenNulo:
	// 	p.avancar()
	// 	return &ConstanteLiteral{token.Valor}, nil
	case lexer.TokenTexto:
		p.avancar()
		return &TextoLiteral{token.Valor}, nil
	case lexer.TokenDecimal:
		p.avancar()
		return &DecimalLiteral{token.Valor}, nil
	case lexer.TokenInteiro:
		p.avancar()
		return &InteiroLiteral{token.Valor}, nil
	case lexer.TokenIdentificador:
		if !IsKeyword(token.Valor) {
			p.avancar()
			return &Identificador{token.Valor}, nil
		}
	}

	// fmt.Printf("%t", p.token)
	return nil, fmt.Errorf("O token '%v' não é reconhecido.", p.token.Valor)
}

func (p *Parser) parseEnquanto() (*Enquanto, error) {
	if err := p.consume("enquanto"); err != nil {
		return nil, err
	}

	if err := p.consume("("); err != nil {
		return nil, err
	}

	condicao, err := p.parseExpressao()
	if err != nil {
		return nil, err
	}

	if err := p.consume(")"); err != nil {
		return nil, err
	}

	corpo, err := p.parseBloco()
	if err != nil {
		return nil, err
	}

	return &Enquanto{Condicao: condicao, Corpo: corpo}, nil
}
