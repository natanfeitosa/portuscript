package embutidos

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/ptst"
)

func met_emb_imprima(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	const (
		final     = ptst.Texto("\n")
		separador = ptst.Texto(" ")
	)

	junta, err := ptst.ObtemItemS(separador, "junta")

	if err != nil {
		return nil, err
	}

	// resultado, err := ptst.Chamar(junta, args)
	resultado, err := ptst.Chamar(
		junta,
		args,
	)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%s%s", resultado, final)
	return nil, nil
}

var _emb_imprima = ptst.NewMetodoOuPanic(
	"imprima",
	met_emb_imprima,
	"imprima(...objetos) -> imprime a representação ou a conversão em string dos objetos separados por espaço",
)
