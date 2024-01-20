package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Lancamento struct {
	ID             uint64          `json:"id,omitempty"`
	Descricao      string          `json:"descricao,omitempty"`
	Valor          decimal.Decimal `json:"valor,omitempty"`
	DataCompra     time.Time       `json:"data-compra,omitempty"`
	DataVencimento time.Time       `json:"data-vencimento,omitempty"`
	DataPagamento  NullTime        `json:"data-pagamento,omitempty"`
	Tipo           string          `json:"tipo,omitempty"`
	FormaPagamento string          `json:"forma-pagamento,omitempty"`
	CantaID        uint64          `json:"conta-id,omitempty"`
	CategoriaID    uint64          `json:"categoria-id,omitempty"`
	ContaNome      string          `json:"conta-nome,omitempty"`
	CategoriaNome  string          `json:"categoria-nome,omitempty"`
	UsuarioID      uint64          `json:"usuario-id,omitempty"`
	Agendada       bool            `json:"agendada,omitempty"`
}

func (lancamento *Lancamento) Preparar() error {
	if erro := lancamento.validar(); erro != nil {
		return erro
	}

	if erro := lancamento.formatar(); erro != nil {
		return erro
	}
	return nil
}

func (lancamento *Lancamento) validar() error {
	if lancamento.Descricao == "" {
		return errors.New("O descrição é obrigatório e não pode estar em branco")
	}

	if lancamento.Valor.IsNegative() {
		return errors.New("O valor é obrigatório")
	}

	if lancamento.DataCompra.IsZero() {
		return errors.New("O data da compra é obrigatório")
	}

	if lancamento.DataVencimento.IsZero() {
		return errors.New("O data de vencimento é obrigatório")
	}

	if lancamento.Tipo == "" {
		return errors.New("O tipo é obrigatório e não pode estar em branco")
	}

	if lancamento.FormaPagamento == "" {
		return errors.New("O forma de pagamento é obrigatório e não pode estar em branco")
	}

	if lancamento.CategoriaID < 0 {
		return errors.New("O categoria é obrigatória")
	}

	if lancamento.CantaID < 0 {
		return errors.New("O conta é obrigatória")
	}

	return nil
}

func (lancamento *Lancamento) formatar() error {
	lancamento.Descricao = strings.TrimSpace(lancamento.Descricao)

	lancamento.Tipo = strings.TrimSpace(lancamento.Tipo)
	lancamento.Tipo = strings.ToUpper(lancamento.Tipo)

	lancamento.FormaPagamento = strings.TrimSpace(lancamento.FormaPagamento)
	lancamento.FormaPagamento = strings.ToUpper(lancamento.FormaPagamento)

	return nil
}
