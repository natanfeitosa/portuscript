package ptst

type Mapa map[string]Objeto

var TipoMapa = NewTipo(
	"Mapa",
	"Objeto chave/valor",
)

func NewMapaVazio() Mapa {
	return make(Mapa)
}

func (m Mapa) Tipo() *Tipo {
	return TipoMapa
}
