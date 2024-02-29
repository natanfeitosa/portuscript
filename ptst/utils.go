package ptst

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

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

// Resolve por exemplo:
//
//	`ResolveArquivoPtst("./atm.ptst", []string{"~/portuscript/exemplo"}, "")` -> ~/portuscript/exemplo/portuscript
func ResolveArquivoPtst(caminhoArqOuMod string, bases []string, curDir string) (string, error) {
	caminhoArqOuMod = strings.TrimSuffix(caminhoArqOuMod, "/")

	if len(curDir) > 0 {
		bases = append([]string{curDir}, bases...)
	}

	stat, err := os.Stat(caminhoArqOuMod)

	if path.IsAbs(caminhoArqOuMod) && err == nil && !stat.IsDir() {
		return caminhoArqOuMod, nil
	}

	for _, base := range bases {
		caminho, _ := filepath.Abs(path.Join(base, caminhoArqOuMod))

		stat, err = os.Stat(caminho)
		if err == nil && stat.IsDir() {
			caminho = path.Join(caminho, "inicio.ptst")
			_, err = os.Stat(caminho)
		}

		if filepath.Ext(caminho) == "" && os.IsNotExist(err) {
			caminho += ".so"
			_, err = os.Stat(caminho)

			if err != nil {
				caminho = strings.Replace(caminho, filepath.Ext(caminho), ".ptst", 1)
				_, err = os.Stat(caminho)
			}
		}

		if err != nil {
			if os.IsNotExist(err) {
				// Talvez dê sorte no próximo ciclo do loop
				continue
			}

			return "", NewErroF(SistemaErro, "Erro ao acessar '%s': %s", caminho, err)
		}

		// if !strings.HasSuffix(caminho, ".ptst") {
		// 	LancarErro(Errorf("o arquivo '%s' pode não ser um arquivo Portuscript válido. Você errou a extensão '.ptst'?", caminho))
		// }

		return caminho, nil
	}

	if err != nil && os.IsNotExist(err) {
		return "", NewErroF(ArquivoNaoEncontradoErro, "Não foi possível encontrar o arquivo '%s'", caminhoArqOuMod)
	}

	// FIXME: talvez isso não seja algo legal de fazer
	return "", nil
}
