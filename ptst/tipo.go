package ptst

import (
	"fmt"
	"strings"
)

type CriaFunc func(meta *Tipo, args Tupla) (Objeto, error)

// func __inicializa__(instancia, argumentos) {}
type InicializaFunc func(inst Objeto, args Tupla) error

// Talvez um MRO seria útil?
type Tipo struct {
	Nome       string         // Nome para o tipo
	Cria       CriaFunc       // Semelhante ao __new__ do python
	Inicializa InicializaFunc // Funcao/Metodo chamado quando instanciar uma classe
	Doc        string         // A documentaçao para ajudar (ou não) a entender a classe
	Base       *Tipo          // A classe da qual a atual herda
	Mapa       Mapa
	// Abstrata // Talvez sinalize que a classe nao deve ser instanciada
}

func NewTipo(nome string, doc string) *Tipo {
	t := &Tipo{Nome: nome, Doc: doc, Mapa: Mapa{}}
	EnfileiraMontagemDoTipo(t)
	return t
}

func (b *Tipo) ObtemDoc() string {
	return b.Doc
}

func (b *Tipo) NewTipo(nome string, doc string) *Tipo {
	t := &Tipo{Nome: nome, Doc: doc, Base: b, Mapa: Mapa{}}
	EnfileiraMontagemDoTipo(t)
	return t
}

func (b *Tipo) NewTipoX(nome string, doc string, cria CriaFunc, inicializa InicializaFunc) *Tipo {
	t := &Tipo{Nome: nome, Doc: doc, Base: b, Cria: cria, Inicializa: inicializa, Mapa: Mapa{}}
	EnfileiraMontagemDoTipo(t)
	return t
}

func (b *Tipo) ObtemMapa() Mapa {
	return b.Mapa
}

func (b *Tipo) Monta() error {
	// FIXME: adicionar aqui as questões de classe base, heranças internas, valores padrão obrigatórios e coisas do tipo

	if b.Mapa == nil {
		b.Mapa = Mapa{}
	}

	if _, ok := b.Mapa["__doc__"]; !ok {
		if b.Doc != "" {
			b.Mapa["__doc__"] = Texto(strings.Trim(b.Doc, "\r\n\t "))
		} else {
			b.Mapa["__doc__"] = Nulo
		}
	}
	
	return nil
}

// G vem de `Genérico`
func (b *Tipo) G_ObtemAtributoOuNil(nome string) Objeto {
	if obj, ok := b.Mapa[nome]; ok {
		return obj
	}

	// FIXME: não deviamos olhar nas bases?
	return nil
}

var TipoTipo *Tipo = NewTipo(
	"Tipo",
	"Tipo raiz para todos os objetos (interno).",
)

var TipoObjeto *Tipo = NewTipo(
	"Objeto",
	"A classe base para todas as outras classes.",
)

func init() {
	// FIXME: faça algo com os erros eventuais
	TipoTipo.Monta()
	TipoObjeto.Monta()
}

var filaMontagem []*Tipo

func EnfileiraMontagemDoTipo(tipo *Tipo) {
	filaMontagem = append(filaMontagem, tipo)
}

func MontaOsTipos() error {
	for _, tipo := range filaMontagem {
		err := tipo.Monta()

		if err != nil {
			return fmt.Errorf("Erro ao montar o tipo %s: %v", tipo.Nome, err)
		}

		filaMontagem = nil
	}

	return nil
}

func init() {
	err := MontaOsTipos()

	if err != nil {
		// FIXME: talvez não seja a melhor abordagem
		panic(err)
	}
}