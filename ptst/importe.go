package ptst

import "path"

func MaquinarioImporteModulo(ctx *Contexto, nome string, escopo *Escopo) (Objeto, error) {
	if modulo, err := ctx.ObterModulo(nome); err == nil {
		return modulo, nil
	}

	if impl := ObtemImplModulo(nome); impl != nil {
		return ctx.InicializarModulo(impl)
	}

	// FIXME: lidar com importações começando em `/` como sendo uma importação desde a raiz do cwd
	// FIXME: prevenir importação circular
	// FIXME: adicionar operador para definir quem deve ser público para importar
	curDir := ""
	if escopo != nil {
		arqAtual, err := escopo.ObterValor("__arquivo__")
		if err == nil {
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
