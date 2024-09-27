package embutidos

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	reader := bufio.NewReader(os.Stdin)
	leitura, err := reader.ReadString('\n')
	if err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao ler a entrada: %v", err)
	}
	return ptst.Texto(strings.TrimRight(leitura, "\n")), nil
}

var _emb_leia = ptst.NewMetodoOuPanic(
	"leia",
	met_emb_leia,
	"leia(frase_para_imprimir) -> imprime um texto se especificado e lê uma entrada do usuário, retornando-a",
)
