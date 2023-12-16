lexer grammar PortuscriptLexer;

FALSO: 'Falso';
VERDADEIRO: 'Verdadeiro';
NULO: 'Nulo';
SE: 'se';
SENAO: 'senao';
VAR: 'var';
CONST: 'const';
IMPORTE: 'importe';
DE: 'de';
RETORNE: 'retorne';
FUNC: 'func';
OU: 'ou';
E: 'e';
NAO: 'nao';

NOVA_LINHA: '\r'? '\n';
DIGITOS:
	'0'
	| '1'
	| '2'
	| '3'
	| '4'
	| '5'
	| '6'
	| '7'
	| '8'
	| '9';
LETRAS: 'A' .. 'Z' | 'a' .. 'z';
OPERADOR_REATRIBUICAO: IGUAL | '+=' | '-=' | '*=' | '@=' | '/=' | '%=' | '&=' | '|=' | '^=' | '<<=' | '>>=' | '**=' | '//=';
IGUAL: '=';
MAIS: '+';
MENOS: '-';
ASTERISCO: '*';
POTENCIA: '**';
DIVISAO: '/';
DIVISAO_INTEIRA: '//';
MODULO: '%';
MENOR_QUE: '<';
MENOR_OU_IGUAL: '<=';
IGUAL_IGUAL: '==';
DIFERENTE: '!=';
MAIOR_QUE: '>';
MAIOR_OU_IGUAL: '>=';
ABRE_PARENTESES: '(';
FECHA_PARENTESES: ')';
PONTO_E_VIRGULA: ';';
VIRGULA: ',';
ABRE_CHAVES: '{';
FECHA_CHAVES: '}';
DOIS_PONTOS: ':';
OU_BIT_A_BIT: '|';
EX_OU_BIT_A_BIT: '^';
E_BIT_A_BIT: '&';
NAO_BIT_A_BIT: '~';
DESLOC_ESQUERDA: '<<';
DESLOC_DIREITA: '>>';
PONTO: '.';

ID: (LETRAS | '_') (LETRAS | DIGITOS)*;

TEXTO: '"' ( ~[\\\r\n"] | ('\\' NOVA_LINHA | '\\' .) )* '"';

WS: [ \t\r\n]+ -> skip;  // Define regra para ignorar espaços em branco.
