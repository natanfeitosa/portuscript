package ptst

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
	return &Tipo{Nome: nome, Doc: doc, Mapa: Mapa{}}
}

func (b *Tipo) ObtemDoc() string {
	return b.Doc
}

func (b *Tipo) NewTipo(nome string, doc string) *Tipo {
	return &Tipo{Nome: nome, Doc: doc, Base: b, Mapa: Mapa{}}
}

func (b *Tipo) NewTipoX(nome string, doc string, cria CriaFunc, inicializa InicializaFunc) *Tipo {
	return &Tipo{Nome: nome, Doc: doc, Base: b, Cria: cria, Inicializa: inicializa, Mapa: Mapa{}}
}

var TipoTipo *Tipo = NewTipo(
	"Tipo",
	"Tipo raiz para todos os objetos (interno).",
)

var TipoObjeto *Tipo = NewTipo(
	"Objeto",
	"A classe base para todas as outras classes.",
)
