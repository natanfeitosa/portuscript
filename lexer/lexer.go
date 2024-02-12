package lexer

import (
	"strings"

	"github.com/natanfeitosa/portuscript/compartilhado"
	"github.com/rivo/uniseg"
)

type Lexer struct {
	entrada string
	carater string
	coluna  int
	linha   int

	// campos de apoio
	tamanho int
	indice  int
}

func NewLexer(entrada string) *Lexer {
	// entrada = strings.Trim(entrada, " \n")
	l := &Lexer{
		entrada: entrada,
		// carater: compartilhado.ObtemCaraterPorIndice(input, 0),
		// coluna:  1,
		// linha:   1,
		tamanho: uniseg.GraphemeClusterCount(entrada),
		indice:  -1,
	}

	l.avancar()

	return l
}

// Verifica se já está no final do arquivo
func (l *Lexer) fimDeArquivo() bool {
	return l.indice >= l.tamanho
}

// Retorna o próximo carater ou uma string vazia se for fim de arquivo, mas não altera o estado do lexer
func (l *Lexer) proximoCarater() string {
	if l.fimDeArquivo() {
		return ""
	}

	return compartilhado.ObtemCaraterPorIndice(l.entrada, l.indice+1)
}

// Atualiza o estado do lexer para o novo carater se não for fim de arquivo
func (l *Lexer) avancar() {
	if l.fimDeArquivo() {
		return
	}

	l.indice += 1
	l.carater = compartilhado.ObtemCaraterPorIndice(l.entrada, l.indice)

	if l.carater == "\n" {
		l.linha += 1
		l.coluna = 0
		return
	}

	l.coluna += 1
}

func (l *Lexer) posicaoAtual() *PosicaoToken {
	return &PosicaoToken{l.coluna, l.linha, l.indice}
}

func (l *Lexer) ignorarEspacos() {
	for (l.carater == " " || l.carater == "\t") && !l.fimDeArquivo() {
		l.avancar()
	}
}

func (l *Lexer) ignorarComentario() {
	for (l.carater != "\n") && !l.fimDeArquivo() {
		l.avancar()
	}
	l.avancar()
}

func (l *Lexer) subString(inicio, fim int) string {
	return l.entrada[compartilhado.IndiceCaraterParaByte(l.entrada, inicio):compartilhado.IndiceCaraterParaByte(l.entrada, fim)]
}

func (l *Lexer) lerIdentificador() *Token {
	inicio := l.posicaoAtual()

	for {
		l.avancar()

		if !(compartilhado.ContemApenasAlfaNum(l.carater) || l.carater == "_") {
			break
		}
	}

	fim := l.posicaoAtual()
	valor := l.subString(inicio.Indice, fim.Indice)
	tipo := TokenIdentificador

	if t, ok := tokensIdentificadores[valor]; ok {
		tipo = t
	}

	return newToken(tipo, valor, inicio, fim)
}

func (l *Lexer) lerNumero() *Token {
	inicio := l.posicaoAtual()

	for {
		l.avancar()

		if !(compartilhado.ContemApenasDigitos(l.carater) || l.carater == ".") {
			break
		}
	}

	fim := l.posicaoAtual()
	valor := l.subString(inicio.Indice, fim.Indice)

	tipo := TokenInteiro
	if strings.Contains(valor, ".") {
		tipo = TokenDecimal
	}

	return newToken(tipo, valor, inicio, fim)
}

func (l *Lexer) lerTexto() *Token {
	inicio := l.posicaoAtual()
	delimitador := l.carater

	for {
		l.avancar()

		if l.carater == delimitador {
			l.avancar()
			break
		}

		if l.carater == "\\" && l.proximoCarater() == delimitador {
			l.avancar()
		}
	}

	fim := l.posicaoAtual()
	return newToken(TokenTexto, l.subString(inicio.Indice, fim.Indice), inicio, fim)
}

func (l *Lexer) ProximoToken() *Token {
	l.ignorarEspacos()

	if l.fimDeArquivo() {
		return &Token{
			Tipo:   TokenFimDeArquivo,
			Valor:  "",
			Inicio: l.posicaoAtual(),
			Fim:    l.posicaoAtual(),
		}
	}

	carater := l.carater
	inicio := l.posicaoAtual()
	if tipo, ok := tokensSimples[carater]; ok || carater == "!" {
		for {
			l.avancar()
			if t, ok := tokensSimples[carater+l.carater]; ok {
				carater += l.carater
				tipo = t
				continue
			}

			break
		}

		return newToken(tipo, carater, inicio, l.posicaoAtual())
	}

	if carater == "#" {
		l.ignorarComentario()
		return l.ProximoToken()
	}

	switch carater {
	case "\"", "'":
		return l.lerTexto()
	// case "/*":
	// 	return l.lerAnotacao()
	default:
		if compartilhado.ContemApenasLetras(carater) || carater == "_" {
			return l.lerIdentificador()
		}

		if compartilhado.ContemApenasDigitos(carater) {
			return l.lerNumero()
		}
	}
	return &Token{Tipo: TokenErro, Valor: l.carater}
}
