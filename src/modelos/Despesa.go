package modelos

import (
	"time"

	"github.com/shopspring/decimal"
)

type Despesa struct {
	ID             uint64          `json:"id,omitempty"`
	Descricao      string          `json:"descricao,omitempty"`
	Valor          decimal.Decimal `json:"valor,omitempty"`
	DataVencimento time.Time       `json:"dataVencimento,omitempty"`
	Tipo           string          `json:"tipo,omitempty"`
}
