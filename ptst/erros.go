package ptst

import "fmt"

type Erro struct {
	Base     *Tipo
	Contexto *Contexto
	Mensagem Objeto // Está certo isso?
}

var BaseErro = TipoObjeto.NewTipo(
	"BaseErro",
	"A classe de erro base para todas as outras.",
)

var (
	TipoErro         = BaseErro.NewTipo("Erro", "Base comum para todos os erros que não são de saída.")
	SintaxeErro      = TipoErro.NewTipo("SintaxeErro", "Sintaxe Invalida.")
	ReatribuicaoErro = TipoErro.NewTipo("ReatribuicaoErro", "Proibido redeclarar.")
	AtributoErro     = TipoErro.NewTipo("AtributoErro", "Atributo não encontrado.")
	TipagemErro      = TipoErro.NewTipo("TipagemErro", "Tipo de argumento inapropriado.")
	NomeErro         = TipoErro.NewTipo("NomeErro", "Erro de nome que não pode ser achado.")
	ImportacaoErro   = TipoErro.NewTipo("ImportacaoErro", "Não é possível encontrar o módulo ou símbolo nele")
)

func NewErro(tipo *Tipo, args Objeto) *Erro {
	return &Erro{Base: tipo, Mensagem: args}
}

func NewErroF(tipo *Tipo, format string, p ...any) *Erro {
	return &Erro{Base: tipo, Mensagem: Texto(fmt.Sprintf(format, p...))}
}

func (e *Erro) AdicionarContexto(contexto *Contexto) {
	if e.Contexto != nil {
		return
	}

	e.Contexto = contexto
}

func (e *Erro) Error() string {
	format := "Arquivo %v:\n  %v: %v"

	return fmt.Sprintf(format, e.Contexto.Caminho, e.Base.Nome, e.Mensagem.(Texto))
}
