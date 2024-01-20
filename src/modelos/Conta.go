package modelos

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type Conta struct {
	ID        uint64          `json:"id,omitempty"`
	Nome      string          `json:"nome,omitempty"`
	Saldo     decimal.Decimal `json:"saldo,omitempty"`
	Imagem    string          `json:"image,omitempty"`
	Status    string          `json:"status,omitempty"`
	Tipo      string          `json:"tipo,omitempty"`
	UsuarioID uint64          `json:"usuario-id,omitempty"`
}

func (conta *Conta) Preparar() error {
	if erro := conta.validar(); erro != nil {
		return erro
	}

	if erro := conta.formatar(); erro != nil {
		return erro
	}
	return nil
}

func (conta *Conta) validar() error {
	if conta.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if conta.Status == "" {
		return errors.New("O status é obrigatório e não pode estar em branco")
	}

	if conta.Tipo == "" {
		return errors.New("O tipo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (conta *Conta) formatar() error {
	conta.Nome = strings.TrimSpace(conta.Nome)
	conta.Tipo = strings.TrimSpace(conta.Tipo)
	conta.Tipo = strings.ToUpper(conta.Tipo)
	conta.Status = strings.TrimSpace(conta.Status)
	conta.Status = strings.ToUpper(conta.Status)

	if conta.Imagem != "" {
		conta.Imagem = strings.TrimSpace(conta.Imagem)
	}

	return nil
}
