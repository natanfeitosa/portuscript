package lexer

// TokenType representa os tipos de tokens suportados.
type TokenType int

// Tokens simples
const (
	// Erro léxico
	TokenErro TokenType = iota
	// Fim do arquivo
	TokenFimDeArquivo
	// nova linha. Ex: \n ou \r\n
	TokenNovaLinha

	// Identificadores e literais

	// Ex: variavel
	TokenIdentificador
	// Ex: 123
	TokenInteiro
	// Ex: 123.4
	TokenDecimal
	// Ex: "texto"
	TokenTexto

	// Palavras-chave

	TokenSe
	TokenSenao
	TokenEnquanto
	TokenPara
	TokenRetorne
	TokenPare
	TokenContinue

	TokenDe
	TokenImporte

	TokenVerdadeiro
	TokenFalso
	TokenNulo

	TokenVar
	TokenConst
	TokenFunc

	// Operadores e pontuação

	// =
	TokenIgual
	// +
	TokenMais
	// -
	TokenMenos
	// *
	TokenAsterisco
	// **
	TokenPotencia
	// /
	TokenDivisao
	// //
	TokenDivisaoInteira
	// %
	TokenModulo
	// <
	TokenMenorQue
	// <=
	TokenMenorOuIgual
	// ==
	TokenIgualIgual
	// !=
	TokenDiferente
	// >
	TokenMaiorQue
	// >=
	TokenMaiorOuIgual
	// (
	TokenAbreParenteses
	// )
	TokenFechaParenteses
	// ;
	TokenPontoEVirgula
	// ,
	TokenVirgula
	// {
	TokenAbreChaves
	// }
	TokenFechaChaves
	// [
	TokenAbreColchetes
	// ]
	TokenFechaColchetes
	// :
	TokenDoisPontos

	// Reatribuicao

	// +=
	TokenMaisIgual
	// -=
	TokenMenosIgual
	// *=
	TokenAsteriscoIgual
	// /=
	TokenBarraIgual

	// Operadores booleanos
	TokenBoolOu
	TokenBoolE
	TokenBoolNao

	// |
	TokenBitABitOu
	// ^
	TokenBitABitExOu
	// &
	TokenBitABitE
	// ~
	TokenBitABitNao

	// <<
	TokenDeslocEsquerda
	// >>
	TokenDeslocDireita

	// .
	TokenPonto

	// const instancia = nova MinhaClasse()
	TokenNova
	TokenClasse
)

type PosicaoToken struct {
	Coluna int
	Linha  int
	Indice int
}

type Token struct {
	Tipo  TokenType
	Valor string

	Inicio *PosicaoToken
	Fim    *PosicaoToken
}

func newToken(tipo TokenType, valor string, inicio, fim *PosicaoToken) *Token {
	return &Token{
		tipo,
		valor,
		inicio,
		fim,
	}
}
