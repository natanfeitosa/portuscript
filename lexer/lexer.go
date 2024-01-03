package lexer

// As palavras chaves devem estar no imperativo sempre que possivel

import (
	"strings"

	"github.com/natanfeitosa/portuscript/compartilhado"
)

// Lexer representa o analisador léxico.
type Lexer struct {
	entrada   string
	posicao   int
	leitura   int
	caractere string
}

// NewLexer cria uma nova instância de Lexer.
func NewLexer(input string) *Lexer {
	l := &Lexer{
		entrada: strings.Trim(input, " \n") + "\n",
		// posicao: 0,
	}
	l.avancar() // Inicializa o caractere atual.
	return l
}

func (l *Lexer) isEof() bool {
	return l.leitura >= len(l.entrada)
}

func (l *Lexer) posicaoReal(idx int) int {
	return compartilhado.IndiceCaractereParaByte(l.entrada, idx)
}

func (l *Lexer) subString(inicio, fim int) string {
	return l.entrada[l.posicaoReal(inicio):l.posicaoReal(fim)]
}

// Avançar lê o próximo caractere da entrada.
func (l *Lexer) avancar() {
	if !l.isEof() {
		l.caractere = compartilhado.ObtemLetraPorIndice(l.entrada, l.posicao)
		l.posicao++
		l.leitura += len(l.caractere)
	}

	// fmt.Printf("\npos: %d\ncar: %s\n", l.posicao, l.caractere)
}

func (l *Lexer) ignorarEspacos() {
	for (l.caractere == " " || l.caractere == "\t") && !l.isEof() {
		l.avancar()
	}
}

func (l *Lexer) ignorarComentarios() {
	// l.avancar() // Avança sobre o cerquilha?
	for (l.caractere != "\n") && !l.isEof() {
		l.avancar()
	}
	l.avancar()
}

func (l *Lexer) lerIdentificador() string {
	inicio := l.posicao - 1
	for compartilhado.ContemApenasLetras(l.caractere) || compartilhado.ContemApenasDigitos(l.caractere) || l.caractere == "_" {
		if l.isEof() {
			break
		}

		l.avancar()
	}

	return l.subString(inicio, l.posicao-1)
}

func (l *Lexer) lerNumero() string {
	inicio := l.posicao - 1

	for compartilhado.ContemApenasDigitos(l.caractere) || strings.ContainsAny(".+-", l.caractere) {
		l.avancar()
	}

	return l.subString(inicio, l.posicao-1)
}

func (l *Lexer) lerTexto() string {
	inicio := l.posicao - 1
	abertura := l.caractere

	for {
		l.avancar()

		if l.isEof() || l.caractere == abertura {
			// FIXME: lança um erro se for EOF (Fim do arquivo)
			break
		}
	}

	l.avancar() // Avança além da última aspa dupla.
	return l.subString(inicio, l.posicao-1)
}

