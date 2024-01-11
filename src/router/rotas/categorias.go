package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasCategorias = []Rota{
	{
		URI:                "/categorias",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCategoriaPorUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/{categoriaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCategoiriaPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/{categoriaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/categorias/{categoriaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCategoria,
		RequerAutenticacao: true,
	},
}
