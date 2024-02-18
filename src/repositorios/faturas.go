package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Faturas struct {
	db *sql.DB
}

func NovoRepositorioDeFaturas(db *sql.DB) *Faturas {
	return &Faturas{db}
}

func (repositorio Faturas) Criar(fatura modelos.Fatura) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into faturas (id_cartao, data_vencimento, data_pagamento, fatura_fechada, valor, codigo_fatura) values (?, ?, ?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(fatura.CartaoID, fatura.DataVencimento, fatura.DataPagamento, fatura.FaturaFechada, fatura.Valor, fatura.CodigoFatura)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Faturas) BuscarPorCartao(cartaoID uint64) ([]modelos.Fatura, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT id, id_cartao, data_vencimento, data_pagamento, fatura_fechada, valor, codigo_fatura FROM faturas WHERE id_cartao = ?`,
		cartaoID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var faturas []modelos.Fatura

	for linhas.Next() {
		var fatura modelos.Fatura

		if erro = linhas.Scan(
			&fatura.ID,
			&fatura.CartaoID,
			&fatura.DataVencimento,
			&fatura.DataPagamento,
			&fatura.FaturaFechada,
			&fatura.Valor,
			&fatura.CodigoFatura,
		); erro != nil {
			return nil, erro
		}

		faturas = append(faturas, fatura)
	}

	return faturas, nil
}

func (repositorio Faturas) BuscarPorId(ID uint64) (modelos.Fatura, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, id_cartao, data_vencimento, data_pagamento, fatura_fechada, valor, codigo_fatura FROM faturas where id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Fatura{}, erro
	}

	defer linha.Close()

	var fatura modelos.Fatura

	if linha.Next() {
		if erro = linha.Scan(
			&fatura.ID,
			&fatura.CartaoID,
			&fatura.DataVencimento,
			&fatura.DataPagamento,
			&fatura.FaturaFechada,
			&fatura.Valor,
			&fatura.CodigoFatura,
		); erro != nil {
			return modelos.Fatura{}, erro
		}
	}

	return fatura, nil
}

func (repositorio Faturas) Atualizar(ID uint64, fatura modelos.Fatura) error {
	statement, erro := repositorio.db.Prepare(
		`UPDATE faturas SET 
			data_vencimento = ?, 
			data_pagamento = ?, 
			fatura_fechada = ?, 
			valor = ? 
		WHERE id = ?`,
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(fatura.DataVencimento, fatura.DataPagamento, fatura.FaturaFechada, fatura.Valor, ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Faturas) BuscarFaturaAtual() (modelos.Fatura, error) {
	linha, erro := repositorio.db.Query(
		`SELECT id, 
			id_cartao, 
			data_vencimento, 
			data_pagamento, 
			fatura_fechada, 
			valor, 
			codigo_fatura 
		FROM faturas 
		where fatura_fechada != true`,
	)

	if erro != nil {
		return modelos.Fatura{}, erro
	}

	defer linha.Close()

	var fatura modelos.Fatura

	if linha.Next() {
		if erro = linha.Scan(
			&fatura.ID,
			&fatura.CartaoID,
			&fatura.DataVencimento,
			&fatura.DataPagamento,
			&fatura.FaturaFechada,
			&fatura.Valor,
			&fatura.CodigoFatura,
		); erro != nil {
			return modelos.Fatura{}, erro
		}
	}

	return fatura, nil
}

func (repositorio Faturas) BuscarFaturasPorCodigo(codigos []uint64) ([]modelos.Fatura, error) {
	placeholders := make([]string, len(codigos))
	for i, codigo := range codigos {
		placeholders[i] = strconv.FormatUint(codigo, 10)
	}

	placeholdersString := strings.Join(placeholders, ",")
	query := fmt.Sprintf("SELECT id, id_cartao, data_vencimento, data_pagamento, fatura_fechada, valor, codigo_fatura FROM faturas WHERE codigo_fatura IN (%s)", placeholdersString)
	linhas, erro := repositorio.db.Query(query)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var faturas []modelos.Fatura

	for linhas.Next() {
		var fatura modelos.Fatura

		if erro = linhas.Scan(
			&fatura.ID,
			&fatura.CartaoID,
			&fatura.DataVencimento,
			&fatura.DataPagamento,
			&fatura.FaturaFechada,
			&fatura.Valor,
			&fatura.CodigoFatura,
		); erro != nil {
			return nil, erro
		}

		faturas = append(faturas, fatura)
	}

	return faturas, nil
}
