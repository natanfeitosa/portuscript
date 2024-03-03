package embutidos

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/ptst"
)

// FIXME: essa provavelmente não é a melhor implementação
func met_emb_leia(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("leia", false, args, 0, 1); err != nil {
		return nil, err
	}

	if len(args) == 1 {
		texto, err := ptst.NewTexto(args[0])

		if err != nil {
			return nil, err
		}

		fmt.Printf("%s", texto)
	}

	var leitura string
	fmt.Scan(&leitura)
	return ptst.Texto(leitura), nil
}

var _emb_leia = ptst.NewMetodoOuPanic(
	"leia",
	met_emb_leia,
	"leia(frase_para_imprimir) -> imprime um texto se especificado e lê uma entrada do usuário, retornando-a",
)
