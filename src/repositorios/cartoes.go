package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Cartoes struct {
	db *sql.DB
}

func NovoRepositorioDeCartoes(db *sql.DB) *Cartoes {
	return &Cartoes{db}
}

func (repositorio Cartoes) Criar(cartao modelos.Cartao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into cartoes (descricao, usuario_id) values (?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(cartao.Descricao, cartao.UsuarioID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Cartoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Cartao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT id, descricao, usuario_id FROM cartoes WHERE usuario_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var catoes []modelos.Cartao

	for linhas.Next() {
		var cartao modelos.Cartao

		if erro = linhas.Scan(
			&cartao.ID,
			&cartao.Descricao,
			&cartao.UsuarioID,
		); erro != nil {
			return nil, erro
		}

		catoes = append(catoes, cartao)
	}

	return catoes, nil
}
