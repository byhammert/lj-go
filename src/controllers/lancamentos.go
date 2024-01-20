package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/config"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

func CriarLancamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var lancamento modelos.Lancamento
	if erro = json.Unmarshal(corpoRequest, &lancamento); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = lancamento.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	lancamento.UsuarioID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	// ATUALIZAR CONTA
	repositorioConta := repositorios.NovoRepositorioDeContas(db)
	conta, erro := repositorioConta.BuscarPorId(lancamento.CantaID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	novoSaldo := conta.Saldo
	if lancamento.Tipo == "RECEITA" {
		novoSaldo = novoSaldo.Add(lancamento.Valor)
	} else {
		novoSaldo = novoSaldo.Sub(lancamento.Valor)
	}

	conta.Saldo = novoSaldo

	erro = repositorioConta.Atualizar(conta.ID, conta)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// CRIAR LANCAMENTO
	repositorioLancamento := repositorios.NovoRepositorioDeLancamentos(db)
	lancamento.ID, erro = repositorioLancamento.Criar(lancamento)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, lancamento)
}

func BuscarLancamentoPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	lancamentoID, erro := strconv.ParseUint(parametros["lancamentoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeLancamentos(db)
	lancamento, erro := repositorio.BuscarPorID(lancamentoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, lancamento)
}

func CriarParcelaLancamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var parcelamento modelos.Parcelamento
	if erro = json.Unmarshal(corpoRequest, &parcelamento); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = parcelamento.Lancamento.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parcelamento.Lancamento.UsuarioID = usuarioID
	parcelamento.Lancamento.Agendada = true

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorioLancamento := repositorios.NovoRepositorioDeLancamentos(db)

	if parcelamento.Tipo == "RECORRENTE" {
		parcelamento.Quantidade = int32(config.ParcelasRecorrentes)
	}
	quantidadeParcelas := decimal.NewFromInt32(parcelamento.Quantidade)
	valorParcela := parcelamento.Lancamento.Valor.Div(quantidadeParcelas)
	for i := 0; i < int(parcelamento.Quantidade); i++ {
		lancamento := parcelamento.Lancamento
		if parcelamento.Tipo != "RECORRENTE" {
			descricaoParcela := fmt.Sprintf("%d/%d", i+1, parcelamento.Quantidade)
			lancamento.Descricao = fmt.Sprintf("%s - %s", lancamento.Descricao, descricaoParcela)
			lancamento.Valor = valorParcela
		}

		if i > 0 {
			lancamento.DataVencimento = lancamento.DataVencimento.AddDate(0, 1, 0)
		}
		_, erro = repositorioLancamento.Criar(lancamento)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func CriarAgendamentoLancamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var lancamento modelos.Lancamento
	if erro = json.Unmarshal(corpoRequest, &lancamento); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = lancamento.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	lancamento.UsuarioID = usuarioID
	lancamento.Agendada = true

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorioLancamento := repositorios.NovoRepositorioDeLancamentos(db)
	lancamento.ID, erro = repositorioLancamento.Criar(lancamento)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, lancamento)
}
