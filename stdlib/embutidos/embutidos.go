package embutidos

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/ptst"
)

func emb_imprima_fn(mod ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
	const (
		final     = ptst.Texto("\n")
		separador = ptst.Texto(" ")
	)

	junta, err := ptst.ObtemItemS(separador, "junta")

	if err != nil {
		return nil, err
	}

	// resultado, err := ptst.Chamar(junta, args)
	resultado, err := ptst.Chamar(
		junta,
		args.(ptst.Tupla),
	)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%s%s", resultado, final)
	return nil, nil
}

var emb_imprima_doc = "imprima(...objetos) -> imprime a representação ou a conversão em string dos objetos separados por espaço"

// FIXME: essa provavelmente não é a melhor implementação
func emb_leia_fn(mod ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
	targs := args.(ptst.Tupla)

	if len(targs) > 1 {
		return nil, ptst.NewErroF(ptst.TipagemErro, "A funçao leia() esperava receber no máximo 1 argumento, mas recebeu um número de %v", len(args.(ptst.Tupla)))
	}

	if len(targs) == 1 {
		texto, err := ptst.NewTexto(targs[0])

		if err != nil {
			return nil, err
		}

		fmt.Printf("%s", texto)
	}

	var leitura string
	fmt.Scan(&leitura)
	return ptst.Texto(leitura), nil
}

var emb_leia_doc = "leia(frase_para_imprimir) -> imprime um texto se especificado e lê uma entrada do usuário, retornando-a"

// FIXME: doc(imprima) e doc(doc) não funcionam
func emb_doc_fn(mod ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
	if len(args.(ptst.Tupla)) != 1 {
		return nil, ptst.NewErroF(ptst.TipagemErro, "A funçao doc() esperava receber 1 argumento, mas recebeu um número de %v", len(args.(ptst.Tupla)))
	}

	arg := args.(ptst.Tupla)[0]
	imp, err := mod.(*ptst.Modulo).Contexto.ObterSimbolo("imprima")

	if err != nil {
		return nil, err
	}

	if obj, ok := arg.(ptst.I_ObtemDoc); ok {
		return ptst.Chamar(imp.Valor, ptst.Tupla{ptst.Texto(obj.ObtemDoc())})
	}

	return ptst.Chamar(imp.Valor, ptst.Tupla{ptst.Texto(arg.Tipo().ObtemDoc())})
}

var emb_doc_doc = "doc(objeto) -> Obtem a documentação do objeto"

// func emb_sair_fn(mod ptst.Objeto, args ptst.Objeto) (ptst.Objeto, error) {
// 	targs := args.(ptst.Tupla)

// 	if len(targs) > 1 {
// 		return nil, ptst.NewErroF(ptst.TipagemErro, "A funçao sair() esperava receber no máximo 1 argumento, mas recebeu um número de %v", len(targs))
// 	}

// 	if !ptst.MesmoTipo(targs[0], ptst.Inteiro(0)) {
// 		return nil, ptst.NewErroF(ptst.TipagemErro, "O tipo de argumento esperado na funçao sair() era 'Inteiro', mas foi recebido o tipo '%s'", targs[0].Tipo().Nome)
// 	}

// 	os.Exit(int(targs[0].(ptst.Inteiro)))
// 	return nil, nil
// }

// var emb_sair_doc = ""

func init() {
	constantes := ptst.Mapa{
		"Verdadeiro": ptst.Verdadeiro,
		"Falso":      ptst.Falso,
		"Nulo":       ptst.Nulo,
	}

	metodos := []*ptst.Metodo{
		ptst.NewMetodoOuPanic("imprima", emb_imprima_fn, emb_imprima_doc),
		ptst.NewMetodoOuPanic("leia", emb_leia_fn, emb_leia_doc),
		ptst.NewMetodoOuPanic("doc", emb_doc_fn, emb_doc_doc),
		// ptst.NewMetodoOuPanic("sair", emb_sair_fn, emb_sair_doc),
	}

	ptst.RegistraModuloImpl(
		&ptst.ModuloImpl{
			Info: ptst.ModuloInfo{
				Nome: "embutidos",
			},
			Constantes: constantes,
			Metodos:    metodos,
		},
	)
}
