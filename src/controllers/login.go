package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorios.BuscarPorEmail(usuario.Email)
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{ID: usuarioID, Token: token})
}
