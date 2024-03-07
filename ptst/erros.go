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

func init() {
	BaseErro.Mapa["__nova_instancia__"] = NewMetodoOuPanic(
		"__nova_instancia__",
		func(inst Objeto, args Tupla) (Objeto, error) {
			tipo := inst.(*Tipo)
			message, ok := args[0].(Texto)
			if !ok {
				return nil, NewErroF(TipagemErro, "O primeiro argumento de '%s' deve ser do tipo '%s', e não '%s'", tipo.Nome, TipoTexto.Nome, args[0].Tipo().Nome)
			}

			return NewErro(tipo, message), nil
		},
		"",
	)
}

var (
	TipoErro = BaseErro.NewTipo("Erro", "Base comum para todos os erros que não são de saída.")

	SintaxeErro       = TipoErro.NewTipo("SintaxeErro", "Sintaxe Invalida.")
	ReatribuicaoErro  = TipoErro.NewTipo("ReatribuicaoErro", "Proibido redeclarar.")
	AtributoErro      = TipoErro.NewTipo("AtributoErro", "Atributo não encontrado.")
	TipagemErro       = TipoErro.NewTipo("TipagemErro", "Tipo de argumento inapropriado.")
	NomeErro          = TipoErro.NewTipo("NomeErro", "Erro de nome que não pode ser achado.")
	ImportacaoErro    = TipoErro.NewTipo("ImportacaoErro", "Não é possível encontrar o módulo ou símbolo nele")
	ValorErro         = TipoErro.NewTipo("ValorErro", "O valor é inapropriádo ou sua ocorrencia não existe")
	IndiceErro        = TipoErro.NewTipo("IndiceErro", "O indice está fora do range aceito")
	RuntimeErro       = TipoErro.NewTipo("RuntimeErro", "Erro no ambiente de execução")
	FimIteracao       = TipoErro.NewTipo("FimIteracao", "Sinaliza o fim da iteração quando `objeto.__proximo__() não retorna mais nada")
	ErroDeAsseguracao = TipoErro.NewTipo("ErroDeAsseguracao", "Erro lançado em um `assegura obj`")

	ConsultaErro = TipoErro.NewTipo("ConsultaErro", "Classe base para erros que envolem chave ou indice em elementos")
	ChaveErro    = ConsultaErro.NewTipo("ChaveErro", "Lançado quando a chave de um mapa não existe ou é inválida")

	SistemaErro              = TipoErro.NewTipo("SistemaErro", "Erro relacionado ao sistema operacional")
	ArquivoNaoEncontradoErro = SistemaErro.NewTipo("ArquivoNaoEncontradoErro", "O arquivo não pôde ser encontrado")

	// Apenas para fins de controle, não são necessariamente erros
	ErroContinue = TipoErro.NewTipo("ErroContinue", "Erro utilizado para representar a instrução 'continue' em loops")
	ErroPare     = TipoErro.NewTipo("ErroPare", "Erro utilizado para representar a instrução 'pare' em loops")
)

func NewErro(tipo *Tipo, mensagem Objeto) *Erro {
	return &Erro{Base: tipo, Mensagem: mensagem}
}

func NewErroF(tipo *Tipo, format string, p ...any) *Erro {
	return &Erro{Base: tipo, Mensagem: Texto(fmt.Sprintf(format, p...))}
}

func (e *Erro) Tipo() *Tipo {
	return e.Base
}

func (e *Erro) AdicionarContexto(contexto *Contexto) {
	if e.Contexto != nil {
		return
	}

	e.Contexto = contexto
}

func (e *Erro) Error() string {
	format := "Arquivo %v:\n  %v: %v"

	// FIXME: corrigir o caminho
	return fmt.Sprintf(format, "e.Contexto.Caminho", e.Base.Nome, e.Mensagem.(Texto))
}
