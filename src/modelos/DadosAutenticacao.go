package modelos

type DadosAutenticacao struct {
	ID    uint64 `json:"id,omitempty"`
	Token string `json:"token,omitempty"`
}
