package ptst

import (
	"os"
	"path"
	"strings"
)

func MaquinarioImporteModulo(ctx *Contexto, nome string, escopo *Escopo) (Objeto, error) {
	if modulo, err := ctx.ObterModulo(nome); err == nil {
		return modulo, nil
	}

	if impl := ObtemImplModulo(nome); impl != nil {
		return ctx.InicializarModulo(impl)
	}

	if !(strings.HasPrefix(nome, "./") || strings.HasPrefix(nome, "/")) {
		return nil, NewErroF(ImportacaoErro, "Importações não relativas só estão disponíveis para módulos embutidos, corrija para './%s'", nome)
	}

	// FIXME: prevenir importação circular
	// FIXME: adicionar operador para definir quem deve ser público para importar
	curDir := ""

	if strings.HasPrefix(nome, "/") {
		// FIXME: vamos apenas ignorar o erro?
		curDir, _ = os.Getwd()
	} else if strings.HasPrefix(nome, "./") {
		if escopo == nil {
			panic("Um escopo atual é necessário quando usar importação relativa do tipo './modulo'")
		}

		if arqAtual, err := escopo.ObterValor("__arquivo__"); err == nil {
			curDir = path.Dir(string(arqAtual.(Texto)))
		}
	}

	mod, err := ExecutarArquivo(ctx, nome, nome, curDir, true)
	if err != nil {
		return nil, err
	}

	return mod, nil
}

func MultiImporteModulo(ctx *Contexto, nomes ...string) error {
	for _, nome := range nomes {
		if _, err := MaquinarioImporteModulo(ctx, nome, nil); err != nil {
			return err
		}
	}

	return nil
}

var Importe func(string, *Escopo) (Objeto, error) = func(s string, e *Escopo) (Objeto, error) {
	panic("Antes de usar a função `Importe` você precisa criar um contexto")
}
