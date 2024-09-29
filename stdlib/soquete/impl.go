package soquete

import (
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/natanfeitosa/portuscript/ptst"
	"golang.org/x/sys/unix"
)

func ultimoElemento[T any](slice []T) T {
	if len(slice) == 0 {
		var none T
		return none
	}
	return slice[len(slice)-1]
}

type Soquete struct {
	descritorDoSoquete       int
	familia, tipo, protocolo ptst.Inteiro
	fechado                  ptst.Booleano
	pollFd                   []unix.PollFd
	p                        *Soquete
}

var TipoSoquete = ptst.TipoObjeto.NewTipo(
	"Soquete",
	`Soquete(familia, tipo) -> Soquete
Cria um novo soquete usando a família de endereços, o tipo de soquete e o número de protocolo fornecidos.`,
)

var _ ptst.Objeto = (*Soquete)(nil)

func (s *Soquete) Tipo() *ptst.Tipo {
	return TipoSoquete
}

func NewSoquete(familia, tipo, protocolo ptst.Inteiro) (ptst.Objeto, error) {
	fd, err := unix.Socket(int(familia), int(tipo), int(protocolo))
	if err != nil {
		if err == unix.EAFNOSUPPORT {
			return nil, ptst.NewErroF(ptst.ValorErro, "Família de endereço não suportada")
		}

		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro externo: %s", err)
	}

	s := &Soquete{descritorDoSoquete: fd, familia: familia, tipo: tipo, protocolo: protocolo, fechado: ptst.Falso}
	s.pollFd = []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}} // Inicializa o pollFd

	return s, nil
}

func (s *Soquete) DefinirNaoBloqueante(naobloqueante ptst.Booleano) (ptst.Objeto, error) {
	if err := unix.SetNonblock(s.descritorDoSoquete, bool(naobloqueante)); err != nil {
		panic(err)
	}

	return ptst.Nulo, nil
}

func (s *Soquete) DefineOpcoes(nivel, opcao, valor ptst.Inteiro) (ptst.Objeto, error) {
	if err := unix.SetsockoptInt(s.descritorDoSoquete, int(nivel), int(opcao), int(valor)); err != nil {
		panic(fmt.Sprintf("Erro ao definir opções do socket: %s", err))
	}

	return ptst.Nulo, nil
}

func (s *Soquete) Fecha() (ptst.Objeto, error) {
	if !s.fechado {
		s.fechado = ptst.Verdadeiro

		if err := unix.Close(s.descritorDoSoquete); err != nil {
			return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao fechar soquete: %s", err)
		}
	}

	return ptst.Nulo, nil
}

func (s *Soquete) AssociaSoquete(ip ptst.Texto, porta ptst.Inteiro) (ptst.Objeto, error) {
	addr := &unix.SockaddrInet4{Port: int(porta)}
	copy(addr.Addr[:], net.ParseIP(string(ip)).To16())

	if err := unix.Bind(s.descritorDoSoquete, addr); err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao associar soquete: %s", err)
	}

	return ptst.Nulo, nil
}

func (s *Soquete) OuveSoquete(backlog ptst.Inteiro) (ptst.Objeto, error) {
	if err := unix.Listen(s.descritorDoSoquete, int(backlog)); err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao ouvir soquete: %s", err)
	}

	return ptst.Nulo, nil
}

