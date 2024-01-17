package modelos

import "time"

type Lancamento struct {
	ID              uint64    `json:"id,omitempty"`
	Valor           float64   `json:"valor,omitempty"`
	DataCompra      time.Time `json:"data-compra,omitempty"`
	DataVencimento  time.Time `json:"data-vencimento,omitempty"`
	DataPagamento   time.Time `json:"data-pagamento,omitempty"`
	Tipo            string    `json:"tipo,omitempty"`
	Forma_Pagamento string    `json:"forma-pagamento,omitempty"`
	CantaID         uint64    `json:"conta-id,omitempty"`
	CategoriaID     uint64    `json:"categoria-id,omitempty"`
	ContaNome       string    `json:"conta-nome,omitempty"`
	CategoriaNome   string    `json:"categoria-nome,omitempty"`
}
