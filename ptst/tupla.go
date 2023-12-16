package ptst

type Tupla []Objeto

var TipoTupla = TipoObjeto.NewTipo(
	"Tupla",
	"Tupla(obj) -> Tupla",
)

func (t Tupla) Tipo() *Tipo {
	return TipoTupla
}
