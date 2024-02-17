package modelos

import (
	"errors"
)

type Parcelamento struct {
	Tipo       int        `json:"tipo,omitempty"`
	Quantidade string     `json:"quantidade,omitempty"`
	Lancamento Lancamento `json:"lancamento,omitempty"`
}

func (parcelamento *Parcelamento) Preparar() error {
	if erro := parcelamento.validar(); erro != nil {
		return erro
	}

	if erro := parcelamento.formatar(); erro != nil {
		return erro
	}
	return nil
}

func (parcelamento *Parcelamento) validar() error {
	if parcelamento.Tipo == 0 {
		return errors.New("O tipo é obrigatório e não pode estar em branco")
	}

	return parcelamento.Lancamento.Preparar()
}

func (parcelamento *Parcelamento) formatar() error {
	if parcelamento.Quantidade == "" {
		parcelamento.Quantidade = "1"
	}
	return nil
}
