package colorize

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/natanfeitosa/portuscript/ptst"
)

type cor struct {
	Nome string
	// Doc  string
	Hex string
}

func criaRenderizadorDeCores(r, g, b ptst.Inteiro, background bool) func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	codigo := InicioCodigo + RgbParaAnsi(r, g, b, background) + FimCodigo

	return func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		saida := codigo
		for _, item := range args {
			itemObj, err := ptst.NewTexto(item)
			if err != nil {
				return nil, err
			}

			saida = fmt.Sprintf("%s%s", saida, itemObj)
		}

		return ptst.Texto(saida + ResetCodigo), nil
	}
}


func HexParaRgb(hex string) (r, g, b ptst.Inteiro, err error) {
	hex = strings.TrimSpace(hex)
	if hex == "" {
		err = fmt.Errorf("O código hex não pode ser vazio")
		return
	}

	// like from css. eg "#ccc" "#ad99c0"
	if hex[0] == '#' {
		hex = hex[1:]
	}

	hex = strings.ToLower(hex)
	switch len(hex) {
	case 3: // "ccc"
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	case 8: // "0xad99c0"
		hex = strings.TrimPrefix(hex, "0x")
	}

	// recheck
	if len(hex) != 6 {
		err = fmt.Errorf("O código '%s' não segue um formato válido de cor hex", hex)
		return
	}

	// convert string to int64
	if i64, err := strconv.ParseInt(hex, 16, 32); err == nil {
		color := int(i64)
		// parse int
		r = ptst.Inteiro(color >> 16)
		g = ptst.Inteiro((color & 0x00FF00) >> 8)
		b = ptst.Inteiro(color & 0x0000FF)
	}
	return
}

func RgbParaAnsi(r, g, b ptst.Inteiro, background bool) string {
	if background {
		return fmt.Sprintf(TplBgRGB, r, g, b)
	}

	return fmt.Sprintf(TplFgRGB, r, g, b)
}