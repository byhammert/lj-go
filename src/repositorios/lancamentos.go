package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
	"time"
)

type Lancamentos struct {
	db *sql.DB
}

func NovoRepositorioDeLancamentos(db *sql.DB) *Lancamentos {
	return &Lancamentos{db}
}

func (repositorio Lancamentos) Criar(lancamento modelos.Lancamento) (uint64, error) {
	var query string
	if lancamento.FaturaID < 1 {
		query = `insert into lancamentos
		(descricao, valor, data_compra, data_vencimento, data_pagamento, tipo, forma_pagamento, id_categoria, id_usuario, id_conta, detalhe)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	} else {
		query = `insert into lancamentos
		(descricao, valor, data_compra, data_vencimento, data_pagamento, tipo, forma_pagamento, id_categoria, id_usuario, id_conta, detalhe, id_fatura)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	}
	statement, erro := repositorio.db.Prepare(
		query,
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()
	if lancamento.FaturaID < 1 {
		resultado, erro := statement.Exec(lancamento.Descricao, lancamento.Valor, lancamento.DataCompra, lancamento.DataVencimento,
			lancamento.DataPagamento, lancamento.Tipo, lancamento.FormaPagamento, lancamento.CategoriaID, lancamento.UsuarioID, lancamento.ContaID, lancamento.Detalhe)
		if erro != nil {
			return 0, erro
		}

		ultimoIDInserido, erro := resultado.LastInsertId()
		if erro != nil {
			return 0, erro
		}

		return uint64(ultimoIDInserido), nil
	} else {
		resultado, erro := statement.Exec(lancamento.Descricao, lancamento.Valor, lancamento.DataCompra, lancamento.DataVencimento,
			lancamento.DataPagamento, lancamento.Tipo, lancamento.FormaPagamento, lancamento.CategoriaID, lancamento.UsuarioID, lancamento.ContaID, lancamento.Detalhe, lancamento.FaturaID)
		if erro != nil {
			return 0, erro
		}

		ultimoIDInserido, erro := resultado.LastInsertId()
		if erro != nil {
			return 0, erro
		}

		return uint64(ultimoIDInserido), nil
	}
}

func (repositorio Lancamentos) Atualizar(ID uint64, lancamento modelos.Lancamento) error {
	statement, erro := repositorio.db.Prepare(
		`UPDATE lancamentos SET 
			descricao = ?, 
			valor = ?, 
			data_compra = ?, 
			data_vencimento = ?, 
			data_pagamento = ?, 
			tipo = ?, 
			forma_pagamento = ?,
			id_categoria = ?,
			detalhe = ?
		WHERE id = ?`,
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(
		lancamento.Descricao,
		lancamento.Valor,
		lancamento.DataCompra,
		lancamento.DataVencimento,
		lancamento.DataPagamento,
		lancamento.Tipo,
		lancamento.FormaPagamento,
		lancamento.CategoriaID,
		lancamento.Detalhe,
		ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Lancamentos) BuscarPorID(ID uint64) (modelos.Lancamento, error) {
	linha, erro := repositorio.db.Query(`
		SELECT 
			l.id,
			l.descricao, 
			l.detalhe,
			l.valor, 
			l.data_compra, 
			l.data_vencimento, 
			l.data_pagamento, 
			l.tipo, 
			l.forma_pagamento, 
			l.id_categoria, 
			l.id_usuario, 
			l.id_conta,
			co.nome,
			ca.nome
		FROM
			lancamentos l
		INNER JOIN
			contas co
		ON
			l.id_conta = co.id
		INNER JOIN
			categorias ca
		ON
			l.id_categoria = ca.id
		WHERE
			l.id = ?
	`, ID)

	if erro != nil {
		return modelos.Lancamento{}, erro
	}

	defer linha.Close()

	var lancamento modelos.Lancamento

	if linha.Next() {
		if erro = linha.Scan(
			&lancamento.ID,
			&lancamento.Descricao,
			&lancamento.Detalhe,
			&lancamento.Valor,
			&lancamento.DataCompra,
			&lancamento.DataVencimento,
			&lancamento.DataPagamento,
			&lancamento.Tipo,
			&lancamento.FormaPagamento,
			&lancamento.CategoriaID,
			&lancamento.UsuarioID,
			&lancamento.ContaID,
			&lancamento.ContaNome,
			&lancamento.CategoriaNome,
		); erro != nil {
			return modelos.Lancamento{}, erro
		}
	}

	return lancamento, nil
}

func (repositorio Lancamentos) BuscarLancamentosDoMes(usuarioID uint64, periodo time.Time) ([]modelos.Lancamento, error) {
	dataInicio, dataFim, erro := obterPeriodoMes(periodo.Year(), int(periodo.Month()))
	if erro != nil {
		return nil, erro
	}

	fmt.Println(periodo)
	fmt.Println(dataInicio)
	fmt.Println(dataFim)
	linhas, erro := repositorio.db.Query(`
		SELECT 
			l.id,
			l.descricao, 
			l.detalhe, 
			l.valor, 
			l.data_compra, 
			l.data_vencimento, 
			l.data_pagamento, 
			l.tipo, 
			l.forma_pagamento, 
			l.id_categoria, 
			l.id_usuario, 
			l.id_conta,
			co.nome,
			ca.nome
		FROM
			lancamentos l
		INNER JOIN
			contas co
		ON
			l.id_conta = co.id
		INNER JOIN
			categorias ca
		ON
			l.id_categoria = ca.id
		WHERE
			l.id_usuario = ?
		AND 
			l.data_vencimento BETWEEN date(?) AND date(?)
	`, usuarioID, dataInicio, dataFim)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var lancamentos []modelos.Lancamento

	for linhas.Next() {
		var lancamento modelos.Lancamento
		if erro = linhas.Scan(
			&lancamento.ID,
			&lancamento.Descricao,
			&lancamento.Detalhe,
			&lancamento.Valor,
			&lancamento.DataCompra,
			&lancamento.DataVencimento,
			&lancamento.DataPagamento,
			&lancamento.Tipo,
			&lancamento.FormaPagamento,
			&lancamento.CategoriaID,
			&lancamento.UsuarioID,
			&lancamento.ContaID,
			&lancamento.ContaNome,
			&lancamento.CategoriaNome,
		); erro != nil {
			return nil, erro
		}

		lancamentos = append(lancamentos, lancamento)
	}

	return lancamentos, nil
}

func (repositorio Lancamentos) BuscarLancamentosDoMesNaoPagasESemFatura(usuarioID uint64, periodo time.Time) ([]modelos.Lancamento, error) {
	dataInicio, dataFim, erro := obterPeriodoMes(periodo.Year(), int(periodo.Month()))
	if erro != nil {
		return nil, erro
	}

	fmt.Println(periodo)
	fmt.Println(dataInicio)
	fmt.Println(dataFim)
	linhas, erro := repositorio.db.Query(`
		SELECT 
			l.id,
			l.descricao, 
			l.detalhe, 
			l.valor, 
			l.data_compra, 
			l.data_vencimento, 
			l.data_pagamento, 
			l.tipo, 
			l.forma_pagamento, 
			l.id_categoria, 
			l.id_usuario, 
			l.id_conta,
			co.nome,
			ca.nome
		FROM
			lancamentos l
		INNER JOIN
			contas co
		ON
			l.id_conta = co.id
		INNER JOIN
			categorias ca
		ON
			l.id_categoria = ca.id
		WHERE
			l.id_usuario = ?
		AND 
			l.data_vencimento BETWEEN date(?) AND date(?)
		AND 
			l.id_fatura IS NULL
		AND 
			l.data_pagamento IS NULL
	`, usuarioID, dataInicio, dataFim)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var lancamentos []modelos.Lancamento

	for linhas.Next() {
		var lancamento modelos.Lancamento
		if erro = linhas.Scan(
			&lancamento.ID,
			&lancamento.Descricao,
			&lancamento.Detalhe,
			&lancamento.Valor,
			&lancamento.DataCompra,
			&lancamento.DataVencimento,
			&lancamento.DataPagamento,
			&lancamento.Tipo,
			&lancamento.FormaPagamento,
			&lancamento.CategoriaID,
			&lancamento.UsuarioID,
			&lancamento.ContaID,
			&lancamento.ContaNome,
			&lancamento.CategoriaNome,
		); erro != nil {
			return nil, erro
		}

		lancamentos = append(lancamentos, lancamento)
	}

	return lancamentos, nil
}

func obterPeriodoMes(ano, mes int) (inicioMes, fimMes time.Time, err error) {
	// Construir o primeiro dia do mês
	dataInicial := time.Date(ano, time.Month(mes), 1, 0, 0, 0, 0, time.Local)

	// Construir o último dia do mês
	ultimoDiaMes := time.Date(ano, time.Month(mes)+1, 0, 0, 0, 0, 0, time.Local)

	return dataInicial, ultimoDiaMes, nil
}

func (repositorio Lancamentos) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from lancamentos where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(ID)
	if erro != nil {
		return erro
	}

	return nil
}
