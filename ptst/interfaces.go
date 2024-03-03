package ptst

// Interfaces para fins de tipagem como a seguinte devem, começar com "I" e seus métodos internos com "M" de "Método"

type I__chame__ interface {
	M__chame__(args Tupla) (Objeto, error)
}

type I__texto__ interface {
	M__texto__() (Objeto, error)
}

type I__inteiro__ interface {
	M__inteiro__() (Objeto, error)
}

type I__decimal__ interface {
	M__decimal__() (Objeto, error)
}

type I__booleano__ interface {
	M__booleano__() (Objeto, error)
}

type I_conversaoEntreTipos interface {
	I__texto__
	I__inteiro__
	I__decimal__
	I__booleano__
}

type I__obtem_attributo__ interface {
	M__obtem_attributo__(nome string) (Objeto, error)
}

// FIXME: inst é a instacia; e dono, o objeto vindo do método `Tipo()`
type I__obtem__ interface {
	M__obtem__(inst Objeto, dono *Tipo) (Objeto, error)
}

type I_ObtemMapa interface {
	ObtemMapa() Mapa
}

type I_ObtemDoc interface {
	ObtemDoc() string
}

type I_Chamar interface {
	Chamar(inst Objeto, args Tupla) (Objeto, error)
}

type I__adiciona__ interface {
	M__adiciona__(outro Objeto) (Objeto, error)
}

type I__multiplica__ interface {
	M__multiplica__(outro Objeto) (Objeto, error)
}

type I__subtrai__ interface {
	M__subtrai__(outro Objeto) (Objeto, error)
}

type I__divide__ interface {
	M__divide__(outro Objeto) (Objeto, error)
}

type I__divide_inteiro__ interface {
	M__divide_inteiro__(outro Objeto) (Objeto, error)
}

type I__mod__ interface {
	M__mod__(outro Objeto) (Objeto, error)
}

type I__neg__ interface {
	M__neg__() (Objeto, error)
}

type I__pos__ interface {
	M__pos__() (Objeto, error)
}

type I_aritmeticaMatematica interface {
	I__adiciona__
	I__multiplica__
	I__subtrai__
	I__divide__
	I__divide_inteiro__
	I__mod__
	I__neg__
	I__pos__
}

type I__menor_que__ interface {
	M__menor_que__(outro Objeto) (Objeto, error)
}

type I__menor_ou_igual__ interface {
	M__menor_ou_igual__(outro Objeto) (Objeto, error)
}

type I__igual__ interface {
	M__igual__(outro Objeto) (Objeto, error)
}

type I__diferente__ interface {
	M__diferente__(outro Objeto) (Objeto, error)
}

type I__maior_que__ interface {
	M__maior_que__(outro Objeto) (Objeto, error)
}

type I__maior_ou_igual__ interface {
	M__maior_ou_igual__(outro Objeto) (Objeto, error)
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
	M__ou__(outro Objeto) (Objeto, error)
}

type I__e__ interface {
	M__e__(outro Objeto) (Objeto, error)
}

type I_aritmeticaBooleana interface {
	I__ou__
	I__e__
}

type I__iter__ interface {
	M__iter__() (Objeto, error)
}

type I__proximo__ interface {
	M__proximo__() (Objeto, error)
}

type I_iterador interface {
	I__iter__
	I__proximo__
}

type I__tamanho__ interface {
	M__tamanho__() (Objeto, error)
}

type I__obtem_item__ interface {
	M__obtem_item__(obj Objeto) (Objeto, error)
}

type I__define_item__ interface {
	M__define_item__(chave, valor Objeto) (Objeto, error)
}

// Semelhante ao __new__ do python
type I__nova_instancia__ interface {
	M__nova_instancia__(meta *Tipo, args Tupla) (Objeto, error)
}

type I__inicializa__ interface {
	M__inicializa__(args Tupla) (error)
}