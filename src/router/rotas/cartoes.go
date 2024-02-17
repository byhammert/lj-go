package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasCartoes = []Rota{
	{
		URI:                "/cartoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCartao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cartoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCartoesPorUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cartoes/{cartaoId}/faturas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarFaturasPorCartao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cartoes/{codigoFatura}/lacamentos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.LancarDespesaNaFatura,
		RequerAutenticacao: true,
	},
}
