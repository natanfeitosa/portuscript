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

// Cria um soquete
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

func (s *Soquete) DefinirOpcoes(nivel, opcao, valor ptst.Inteiro) (ptst.Objeto, error) {
	if err := unix.SetsockoptInt(s.descritorDoSoquete, int(nivel), int(opcao), int(valor)); err != nil {
		panic(fmt.Sprintf("Erro ao definir opções do socket: %s", err))
	}

	return ptst.Nulo, nil
}

func (s *Soquete) Fechar() (ptst.Objeto, error) {
	if !s.fechado {
		s.fechado = ptst.Verdadeiro

		if err := unix.Close(s.descritorDoSoquete); err != nil {
			return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao fechar soquete: %s", err)
		}
	}

	return ptst.Nulo, nil
}

// Associa um soquete a um endereço
func (s *Soquete) AssociarSoquete(ip ptst.Texto, porta ptst.Inteiro) (ptst.Objeto, error) {
	addr := &unix.SockaddrInet4{Port: int(porta)}
	copy(addr.Addr[:], net.ParseIP(string(ip)).To16())

	if err := unix.Bind(s.descritorDoSoquete, addr); err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao associar soquete: %s", err)
	}

	return ptst.Nulo, nil
}

// Escuta por conexões em um soquete
func (s *Soquete) EscutarSoquete(backlog ptst.Inteiro) (ptst.Objeto, error) {
	if err := unix.Listen(s.descritorDoSoquete, int(backlog)); err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao escutar soquete: %s", err)
	}

	return ptst.Nulo, nil
}

// Aceita uma conexão
func (s *Soquete) AceitarConexao() (*Soquete, error) {
	for {
		// Poll para verificar eventos
		_, err := unix.Poll(s.pollFd, 1000) // Timeout de 1 segundo
		if err != nil {
			return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro no poll: %s", err)
		}

		if s.pollFd[0].Revents&unix.POLLIN != 0 {
			fd, _, err := unix.Accept(s.descritorDoSoquete)
			if err != nil {
				return nil, ptst.NewErroF(ptst.ErroDeSistema, "Erro ao aceitar conexão: %s", err)
			}

			// Adiciona o novo socket ao pollFd
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
		time.Sleep(10 * time.Millisecond) // Ajusta a espera entre as tentativas
	}
}

// Lê dados de uma conexão
func (s *Soquete) ReceberDados(tamanhoBuffer ptst.Inteiro) (*ptst.Bytes, error) {
	buffer := make([]byte, int(tamanhoBuffer))

	// loop := 0
	// Usar poll para verificar dados prontos para leitura
	for {
		n, err := unix.Poll(s.pollFd, 1) // Timeout de 0 para não bloquear
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
				return &ptst.Bytes{}, nil // Retorna vazio se não houver dados
			}
			val := &ptst.Bytes{Itens: buffer[:n]}
			return val, nil
		}
	}

	// Se não há dados prontos, retornar vazio
	return &ptst.Bytes{}, nil
}

// Envia dados para uma conexão
func (s *Soquete) EnviarDados(dados *ptst.Bytes) (ptst.Objeto, error) {
	_, err := unix.Write(s.descritorDoSoquete, dados.Itens)
	if err != nil {
		return nil, ptst.NewErroF(ptst.ErroDeSistema, "erro ao enviar dados: %v", err)
	}

	return ptst.Nulo, nil
}

// Conecta-se a um servidor
func (s *Soquete) Conectar(endereco ptst.Texto, porta ptst.Inteiro) (ptst.Objeto, error) {
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

	TipoSoquete.Mapa["associar"] = ptst.NewMetodoOuPanic("associar", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("associar", true, args, 2, 2); err != nil {
			return nil, err
		}

		return inst.(*Soquete).AssociarSoquete(args[0].(ptst.Texto), args[1].(ptst.Inteiro))
	}, "")

	TipoSoquete.Mapa["escutar"] = ptst.NewMetodoOuPanic("escutar", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("escutar", true, args, 0, 1); err != nil {
			return nil, err
		}

		backlog := ptst.Inteiro(0)
		if len(args) == 1 {
			backlog = args[0].(ptst.Inteiro)
		}

		return inst.(*Soquete).EscutarSoquete(backlog)
	}, "")

	TipoSoquete.Mapa["aceitar"] = ptst.NewMetodoOuPanic("aceitar", func(inst ptst.Objeto) (ptst.Objeto, error) {
		return inst.(*Soquete).AceitarConexao()
	}, "")

	TipoSoquete.Mapa["fechar"] = ptst.NewMetodoOuPanic("fechar", func(inst ptst.Objeto) (ptst.Objeto, error) {
		return inst.(*Soquete).Fechar()
	}, "")

	TipoSoquete.Mapa["receber"] = ptst.NewMetodoOuPanic("receber", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("receber", true, args, 0, 1); err != nil {
			return nil, err
		}

		tamanhoBuffer := ptst.Inteiro(0)
		if len(args) == 1 {
			tamanhoBuffer = args[0].(ptst.Inteiro)
		}

		return inst.(*Soquete).ReceberDados(tamanhoBuffer)
	}, "")

	TipoSoquete.Mapa["enviar"] = ptst.NewMetodoOuPanic("enviar", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("enviar", true, args, 1, 1); err != nil {
			return nil, err
		}

		return inst.(*Soquete).EnviarDados(args[0].(*ptst.Bytes))
	}, "")

	TipoSoquete.Mapa["conectar"] = ptst.NewMetodoOuPanic("conectar", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("conectar", true, args, 2, 2); err != nil {
			return nil, err
		}

		return inst.(*Soquete).Conectar(args[0].(ptst.Texto), args[1].(ptst.Inteiro))
	}, "")

	TipoSoquete.Mapa["def_nao_bloqueante"] = ptst.NewMetodoOuPanic("def_nao_bloqueante", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("def_nao_bloqueante", true, args, 1, 1); err != nil {
			return nil, err
		}

		return inst.(*Soquete).DefinirNaoBloqueante(args[0].(ptst.Booleano))
	}, "")

	TipoSoquete.Mapa["definir_opcoes"] = ptst.NewMetodoOuPanic("definir_opcoes", func(inst ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
		if err := ptst.VerificaNumeroArgumentos("definir_opcoes", true, args, 2, 3); err != nil {
			return nil, err
		}

		valor := ptst.Inteiro(1)

		if len(args) == 3 {
			valor = args[2].(ptst.Inteiro)
		}

		return inst.(*Soquete).DefinirOpcoes(args[0].(ptst.Inteiro), args[1].(ptst.Inteiro), valor)
	}, "")
}
