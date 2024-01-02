package ptst

func MaquinarioImporteModulo(ctx *Contexto, nome string) (Objeto, error) {
	if modulo, err := ctx.ObterModulo(nome); err == nil {
		return modulo, nil
	}

	if impl := ObtemImplModulo(nome); impl != nil {
		return ctx.InicializarModulo(impl)
	}

	// FIXME: adicionar lógica para aquisição e inicialização de métodos "locais"

	return nil, NewErroF(ImportacaoErro, "Não foi possível achar o módulo '%s', talvez seja um módulo local que ainda não é suportado", nome)
}

func MultiImporteModulo(ctx *Contexto, nomes ...string) error {
	for _, nome := range nomes {
		if _, err := MaquinarioImporteModulo(ctx, nome); err != nil {
			return err
		}
	}

	return nil
}
