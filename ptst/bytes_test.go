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
	if _, err := ptst.NovaInstancia(ptst.TipoBytes, ptst.Tupla{ptst.Bytes{}}); err != nil {
		t.Error(err)
	}
	
}