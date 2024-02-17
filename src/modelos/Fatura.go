package modelos

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type Fatura struct {
	ID             uint64          `json:"id,omitempty"`
	CartaoID       uint64          `json:"cartaoID,omitempty"`
	DataVencimento time.Time       `json:"dataVencimento,omitempty"`
	DataPagamento  NullTime        `json:"dataPagamento,omitempty"`
	FaturaFechada  bool            `json:"faturaFechada,omitempty"`
	Valor          decimal.Decimal `json:"valor,omitempty"`
	CodigoFatura   uint64          `json:"codigoFatura,omitempty"`
}

func (fatura *Fatura) Preparar() error {
	if erro := fatura.validar(); erro != nil {
		return erro
	}

	return nil
}

func (fatura *Fatura) validar() error {
	if fatura.DataVencimento.IsZero() {
		return errors.New("A data de vencimento é obrigatório e não pode estar em branco")
	}

	return nil
}