func (s *Soquete) AceitaConexao() (*Soquete, error) {
	for {
		_, err := unix.Poll(s.pollFd, 1000)
		if err != nil {
			return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro no poll: %s", err)
		}

		if s.pollFd[0].Revents&unix.POLLIN != 0 {
			fd, _, err := unix.Accept(s.descritorDoSoquete)
			if err != nil {
				return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao aceitar conexão: %s", err)
			}

			s.pollFd = append(s.pollFd, unix.PollFd{Fd: int32(fd), Events: unix.POLLIN})

			soq := &Soquete{
				descritorDoSoquete: fd,
				familia:            s.familia,
				tipo:               s.tipo,
				protocolo:          s.protocolo,
				fechado:            ptst.Falso,
				pollFd:             []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}},
				p:                  s,
			}
			return soq, nil
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (s *Soquete) RecebeDados(tamanhoBuffer ptst.Inteiro) (*ptst.Bytes, error) {
	buffer := make([]byte, int(tamanhoBuffer))

	// loop := 0
	for {
		n, err := unix.Poll(s.pollFd, 1)
		if err != nil {
			if err == unix.EINTR {
				// if loop > 0 {
				// 	break
				// }
				// loop += 1
				// // Se o poll for interrompido, continue o loop
				// continue
				break
			}
			return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro no poll: %s", err)
		}

		if n > 0 && ultimoElemento(s.pollFd).Revents&unix.POLLIN != 0 {
			n, _, err := unix.Recvfrom(s.descritorDoSoquete, buffer, 0)
			if err != nil {
				return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao receber dados: %v", err)
			}

			if n <= 0 {
				return &ptst.Bytes{}, nil
			}
			val := &ptst.Bytes{Itens: buffer[:n]}
			return val, nil
		}
	}

	// Se não há dados prontos, retornar vazio
	return &ptst.Bytes{}, nil
}

func (s *Soquete) EnviaDados(dados *ptst.Bytes) (ptst.Objeto, error) {
	_, err := unix.Write(s.descritorDoSoquete, dados.Itens)
	if err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "erro ao enviar dados: %v", err)
	}

	return ptst.Nulo, nil
}

func (s *Soquete) Conecta(endereco ptst.Texto, porta ptst.Inteiro) (ptst.Objeto, error) {
	addr, err := s.resolveEndereco(string(endereco), int(porta))
	if err != nil {
		return nil, err
	}

	if err := unix.Connect(s.descritorDoSoquete, addr); err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "erro ao conectar ao servidor: %v", err)
	}

	return ptst.Nulo, nil
}

func (s *Soquete) resolveEndereco(endereco string, porta int) (unix.Sockaddr, error) {
	ips, err := net.LookupIP(string(endereco))
	if err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao resolver o endereço: %v", err)
	}

	if len(ips) == 0 {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Nenhum endereço IP encontrado para: %s", endereco)
	}

	for _, ip := range ips {
		switch s.familia {
		case syscall.AF_INET6:
			if ip6 := ip.To16(); ip6 != nil && ip.To4() == nil {
				addr := &unix.SockaddrInet6{Port: porta}
				copy(addr.Addr[:], ip6)
				return addr, nil
			}
		case syscall.AF_INET:
			if ip4 := ip.To4(); ip4 != nil {
				addr := &unix.SockaddrInet4{Port: porta}
				copy(addr.Addr[:], ip)
				return addr, nil
			}
		}
	}

	return nil, nil
}