// FIXME: refatorar para deixar menor
func (l *Lexer) ProximoToken() Token {
	if l.isEof() {
		return Token{TokenEOF, ""}
	}
	l.ignorarEspacos()

	switch l.caractere {
	case "\n":
		l.avancar()
		return Token{TokenNovaLinha, "\n"}
	// FIXME: isso com certeza não é a melhor das ideias
	case "#":
		l.ignorarComentarios()
		return l.ProximoToken()
	case ";":
		l.avancar()
		return Token{TokenPontoEVirgula, ";"}
	case ":":
		l.avancar()
		return Token{TokenDoisPontos, ":"}
	case ",":
		l.avancar()
		return Token{TokenVirgula, ","}
	case "=":
		l.avancar()
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenIgualIgual, "=="}
		}
		return Token{TokenIgual, "="}
	case "+":
		if compartilhado.ContemApenasDigitos(l.peekProximoCaractere()) {
			goto LERNUMERO
		}
		
		l.avancar()
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenMaisIgual, "+="}
		}
		return Token{TokenMais, "+"}
	case "-":
		if compartilhado.ContemApenasDigitos(l.peekProximoCaractere()) {
			goto LERNUMERO
		}

		l.avancar()
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenMenosIgual, "-="}
		}
		return Token{TokenMenos, "-"}
	case "*":
		l.avancar()
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenAsteriscoIgual, "*="}
		}

		if l.caractere == "*" {
			l.avancar()
			return Token{TokenPotencia, "**"}
		}

		return Token{TokenAsterisco, "*"}
	case "/":
		l.avancar()
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenBarraIgual, "/="}
		}

		if l.caractere == "/" {
			l.avancar()
			return Token{TokenDivisaoInteira, "//"}
		}

		return Token{TokenDivisao, "/"}
	case "!":
		l.avancar()
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenDiferente, "!="}
		}
		// FIXME: deve lançar um erro aqui
	case ">":
		l.avancar()
		if l.caractere == ">" {
			l.avancar()
			return Token{TokenDeslocDireita, ">>"}
		}
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenMaiorOuIgual, ">="}
		}
		return Token{TokenMaiorQue, ">"}
	case "<":
		l.avancar()
		if l.caractere == "<" {
			l.avancar()
			return Token{TokenDeslocEsquerda, "<<"}
		}
		if l.caractere == "=" {
			l.avancar()
			return Token{TokenMenorOuIgual, "<="}
		}
		return Token{TokenMenorQue, "<"}
	case "(":
		l.avancar()
		return Token{TokenAbreParenteses, "("}
	case ")":
		l.avancar()
		return Token{TokenFechaParenteses, ")"}
	case "{":
		l.avancar()
		return Token{TokenAbreChaves, "{"}
	case "}":
		l.avancar()
		return Token{TokenFechaChaves, "}"}
	case "[":
		l.avancar()
		return Token{TokenAbreColchetes, "["}
	case "]":
		l.avancar()
		return Token{TokenFechaColchetes, "]"}
	case "|":
		l.avancar()
		return Token{TokenBitABitOu, "|"}
	case "^":
		l.avancar()
		return Token{TokenBitABitExOu, "^"}
	case "&":
		l.avancar()
		return Token{TokenBitABitE, "&"}
	case "~":
		l.avancar()
		return Token{TokenBitABitNao, "~"}
	case "\"":
		return Token{TokenTexto, l.lerTexto()}
	case ".":
		l.avancar()
		return Token{TokenPonto, "."}
	default:
		if compartilhado.ContemApenasLetras(l.caractere) || l.caractere == "_" {
			identificador := l.lerIdentificador()
			switch strings.ToLower(identificador) {
			case "se":
				return Token{TokenSe, identificador}
			case "senao":
				return Token{TokenSenao, identificador}
			case "enquanto":
				return Token{TokenEnquanto, identificador}
			case "para":
				return Token{TokenPara, identificador}
			case "retorne":
				return Token{TokenRetorne, identificador}
			case "ou":
				return Token{TokenBoolOu, identificador}
			case "e":
				return Token{TokenBoolE, identificador}
			case "nao":
				return Token{TokenBoolNao, identificador}
			case "Verdadeiro":
				return Token{TokenVerdadeiro, identificador}
			case "Falso":
				return Token{TokenFalso, identificador}
			case "Nulo":
				return Token{TokenNulo, identificador}
			case "pare":
				return Token{TokenPare, identificador}
			case "continue":
				return Token{TokenContinue, identificador}
			default:
				return Token{TokenIdentificador, identificador}
			}
		} else if compartilhado.ContemApenasDigitos(l.caractere) {
			goto LERNUMERO
		}
	}

	return Token{TokenErro, l.caractere}

LERNUMERO:
	{
		numero := l.lerNumero()

		if strings.Contains(numero, ".") {
			return Token{TokenDecimal, numero}
		}

		return Token{TokenInteiro, numero}
	}
}

// peekProximoCaractere retorna o próximo caractere sem avançar o analisador.
func (l *Lexer) peekProximoCaractere() string {
	if !l.isEof() {
		return string(l.entrada[l.posicao+1])
	}

	return ""
}
