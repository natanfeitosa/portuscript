package main

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/ptst"
)

// Função obrigatória para quando o interpretador tentar obter o módulo
func InicializaModulo() *ptst.ModuloImpl {
	return &ptst.ModuloImpl{
		Info: ptst.ModuloInfo{
			Nome: "externos",
			Doc: "Um módulo externo para teste",
		},
		Metodos: []*ptst.Metodo{
			ptst.NewMetodoOuPanic("exiba", func(_ ptst.Objeto, args ptst.Tupla) (obj ptst.Objeto, err error) {
				junta, err := ptst.ObtemAtributoS(ptst.Texto(", "), "junta")
				if err != nil {
					return
				}

				juntos, err := ptst.Chamar(junta, args)
				if err != nil {
					return
				}

				fmt.Printf("externos: %s", juntos.(ptst.Texto))
				return
			}, "Exibe algo no terminal, ok?"),
		},
	}
}
