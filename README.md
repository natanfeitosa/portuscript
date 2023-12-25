# PortuScript

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Open Source Helpers](https://www.codetriage.com/natanfeitosa/portuscript/badges/users.svg)](https://www.codetriage.com/natanfeitosa/portuscript)
[![Documentation Status](https://readthedocs.org/projects/portudoc/badge/?version=latest)](https://portudoc.readthedocs.io/pt/latest/?badge=latest)

## Sobre

**PortuScript** é uma linguagem de programação brasileira, desenvolvida por brasileiros, totalmente em português. Mais do que uma simples linguagem para treino de lógica, o PortuScript visa proporcionar uma experiência de programação acessível e envolvente para a comunidade de língua portuguesa.

### Características Principais

- **Brasileira por Natureza**: Desenvolvida com o objetivo de ser inclusiva e acessível para falantes de português.
- **Simples e Poderosa**: Projetada para facilitar o aprendizado de programação, mantendo a capacidade de lidar com tarefas complexas.
- **Comunidade Ativa**: Contribua e faça parte de uma comunidade que apoia o crescimento e desenvolvimento do PortuScript.

## Como Começar

### Baixar Binários (Recomendado)

Os binários pré-compilados estão disponíveis para os seguintes sistemas operacionais e arquiteturas:

- Darwin (macOS)
- Linux
- Windows

Navege até a [página de lançamentos](https://github.com/natanfeitosa/portuscript/releases) e escolha o binário correspondente ao seu sistema operacional e arquitetura, faça o download e comece a usar o PortuScript imediatamente.

### Compilar a partir do Código Fonte

Se preferir compilar a partir do código fonte, certifique-se de ter o [Go](https://golang.org/doc/install) instalado. Em seguida, execute os seguintes comandos:

```bash
git clone https://github.com/natanfeitosa/portuscript.git
cd portuscript
go build
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

Sinta-se à vontade para contribuir para o desenvolvimento do PortuScript.

Sua contribuição é super bem vinda. Caso não tenha ideias de melhorias nem perceba um bug, você ainda pode ajudar dando uma olhadinha nas nossas [metas](/metas.md) e vendo o que você pode fazer, ou talvez tendo alguma ideia a partir daí.

Para mais informações, consulte o [Guia de Contribuição](/CONTRIBUTING.md).

> Este projeto está em constante evolução. Se você encontrar problemas ou tiver sugestões, por favor, abra uma [issue](https://github.com/natanfeitosa/portuscript/issues).

Agradecemos pela sua contribuição!
