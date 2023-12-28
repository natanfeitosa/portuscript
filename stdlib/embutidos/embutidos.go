package embutidos

import (
	"fmt"

	"github.com/natanfeitosa/portuscript/ptst"
)

func emb_imprima_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
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
		args,
	)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%s%s", resultado, final)
	return nil, nil
}

var emb_imprima_doc = "imprima(...objetos) -> imprime a representação ou a conversão em string dos objetos separados por espaço"

// FIXME: essa provavelmente não é a melhor implementação
func emb_leia_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("leia", false, args, 0, 1); err != nil {
		return nil, err
	}

	if len(args) == 1 {
		texto, err := ptst.NewTexto(args[0])

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
func emb_doc_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("doc", false, args, 1, 1); err != nil {
		return nil, err
	}

	arg := args[0]
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

func emb_int_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("int", false, args, 0, 1); err != nil {
		return nil, err
	}

	return ptst.NewInteiro(args[0])
}

var emb_int_doc = "int(objeto) -> Recebe um objeto e retorna uma representação numérica do tipo inteiro, se possível"

func emb_texto_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("texto", false, args, 0, 1); err != nil {
		return nil, err
	}

	return ptst.NewTexto(args[0])
}

var emb_texto_doc = "texto(objeto) -> Recebe um objeto e retorna uma representação no tipo texto, se possível"

func emb_tamanho_fn(mod ptst.Objeto, args ptst.Tupla) (ptst.Objeto, error) {
	if err := ptst.VerificaNumeroArgumentos("tamanho", false, args, 1, 1); err != nil {
		return nil, err
	}

	if obj, ok := args[0].(ptst.I__tamanho__); ok {
		return obj.O__tamanho__()
	}

	return nil, ptst.NewErroF(ptst.TipagemErro, "Objeto do tipo '%s' não implementa a interface '__tamanho__'.", args[0].Tipo().Nome)
}

var emb_tamanho_doc = ""

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
		ptst.NewMetodoOuPanic("int", emb_int_fn, emb_int_doc),
		ptst.NewMetodoOuPanic("texto", emb_texto_fn, emb_texto_doc),
		ptst.NewMetodoOuPanic("tamanho", emb_tamanho_fn, emb_tamanho_doc),
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