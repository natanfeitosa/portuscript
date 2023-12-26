package ptst

// Interfaces para fins de tipagem como a seguinte devem, começar com "I" e seus métodos internos com "O" de "Objeto"

type I__chame__ interface {
	O__chame__(args Tupla) (Objeto, error)
}

type I__texto__ interface {
	O__texto__() (Objeto, error)
}

type I__inteiro__ interface {
	O__inteiro__() (Objeto, error)
}

type I__decimal__ interface {
	O__decimal__() (Objeto, error)
}

type I__booleano__ interface {
	O__booleano__() (Objeto, error)
}

type I_conversaoEntreTipos interface {
	I__texto__
	I__inteiro__
	I__decimal__
	I__booleano__
}

type I__obtem_attributo__ interface {
	O__obtem_attributo__(nome string) (Objeto, error)
}

// FIXME: inst é a instacia; e dono, o objeto vindo do método `Tipo()`
type I__obtem__ interface {
	O__obtem__(inst Objeto, dono *Tipo) (Objeto, error)
}

type I_Mapa interface {
	ObtemMapa() Mapa
}

type I_ObtemDoc interface {
	ObtemDoc() string
}

type I_Chamar interface {
	Chamar(inst Objeto, args Tupla) (Objeto, error)
}

type I__adiciona__ interface {
	O__adiciona__(outro Objeto) (Objeto, error)
}

type I__multiplica__ interface {
	O__multiplica__(outro Objeto) (Objeto, error)
}

type I__subtrai__ interface {
	O__subtrai__(outro Objeto) (Objeto, error)
}

type I__divide__ interface {
	O__divide__(outro Objeto) (Objeto, error)
}

type I__divide_inteiro__ interface {
	O__divide_inteiro__(outro Objeto) (Objeto, error)
}

type I_aritmeticaMatematica interface {
	I__adiciona__
	I__multiplica__
	I__subtrai__
	I__divide__
	I__divide_inteiro__
}

type I__menor_que__ interface {
	O__menor_que__(outro Objeto) (Objeto, error)
}

type I__menor_ou_igual__ interface {
	O__menor_ou_igual__(outro Objeto) (Objeto, error)
}

type I__igual__ interface {
	O__igual__(outro Objeto) (Objeto, error)
}

type I__diferente__ interface {
	O__diferente__(outro Objeto) (Objeto, error)
}

type I__maior_que__ interface {
	O__maior_que__(outro Objeto) (Objeto, error)
}

type I__maior_ou_igual__ interface {
	O__maior_ou_igual__(outro Objeto) (Objeto, error)
}

type I_comparacaoRica interface {
	I__menor_que__
	I__menor_ou_igual__
	I__igual__
	I__diferente__
	I__maior_que__
	I__maior_ou_igual__
}

type I__ou__ interface {
	O__ou__(outro Objeto) (Objeto, error)
}

type I__e__ interface {
	O__e__(outro Objeto) (Objeto, error)
}

type I_aritmeticaBooleana interface {
	I__ou__
	I__e__
}

type I__iter__ interface {
	O__iter__() (Objeto, error)
}

type I__proximo__ interface {
	O__proximo__() (Objeto, error)
}

type I_iterador interface {
	I__iter__
	I__proximo__
}