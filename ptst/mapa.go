package ptst

type Mapa map[string]Objeto

var TipoMapa = NewTipo(
	"Mapa",
	"Objeto chave/valor",
)

func (m Mapa) Tipo() *Tipo {
	return TipoMapa
}
