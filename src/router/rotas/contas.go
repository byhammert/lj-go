package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasContas = []Rota{
	{
		URI:                "/contas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarConta,
		RequerAutenticacao: true,
	},
	{
		URI:                "/contas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarContasPorUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/contas/{contaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarContaPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/contas/{contaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarConta,
		RequerAutenticacao: true,
	},
	{
		URI:                "/contas/{contaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarConta,
		RequerAutenticacao: true,
	},
}
