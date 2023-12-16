# PortuScript

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Sobre

**PortuScript** é uma linguagem de programação brasileira, desenvolvida por brasileiros, totalmente em português. Mais do que uma simples linguagem para treino de lógica, o PortuScript visa proporcionar uma experiência de programação acessível e envolvente para a comunidade de língua portuguesa.

### Características Principais

- **Brasileira por Natureza**: Desenvolvida com o objetivo de ser inclusiva e acessível para falantes de português.
- **Simples e Poderosa**: Projetada para facilitar o aprendizado de programação, mantendo a capacidade de lidar com tarefas complexas.
- **Comunidade Ativa**: Contribua e faça parte de uma comunidade que apoia o crescimento e desenvolvimento do PortuScript.

## Como Começar

### Baixar Binários (Recomendado)

Os binários pré-compilados estão disponíveis para os seguintes sistemas operacionais e arquiteturas:

#### Arquivos binários direto
| Arquitetura | Darwin (macOS) | Linux | Windows |
|-------------|----------------|-------|---------|
| arm64    | [portuscript](dist/portuscript_darwin_arm64/portuscript) | [portuscript](dist/portuscript_linux_arm64/portuscript) | [portuscript.exe](dist/portuscript_windows_arm64/portuscript.exe) |
| i386     | N/A                                                      | [portuscript](dist/portuscript_linux_386/portuscript)   | [portuscript.exe](dist/portuscript_windows_386/portuscript.exe) |
| amd64_v1 | [portuscript](dist/portuscript_darwin_amd64_v1/portuscript) | [portuscript](dist/portuscript_linux_amd64_v1/portuscript)    | [portuscript.exe](dist/portuscript_windows_amd64_v1/portuscript.exe) |


#### Arquivo binário e README compactados

| Arquitetura | Darwin (macOS)                        | Linux                                  | Windows                                |
|-------------|--------------------------------------|----------------------------------------|----------------------------------------|
| arm64       | [Darwin_arm64.tar.gz](dist/Darwin_arm64.tar.gz) | [Linux_arm64.tar.gz](dist/Linux_arm64.tar.gz) | [Windows_arm64.zip](dist/Windows_arm64.zip) |
| i386        | N/A                                  | [Linux_i386.tar.gz](dist/Linux_i386.tar.gz) | [Windows_i386.zip](dist/Windows_i386.zip) |
| amd64_v1    | [Darwin_x86_64.tar.gz](dist/Darwin_x86_64.tar.gz) | [Linux_x86_64.tar.gz](dist/Linux_x86_64.tar.gz) | [Windows_x86_64.zip](dist/Windows_x86_64.zip) |

Escolha o binário correspondente ao seu sistema operacional e arquitetura, faça o download e comece a usar o PortuScript imediatamente.

### Compilar a partir do Código Fonte

Se preferir compilar a partir do código fonte, certifique-se de ter o [Go](https://golang.org/doc/install) instalado. Em seguida, execute os seguintes comandos:

```bash
git clone https://github.com/natanfeitosa/portuscript.git
cd portuscript
# Instruções de compilação aqui
```

## CLI - Utilização Básica

A CLI do PortuScript oferece as seguintes funcionalidades:

- **Abrir o Playground**: Se nenhum argumento for passado, a CLI abrirá o Playground interativo.
```bash
portuscript
```

- **Executar Arquivo `*.ptst`**: Se o caminho de um arquivo `.ptst` for fornecido como argumento, o PortuScript executará o script contido no arquivo.

```bash
portuscript caminho/do/arquivo.ptst
```

- **Executar Código Inline**: Se a flag `-c` ou `--codigo` for usada, é possível executar código inline diretamente na linha de comando.

```bash
portuscript -c "seu código aqui"
```

## Exemplos de Uso

Explore o diretório [exemplos](/exemplos/) para encontrar uma variedade de scripts de exemplo que demonstram a versatilidade do PortuScript.

## Gramática

O diretório gramatica contém informações detalhadas sobre a gramática da linguagem. Consulte para uma compreensão mais profunda da estrutura da linguagem.

## Contribuindo
Sinta-se à vontade para contribuir para o desenvolvimento do PortuScript. Para mais informações, consulte o Guia de Contribuição.

> Este projeto está em constante evolução. Se você encontrar problemas ou tiver sugestões, por favor, abra uma [issue](https://github.com/natanfeitosa/portuscript/issues).

Agradecemos pela sua contribuição!
