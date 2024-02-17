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
		URI:                "/lancamentos/{lancamentoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarLancamento,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos/parcelamento",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarParcelaLancamento,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos/{periodo}/por-mes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarLancamentoDoMes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/lancamentos/{periodo}/despesas-do-mes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDespesasDoMes,
		RequerAutenticacao: true,
	},
}
