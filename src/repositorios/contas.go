package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Contas struct {
	db *sql.DB
}

func NovoRepositorioDeContas(db *sql.DB) *Contas {
	return &Contas{db}
}

func (repositorio Contas) Criar(conta modelos.Conta) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into contas (nome, saldo, imagem, status_conta, tipo, usuario_id) values (?, ?, ?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(conta.Nome, conta.Saldo, conta.Imagem, conta.Status, conta.Tipo, conta.UsuarioID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Contas) Buscar(usuarioID uint64) ([]modelos.Conta, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT id, nome, saldo, imagem, status_conta, tipo, usuario_id FROM contas WHERE usuario_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var contas []modelos.Conta

	for linhas.Next() {
		var conta modelos.Conta

		if erro = linhas.Scan(
			&conta.ID,
			&conta.Nome,
			&conta.Saldo,
			&conta.Imagem,
			&conta.Status,
			&conta.Tipo,
			&conta.UsuarioID,
		); erro != nil {
			return nil, erro
		}

		contas = append(contas, conta)
	}

	return contas, nil
}

func (repositorio Contas) BuscarPorId(ID uint64) (modelos.Conta, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, nome, saldo, imagem, status_conta, tipo, usuario_id FROM contas where id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Conta{}, erro
	}

	defer linha.Close()

	var conta modelos.Conta

	if linha.Next() {
		if erro = linha.Scan(
			&conta.ID,
			&conta.Nome,
			&conta.Saldo,
			&conta.Imagem,
			&conta.Status,
			&conta.Tipo,
			&conta.UsuarioID,
		); erro != nil {
			return modelos.Conta{}, erro
		}
	}

	return conta, nil
}

func (repositorio Contas) Atualizar(ID uint64, conta modelos.Conta) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE contas SET nome = ?, saldo = ?, imagem = ?, status_conta = ?, tipo = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(conta.Nome, conta.Saldo, conta.Imagem, conta.Status, conta.Tipo, ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Contas) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM contas WHERE id = ?",
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
