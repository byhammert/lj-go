package modelos

type Parcelamento struct {
	Tipo       string     `json:"tipo,omitempty"`
	Quantidade int32      `json:"quantidade,omitempty"`
	Lancamento Lancamento `json:"lancamento,omitempty"`
}
