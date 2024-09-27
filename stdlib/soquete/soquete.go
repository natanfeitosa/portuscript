package soquete

import (
	"syscall"

	"github.com/natanfeitosa/portuscript/ptst"
)

var familia = ptst.Mapa{
	"AF_INET":     ptst.Inteiro(syscall.AF_INET),
	"AF_INET6":    ptst.Inteiro(syscall.AF_INET6),
	"SOCK_STREAM": ptst.Inteiro(syscall.SOCK_STREAM),
	"SOCK_DGRAM":  ptst.Inteiro(syscall.SOCK_DGRAM),
}

func init() {
	constantes := ptst.Mapa{
		TipoSoquete.Nome: TipoSoquete,
	}
	constantes.Atualizar(familia, false)

	metodos := []*ptst.Metodo{}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome:    "soquete",
				Arquivo: "stdlib/soquete",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}
