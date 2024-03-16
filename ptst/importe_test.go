package ptst_test

import (
	"testing"

	"github.com/natanfeitosa/portuscript/ptst"
	_ "github.com/natanfeitosa/portuscript/stdlib"
)

func assertPanic(t *testing.T, f func()) {
    t.Helper()
    defer func() { _ = recover() }()
    f()
    t.Errorf("Era esperado que houvesse um `panic`")
}

func TestMaquinarioImporteModulo(t *testing.T) {
	var mod, obj ptst.Objeto
	var err error

	ctx := ptst.NewContexto(ptst.OpcsContexto{})
	defer ctx.Terminar()

	if mod, err = ptst.MaquinarioImporteModulo(ctx, "colorize", nil); err != nil {
		t.Error(err)
	}

	if obj, err = ptst.ObtemAtributoS(mod, "converteRGB"); err != nil {
		t.Error(err)
	}

	if obj.(*ptst.Metodo).Nome != "converteRGB" {
		t.Error("erro no nome do método")
	}
}

func TestMultiImporteModulo(t *testing.T) {
	ctx := ptst.NewContexto(ptst.OpcsContexto{})
	defer ctx.Terminar()

	if  err := ptst.MultiImporteModulo(ctx, "colorize", "embutidos"); err != nil {
		t.Error(err)
	}

	if _, err := ctx.Modulos.ObterModulo("embutidos"); err != nil {
		t.Error(err)
	}
}

func TestImporteSemCriarContexto(t *testing.T) {
	teste := func ()  {
		ptst.Importe("embutidos", nil)
	}

	assertPanic(t, teste)
}

func TestImporteComContexto(t *testing.T) {
	var mod, obj ptst.Objeto
	var err error

	ctx := ptst.NewContexto(ptst.OpcsContexto{})
	defer ctx.Terminar()

	if mod, err = ptst.Importe("colorize", nil); err != nil {
		t.Error(err)
	}

	if obj, err = ptst.ObtemAtributoS(mod, "converteRGB"); err != nil {
		t.Error(err)
	}

	if obj.(*ptst.Metodo).Nome != "converteRGB" {
		t.Error("erro no nome do método")
	}
}

func TestImportacoesRelativas(t *testing.T) {
	ctx := ptst.NewContexto(ptst.OpcsContexto{})
	defer ctx.Terminar()

	if _, err := ptst.Importe("./algo", nil); err == nil {
		t.Error("Era esperado um erro, pois `algo.ptst` não existe")
	}
}