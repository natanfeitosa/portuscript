package playground

import (
	"strings"
)

type indicador string

const (
	Normal   indicador = ">>> "
	Continua indicador = "... "
)

type Estado struct {
	Indicador indicador
	Continua bool
	Codigo string
}

func NewEstado() *Estado {
	return &Estado{Normal, false, ""}
}

func (e *Estado) RecalcularEstado(cod string) {
	e.Codigo += cod
	continua := e.continuaEmNovaLinha("[", "]") || e.continuaEmNovaLinha("(", ")") || e.continuaEmNovaLinha("{", "}")

	if continua == e.Continua {
		return
	}
	
	e.Continua = continua
	if e.Continua {
		e.Indicador = Continua
		return
	}

	e.Indicador = Normal
}

func (e *Estado) continuaEmNovaLinha(abre, fecha string) bool {
	return strings.Count(e.Codigo, abre) > strings.Count(e.Codigo, fecha)
}