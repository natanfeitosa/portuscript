package cmd

// FIXME: aqui quase tudo pode ser refatorado e ser feito reaproveitamento de código

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
)

func isWindows() bool {
	return runtime.GOOS == "windows"
}

// `a` é a versão instalada
func jaAtualizado(a, b string) bool {
	i, _ := semver.NewConstraint("< " + b)
	n, _ := semver.NewVersion(a)

	return i.Check(n)
}

func urlDaVersao() string {
	url := "https://github.com/natanfeitosa/portuscript/releases/latest/download/"
	
	if OS := strings.ToTitle(runtime.GOOS); OS == "Darwind" || OS == "Linux" || OS == "Windows" {
		url += OS + "_"
	} else {
		url += "Linux_"
	}

	switch arch := runtime.GOARCH; arch {
	case "amd64":
		url += "x86_64"
	case "386":
		url += "i" + arch
	default:
		url += arch
	}

	if isWindows() {
		return url + ".zip"
	}

	return url + ".tar.gz"
}

type Tag struct {
	Name string `json:"name"`
}

func atualize() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("erro ao tentar montar o caminho da versão atual: %s", err)
	}

	raizPortuscript := path.Join(home, ".portuscript/bin/")
	binario := path.Join(raizPortuscript, "portuscript")

	if isWindows() {
		binario = binario + ".exe"
	}

	comandoEx, err := exec.Command(binario, "-v").Output()
	if err != nil {
		return fmt.Errorf("erro ao obter a versão instalada, provavelmente você ainda não instalou nenhuma versão, veja: <https://github.com/natanfeitosa/portuscript/?tab=readme-ov-file#com-bash>")
	}
	comandoExString := strings.Split(strings.Trim(string(comandoEx), " \t\n"), " ")
	versaoInstalada := comandoExString[len(comandoExString)-1]

	if versaoInstalada == "dev" {
		return fmt.Errorf("você tem a versão 'dev' instalada, este comando ainda não é capaz de instalar nesse cenário")
	}

	response, err := http.Get("https://api.github.com/repos/natanfeitosa/portuscript/tags")
	if err != nil {
		return fmt.Errorf("erro ao obter as versões no repositório")
	}
	defer response.Body.Close()

	var tags []Tag
	if err := json.NewDecoder(response.Body).Decode(&tags); err != nil {
		return fmt.Errorf("erro ao decodificar a resposta JSON")
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("erro na resposta do servidor: %s", response.Status)
	}

	if len(tags) == 0 {
		return fmt.Errorf("nenhuma versão encontrada no repositório")
	}

	ultimaTag := strings.TrimPrefix(tags[0].Name, "v")

	if !jaAtualizado(versaoInstalada, ultimaTag) {
		fmt.Printf("Você já tem a versão mais recente (%s) instalada.\n", versaoInstalada)
	} else {
		fmt.Printf("Nova versão disponível: %s\n", ultimaTag)

		f, err := os.CreateTemp("", "-ptst")
		if err != nil {
			return fmt.Errorf("erro ao criar um diretorio temporário")
		}
		defer os.Remove(f.Name())

		fmt.Println("Baixando arquivos necessários")

		compactTemp := f.Name()
		curl := exec.Command(
			"curl", "--fail", "--location", "--progress-bar", "--output", compactTemp, urlDaVersao(),
		)
		curl.Stdout = os.Stdout
		curl.Stderr = os.Stderr

		if curl.Run() != nil {
			return fmt.Errorf("falha ao baixar os arquivos")
		}

		fmt.Println("Instalando a nova versão...")

		if isWindows() {
			if _, err := exec.LookPath("unzip"); err == nil {
				cmd := exec.Command("unzip", "-d", raizPortuscript, "-o", compactTemp)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					return fmt.Errorf("erro ao descompactar com unzip:\n%s", err)
				}
			} else {
				cmd := exec.Command("7z", "x", "-o", raizPortuscript, "-y", compactTemp)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					return fmt.Errorf("erro ao descompactar com 7z:\n%s", err)
				}
			}
		} else {
			cmd := exec.Command("tar", "-xf", compactTemp, "-C", raizPortuscript)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return fmt.Errorf("erro ao descompactar com tar:\n%s", err)
			}
		}

		fmt.Println("Nova versão instalada com sucesso!")
	}

	return nil
}

func comandoAtualize() *cobra.Command {
	return &cobra.Command{
		Use:   "atualize",
		Short: "Atualiza a versão do portuscript instalada se já não for a mais recente",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := atualize(); err != nil {
				fmt.Fprintf(os.Stderr, "\033[1m\033[31m%s\033[0m\n", err)
				os.Exit(1)
			}

			os.Exit(0)
		},
	}
}
