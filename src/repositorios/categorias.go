package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Categorias struct {
	db *sql.DB
}

func NovoRepositorioDeCategorias(db *sql.DB) *Categorias {
	return &Categorias{db}
}

func (repositorio Categorias) Criar(categoria modelos.Categoria) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into categorias (nome, tipo, usuario_id) values (?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(categoria.Nome, categoria.Tipo, categoria.UsuarioID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Categorias) Buscar(usuarioID uint64) ([]modelos.Categoria, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT id, nome, tipo FROM categorias WHERE usuario_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var categorias []modelos.Categoria

	for linhas.Next() {
		var categoria modelos.Categoria

		if erro = linhas.Scan(
			&categoria.ID,
			&categoria.Nome,
			&categoria.Tipo,
		); erro != nil {
			return nil, erro
		}

		categorias = append(categorias, categoria)
	}

	return categorias, nil
}

func (repositorio Categorias) BuscarPorId(ID uint64) (modelos.Categoria, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, nome, tipo, usuario_id FROM categorias where id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Categoria{}, erro
	}

	defer linha.Close()

	var categoria modelos.Categoria

	if linha.Next() {
		if erro = linha.Scan(
			&categoria.ID,
			&categoria.Nome,
			&categoria.Tipo,
			&categoria.UsuarioID,
		); erro != nil {
			return modelos.Categoria{}, erro
		}
	}

	return categoria, nil
}

func (repositorio Categorias) Atualizar(ID uint64, categoria modelos.Categoria) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE categorias SET nome = ?, tipo = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(categoria.Nome, categoria.Tipo, ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Categorias) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM categorias WHERE id = ?",
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
