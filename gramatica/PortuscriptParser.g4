parser grammar PortuscriptParser;

options {
	tokenVocab = PortuscriptLexer;
}

programa: declaracoes? EOF;

declaracoes: declaracao+;

declaracao: declaracao_composta | declaracao_simples;

declaracao_simples:
	atribuicao
	| expressao OPERADOR_REATRIBUICAO expressao
	| expressao
	| declaracao_retorne
	| declaracao_importacao
	| PARE
	| CONTINUE;

declaracao_composta:
	declaracao_funcao
	| declaracao_se
	| declaracao_para;

atribuicao: atribuicao_constante | atribuicao_variavel;

atribuicao_variavel:
	VAR ID (IGUAL expressao | DOIS_PONTOS ID)? PONTO_E_VIRGULA;

atribuicao_constante: CONST ID IGUAL expressao PONTO_E_VIRGULA;

declaracao_importacao:
	declaracao_importacao_simples
	| declaracao_importacao_de;

declaracao_importacao_simples:
	IMPORTE TEXTO (VIRGULA TEXTO)* PONTO_E_VIRGULA;

declaracao_importacao_de:
	DE TEXTO IMPORTE (ASTERISCO | ID (VIRGULA ID)*) PONTO_E_VIRGULA;

declaracao_retorne: RETORNE expressao? PONTO_E_VIRGULA;

declaracao_funcao:
	FUNC ID ABRE_PARENTESES parametros? FECHA_PARENTESES bloco;

parametros: parametro (VIRGULA parametro)*;

parametro: ID (':' ID)?;

declaracao_se:
	SE ABRE_PARENTESES expressao FECHA_PARENTESES bloco (
		declaracao_senao_se
		| SENAO bloco
	)?;

declaracao_senao_se:
	SENAO SE ABRE_PARENTESES expressao FECHA_PARENTESES bloco (
		declaracao_senao_se
		| SENAO bloco
	)?;

declaracao_senao: SENAO bloco;

declaracao_para: PARA ID EM primario;

bloco: ABRE_CHAVES declaracoes? FECHA_CHAVES;

expressao: disjuncao;

disjuncao: conjuncao (OU conjuncao)*;

conjuncao: inversao (E inversao)*;

inversao: NAO inversao | comparacao;

comparacao:
	ou_bitabit (
		(
			IGUAL_IGUAL
			| DIFERENTE
			| MENOR_QUE
			| MENOR_OU_IGUAL
			| MAIOR_QUE
			| MAIOR_OU_IGUAL
		) ou_bitabit
	)?;

ou_bitabit: exou_bitabit (OU_BIT_A_BIT exou_bitabit)*;

exou_bitabit: e_bitabit (EX_OU_BIT_A_BIT e_bitabit)*;

e_bitabit: deslocamento (E_BIT_A_BIT deslocamento)*;

deslocamento:
	arit_basica ((DESLOC_ESQUERDA | DESLOC_DIREITA) arit_basica)?;

arit_basica: termo ((MAIS | MENOS) termo)*;

termo:
	fator (
		(ASTERISCO | DIVISAO | DIVISAO_INTEIRA | MODULO) fator
	)*;

fator: (MAIS | MENOS | NAO_BIT_A_BIT)* potencia;

potencia: primario (POTENCIA fator)?;

primario:
	primario PONTO primario
	| primario ABRE_PARENTESES argumentos FECHA_PARENTESES
	| primario ABRE_COLCHETES expressao FECHA_COLCHETES
	| atomo;

argumentos: argumento (VIRGULA argumento)*;

argumento: expressao;

atomo:
	ID
	| 'Verdadeiro'
	| 'Falso'
	| 'Nulo'
	| TEXTO+
	| DIGITOS
	| tupla
	| grupo
	| lista;

tupla:
	ABRE_CHAVES expressao VIRGULA (expressao VIRGULA?)* FECHA_CHAVES;

grupo: ABRE_CHAVES expressao FECHA_CHAVES;

lista:
	ABRE_COLCHETES expressao (VIRGULA expressao)* FECHA_COLCHETES;