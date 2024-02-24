package colorize

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/natanfeitosa/portuscript/ptst"
)

var regexCor = regexp.MustCompile(`\033\[[\d;?]+m`)

func met_color_imprimac(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if len(args) == 0 {
		return nil, ptst.NewErro(ptst.TipagemErro, ptst.Texto("A função imprimac esperava receber ao menos 1 argumento"))
	}

	var junta, textoObj ptst.Objeto
	var err error

	if junta, err = ptst.ObtemAtributoS(ptst.Texto(""), "junta"); err != nil {
		return nil, err
	}
	
	if textoObj, err = ptst.Chamar(junta, args);err != nil {
		return nil, err
	}

	saida := string(textoObj.(ptst.Texto))
	
	if !SuportaCores && strings.Contains(saida, InicioCodigo) {
		saida = regexCor.ReplaceAllString(saida, "")
	}
	
	fmt.Println(saida)
	return nil, nil
}

var _color_imprimac = ptst.NewMetodoOuPanic(
	"imprimac",
	met_color_imprimac,
	"imprimac(...objeto) -> O mesmo que a função embutida imprima, porém mais apto a trabalhar com as cores",
)