func init() {
	TipoSoquete.Nova = func(args ptst.Tupla) (ptst.Objeto, error) {
		if argsLen := len(args); argsLen != 3 {
			if argsLen < 2 {
				return nil, ptst.NewErroF(ptst.TipagemErro, "Soquete() esperava receber no mínimo 2 argumentos, mas recebeu %d", argsLen)
			}

			if argsLen > 3 {
				return nil, ptst.NewErroF(ptst.TipagemErro, "Soquete() esperava receber no máximo 3 argumentos, mas recebeu %d", argsLen)
			}
		}

		var familia, tipo, protocolo ptst.Inteiro = args[0].(ptst.Inteiro), args[1].(ptst.Inteiro), ptst.Inteiro(0)

		if len(args) == 3 {
			protocolo = args[2].(ptst.Inteiro)
		}

		return NewSoquete(familia, tipo, protocolo)
	}

	TipoSoquete.Mapa["associa"] = ptst.NewMetodoOuPanic("associa", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("associa", true, args, 2, 2); err != nil {
			return nil, err
		}

		return inst.(*Soquete).AssociaSoquete(args[0].(ptst.Texto), args[1].(ptst.Inteiro))
	}, "soquete.associa(ip, porta) -> Nulo\n\nAssocia um soquete a um endereço IP e porta.")

	TipoSoquete.Mapa["ouve"] = ptst.NewMetodoOuPanic("ouve", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("ouve", true, args, 0, 1); err != nil {
			return nil, err
		}

		backlog := ptst.Inteiro(0)
		if len(args) == 1 {
			backlog = args[0].(ptst.Inteiro)
		}

		return inst.(*Soquete).OuveSoquete(backlog)
	}, "soquete.ouve(backlog?) -> Nulo\n\nInicia a escuta por conexões em um soquete.\nSe não for passado o backlog, que é o número máximo de conexões pendentes na fila, por padrão será 1.")

	TipoSoquete.Mapa["aceita"] = ptst.NewMetodoOuPanic("aceita", func(inst ptst.Objeto) (ptst.Objeto, error) {
		return inst.(*Soquete).AceitaConexao()
	}, "soquete.aceita() -> Soquete\n\nAceita uma nova conexão em um soquete que está escutando e retorna o soquete referente ao cliente.")

	TipoSoquete.Mapa["fecha"] = ptst.NewMetodoOuPanic("fecha", func(inst ptst.Objeto) (ptst.Objeto, error) {
		return inst.(*Soquete).Fecha()
	}, "soquete.fecha() -> Nulo\n\nFecha o soquete.")

	TipoSoquete.Mapa["recebe"] = ptst.NewMetodoOuPanic("recebe", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("recebe", true, args, 0, 1); err != nil {
			return nil, err
		}

		tamanhoBuffer := ptst.Inteiro(0)
		if len(args) == 1 {
			tamanhoBuffer = args[0].(ptst.Inteiro)
		}

		return inst.(*Soquete).RecebeDados(tamanhoBuffer)
	}, "soquete.recebe(tamanhoBuffer?) -> Bytes\n\nRecebe os dados de uma conexão e retorna no tipo Bytes\nSe não for definido um tamanho de buffer, o padrão será 0")

	TipoSoquete.Mapa["envia"] = ptst.NewMetodoOuPanic("envia", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("envia", true, args, 1, 1); err != nil {
			return nil, err
		}

		return inst.(*Soquete).EnviaDados(args[0].(*ptst.Bytes))
	}, "soquete.envia(dados) -> Nulo\n\nEnvia um objeto do tipo Bytes para o outro lado da conexão")

	TipoSoquete.Mapa["conecta"] = ptst.NewMetodoOuPanic("conecta", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("conecta", true, args, 2, 2); err != nil {
			return nil, err
		}

		return inst.(*Soquete).Conecta(args[0].(ptst.Texto), args[1].(ptst.Inteiro))
	}, "soquete.conecta(endereco, porta) -> Nulo\n\nSe conecta a um servidor pela porta e endereço informado.\nO endereço pode ser um IP ou nome de domínio como: exemplo.com")

	TipoSoquete.Mapa["def_nao_bloqueante"] = ptst.NewMetodoOuPanic("def_nao_bloqueante", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("def_nao_bloqueante", true, args, 1, 1); err != nil {
			return nil, err
		}

		return inst.(*Soquete).DefinirNaoBloqueante(args[0].(ptst.Booleano))
	}, "soquete.def_nao_bloqueante(naoBloqueante) -> Nulo\n\nDefine se o soquete deve operar em modo não bloqueante")

	TipoSoquete.Mapa["define_opcoes"] = ptst.NewMetodoOuPanic("define_opcoes", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("define_opcoes", true, args, 2, 3); err != nil {
			return nil, err
		}

		valor := ptst.Inteiro(1)

		if len(args) == 3 {
			valor = args[2].(ptst.Inteiro)
		}

		return inst.(*Soquete).DefineOpcoes(args[0].(ptst.Inteiro), args[1].(ptst.Inteiro), valor)
	}, "soquete.define_opcoes(nivel, opcao, valor) -> Nulo\n\nDefine opções para o soquete.")
}
