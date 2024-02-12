package compartilhado

import (
	"unicode"
	"unicode/utf8"
)

func ContemApenasLetras(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func ContemApenasDigitos(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func IndiceCaraterParaByte(str string, indice int) int {
	byteIndex := 0
	for i := 0; i < indice; i++ {
		_, tamanho := utf8.DecodeRuneInString(str[byteIndex:])
		byteIndex += tamanho
	}
	return byteIndex
}

func ObtemCaraterPorIndice(str string, indice int) string {
	byteIndex := IndiceCaraterParaByte(str, indice)
	letra, _ := utf8.DecodeRuneInString(str[byteIndex:])
	return string(letra)
}
