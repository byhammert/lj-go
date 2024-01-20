package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasLancamentos = []Rota{
	{
		URI:                "/lancamentos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarLancamento,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos/{lancamentoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarLancamentoPorId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos/parcelamento",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarParcelaLancamento,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos/agendamento",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarAgendamentoLancamento,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarLancamento,
		RequerAutenticacao: true,
	},
}
