package lexer

var tokensSimples = map[string]TokenType{
	"\n": TokenNovaLinha,

	"=":  TokenIgual,
	"+":  TokenMais,
	"-":  TokenMenos,
	"*":  TokenAsterisco,
	"**": TokenPotencia,
	"/":  TokenDivisao,
	"//": TokenDivisaoInteira,
	"%":  TokenModulo,
	"<":  TokenMenorQue,
	"<=": TokenMenorOuIgual,
	"==": TokenIgualIgual,
	"!=": TokenDiferente,
	">":  TokenMaiorQue,
	">=": TokenMaiorOuIgual,
	"(":  TokenAbreParenteses,
	")":  TokenFechaParenteses,
	";":  TokenPontoEVirgula,
	",":  TokenVirgula,
	"{":  TokenAbreChaves,
	"}":  TokenFechaChaves,
	"[":  TokenAbreColchetes,
	"]":  TokenFechaColchetes,
	":":  TokenDoisPontos,

	// Reatribuicao
	"+=": TokenMaisIgual,
	"-=": TokenMenosIgual,
	"*=": TokenAsteriscoIgual,
	"/=": TokenBarraIgual,
	"//=": TokenBarraBarraIgual,

	"|":  TokenBitABitOu,
	"^":  TokenBitABitExOu,
	"&":  TokenBitABitE,
	"~":  TokenBitABitNao,
	"<<": TokenDeslocEsquerda,
	">>": TokenDeslocDireita,

	".": TokenPonto,
}

var tokensIdentificadores = map[string]TokenType{
	"se":       TokenSe,
	"senao":    TokenSenao,
	"enquanto": TokenEnquanto,
	"para":     TokenPara,
	"retorne":  TokenRetorne,
	"pare":     TokenPare,
	"continue": TokenContinue,

	"de":      TokenDe,
	"importe": TokenImporte,

	"Verdadeiro": TokenVerdadeiro,
	"Falso":      TokenFalso,
	"Nulo":       TokenNulo,

	"var":   TokenVar,
	"const": TokenConst,
	"func":  TokenFunc,

	// Operadores booleanos

	"ou":  TokenBoolOu,
	"e":   TokenBoolE,
	"nao": TokenBoolNao,

	"nova":   TokenNova,
	"classe": TokenClasse,

	"assegura": TokenAssegura,

	"em": TokenEm,
}
