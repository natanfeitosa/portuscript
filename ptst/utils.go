package ptst

func MesmoTipo(a, b Objeto) bool {
	return a.Tipo() == b.Tipo()
}

func VerificaNumeroArgumentos(nome string, ehMetodo bool, args Objeto, min, max int) error {
	numArgs := len(args.(Tupla))

	if numArgs < min || numArgs > max {
		tipo := "a função"
		if ehMetodo {
			tipo = "o método"
		}
		return NewErroF(TipagemErro, "Número incorreto de argumentos para %s %s. Esperava entre %d e %d, mas recebeu %d", tipo, nome, min, max, numArgs)
	}

	return nil
}