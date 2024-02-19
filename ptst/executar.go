package ptst

func ExecutarString(ctx *Contexto, codigo string) (*Modulo, error) {
	ast, err := ctx.StringParaAst(codigo)
	if err != nil {
		return nil, err
	}

	impl := &ModuloImpl{
		Info: ModuloInfo{},
		Ast: ast,
	}

	return ctx.InicializarModulo(impl)
}

func ExecutarArquivo(ctx *Contexto, nome, caminho, curDir string, useSysPaths bool) (*Modulo, error) {
	caminhoCalculado, ast, err := ctx.TransformarEmAst(caminho, useSysPaths, curDir)
	if err != nil {
		return nil, err
	}

	impl := &ModuloImpl{
		Info: ModuloInfo{
			Nome: nome,
			Arquivo: caminhoCalculado,
		},
		Ast: ast,
	}

	return ctx.InicializarModulo(impl)
}