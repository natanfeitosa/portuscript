package lexer

// TokenType representa os tipos de tokens suportados.
type TokenType int

const (
	// Tokens simples
	TokenEOF       TokenType = iota // Fim do arquivo
	TokenErro                       // Erro léxico
	TokenNovaLinha                  // \n

	// Identificadores e literais
	TokenIdentificador // Ex: variavel
	TokenInteiro       // Ex: 123
	TokenDecimal       // Ex: 123.4
	TokenTexto         // Ex: "texto"

	// Palavras-chave
	TokenSe
	TokenSenao
	TokenEnquanto
	TokenPara
	TokenRetorne

	TokenVerdadeiro // Verdadeiro
	TokenFalso      // Falso
	TokenNulo       // Nulo

	// Operadores e pontuação
	TokenIgual           // =
	TokenMais            // +
	TokenMenos           // -
	TokenAsterisco       // *
	TokenPotencia        // **
	TokenDivisao         // /
	TokenDivisaoInteira  // //
	TokenModulo          // %
	TokenMenorQue        // <
	TokenMenorOuIgual    // <=
	TokenIgualIgual      // ==
	TokenDiferente       // !=
	TokenMaiorQue        // >
	TokenMaiorOuIgual    // >=
	TokenAbreParenteses  // (
	TokenFechaParenteses // )
	TokenPontoEVirgula   // ;
	TokenVirgula         // ,
	TokenAbreChaves      // {
	TokenFechaChaves     // }
	TokenDoisPontos      // :

	// Reatribuicao
	TokenMaisIgual      // +=
	TokenMenosIgual     // -=
	TokenAsteriscoIgual // *=
	TokenBarraIgual     // /=

	// Operadores booleanos
	TokenBoolOu
	TokenBoolE
	TokenBoolNao

	TokenBitABitOu   // |
	TokenBitABitExOu // ^
	TokenBitABitE    // &
	TokenBitABitNao  // ~

	TokenDeslocEsquerda // <<
	TokenDeslocDireita  // >>

	TokenPonto // .
)

// Token representa um token com seu tipo e valor.
type Token struct {
	Tipo  TokenType
	Valor string
}
