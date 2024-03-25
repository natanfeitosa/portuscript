package ptst_test

import (
	"reflect"
	"testing"

	"github.com/natanfeitosa/portuscript/ptst"
)

func TestStringParaBytes(t *testing.T) {
	frase := "tipo bytes"
	bytes, err := ptst.NewBytes(frase)
	if err != nil {
		t.Errorf("Não foi possível fazer a conversão, o seguinte erro foi retornado: %s", err)
	}

	if bytes == nil {
		t.Error("Era esperado uma instancia do tipo Bytes, mas foi retornado um 'nil'")
	}

	if !reflect.DeepEqual(bytes, &ptst.Bytes{[]byte(frase), false}) {
		t.Error("Aparentemente o construtor não está lidando direito com a conversão")
	}
}

func TestConversaoParaBytesPorMetodoImplementado(t *testing.T) {
	ptst.TipoTexto.Mapa["__bytes__"] = ptst.NewMetodoOuPanic(
		"__bytes__",
		func(inst ptst.Objeto) (ptst.Objeto, error) {
			return ptst.NewBytes(string(inst.(ptst.Texto)))
		},
		"",
	)

	texto := ptst.Texto("tipo bytes")
	bytes, err := ptst.NewBytes(texto)
	if err != nil {
		t.Errorf("Não foi possível fazer a conversão, o seguinte erro foi retornado: %s", err)
	}

	if bytes == nil {
		t.Error("Era esperado uma instancia do tipo Bytes, mas foi retornado um 'nil'")
	}

	if !reflect.DeepEqual(bytes, &ptst.Bytes{[]byte(texto), false}) {
		t.Error("Aparentemente o construtor não está lidando direito com a conversão")
	}
}

func TestCriacaoDeInstanciaVazia(t *testing.T) {
	if _, err := ptst.NewBytes(nil); err != nil {
		t.Error(err)
	}
}

func TestConversaoPelaChamadaDoConstrutor(t *testing.T) {
	if _, err := ptst.NovaInstancia(ptst.TipoBytes, ptst.Tupla{&ptst.Bytes{}}); err != nil {
		t.Error(err)
	}
}

func TestComparacaoRica(t *testing.T) {
	var a, b = &ptst.Bytes{}, &ptst.Bytes{}

	t.Run("`a == b`", func(t *testing.T) {
		res, err := ptst.Igual(a, b)
		if err != nil {
			t.Error(err)
		}

		if !res.(ptst.Booleano) {
			t.Error("deveria ser Verdadeiro, mas deu Falso")
		}
	})

	t.Run("`a != b`", func(t *testing.T) {
		res, err := ptst.Diferente(a, b)
		if err != nil {
			t.Error(err)
		}

		if res.(ptst.Booleano) {
			t.Error("deveria ser Falso, mas deu Verdadeiro")
		}
	})

	t.Run("`a >= b`", func(t *testing.T) {
		res, err := ptst.MaiorOuIgual(a, b)
		if err != nil {
			t.Error(err)
		}

		if !res.(ptst.Booleano) {
			t.Error("deveria ser Verdadeiro, mas deu Falso")
		}
	})

	t.Run("`a > b`", func(t *testing.T) {
		res, err := ptst.MaiorQue(a, b)
		if err != nil {
			t.Error(err)
		}

		if res.(ptst.Booleano) {
			t.Error("deveria ser Falso, mas deu Verdadeiro")
		}
	})

	t.Run("`a <= b`", func(t *testing.T) {
		res, err := ptst.MenorOuIgual(a, b)
		if err != nil {
			t.Error(err)
		}

		if !res.(ptst.Booleano) {
			t.Error("deveria ser Verdadeiro, mas deu Falso")
		}
	})

	t.Run("`a < b`", func(t *testing.T) {
		res, err := ptst.MenorQue(a, b)
		if err != nil {
			t.Error(err)
		}

		if res.(ptst.Booleano) {
			t.Error("deveria ser Falso, mas deu Verdadeiro")
		}
	})
}

func TestConversaoDeTipos(t *testing.T) {
	a := &ptst.Bytes{[]byte("a"), false}

	t.Run("=> booleano", func(t *testing.T) {
		res, err := ptst.NewBooleano(a)
		if err != nil {
			t.Error(err)
		}

		if !res.(ptst.Booleano) {
			t.Error("deveria ser Verdadeiro, mas deu Falso")
		}
	})

	t.Run("=> texto", func(t *testing.T) {
		text, err := ptst.NewTexto(a)
		if err != nil {
			t.Error(err)
		}

		if res, _ := ptst.Igual(ptst.Texto("a"), text); !res.(ptst.Booleano) {
			t.Error("a comparação com o tipo convertido não deu verdadeiro")
		}
	})
}
