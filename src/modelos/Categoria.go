package modelos

import (
	"errors"
	"strings"
)

type Categoria struct {
	ID        uint64 `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Tipo      string `json:"tipo,omitempty"`
	UsuarioID uint64 `json:"usuarioID,omitempty"`
}

func (categoria *Categoria) Preparar() error {
	if erro := categoria.validar(); erro != nil {
		return erro
	}

	if erro := categoria.formatar(); erro != nil {
		return erro
	}
	return nil
}

func (categoria *Categoria) validar() error {
	if categoria.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if categoria.Tipo == "" {
		return errors.New("O tipo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (categoria *Categoria) formatar() error {
	categoria.Nome = strings.TrimSpace(categoria.Nome)
	categoria.Tipo = strings.TrimSpace(categoria.Tipo)
	categoria.Tipo = strings.ToUpper(categoria.Tipo)

	return nil
}
