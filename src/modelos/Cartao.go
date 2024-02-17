package modelos

import (
	"errors"
	"strings"
)

type Cartao struct {
	ID        uint64 `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	UsuarioID uint64 `json:"usuarioID,omitempty"`
}

func (cartao *Cartao) Preparar() error {
	if erro := cartao.validar(); erro != nil {
		return erro
	}

	if erro := cartao.formatar(); erro != nil {
		return erro
	}
	return nil
}

func (cartao *Cartao) validar() error {
	if cartao.Descricao == "" {
		return errors.New("O descrição é obrigatório e não pode estar em branco")
	}

	return nil
}

func (cartao *Cartao) formatar() error {
	cartao.Descricao = strings.TrimSpace(cartao.Descricao)

	return nil
}
