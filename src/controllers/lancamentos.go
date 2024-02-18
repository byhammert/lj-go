package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/config"
	"api/src/modelos"
	"api/src/modelos/enums"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

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
	if lancamento.DataPagamento.Valid {
		repositorioConta := repositorios.NovoRepositorioDeContas(db)
		conta, erro := repositorioConta.BuscarPorId(lancamento.ContaID)
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

func AtualizarLancamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

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

	if lancamento.UsuarioID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um lancamento que não seja o seu"))
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var lancamentoAtualizado modelos.Lancamento
	if erro = json.Unmarshal(corpoRequest, &lancamentoAtualizado); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = lancamentoAtualizado.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	lancamentoAtualizado.ContaID = lancamento.ContaID

	// ATUALIZAR CONTA
	if lancamentoAtualizado.DataPagamento.Valid && !lancamento.DataPagamento.Valid {
		repositorioConta := repositorios.NovoRepositorioDeContas(db)
		conta, erro := repositorioConta.BuscarPorId(lancamento.ContaID)
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
	}

	erro = repositorio.Atualizar(lancamentoID, lancamentoAtualizado)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeletarLancamento(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

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

	if lancamento.UsuarioID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um lancamento que não seja o seu"))
		return
	}

	// ATUALIZAR CONTA
	if lancamento.DataPagamento.Valid {
		repositorioConta := repositorios.NovoRepositorioDeContas(db)
		conta, erro := repositorioConta.BuscarPorId(lancamento.ContaID)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		novoSaldo := conta.Saldo
		if lancamento.Tipo != "RECEITA" {
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
	}

	erro = repositorio.Deletar(lancamentoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
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

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorioLancamento := repositorios.NovoRepositorioDeLancamentos(db)

	quantidade, erro := strconv.Atoi(parcelamento.Quantidade)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if parcelamento.Tipo == enums.Recorrente {
		quantidade = config.ParcelasRecorrentes
	}

	quantidadeParcelas := decimal.NewFromInt32(int32(quantidade))
	valorParcela := parcelamento.Lancamento.Valor.Div(quantidadeParcelas)
	mes := 1

	for i := 0; i < quantidade; i++ {
		lancamento := parcelamento.Lancamento
		if parcelamento.Tipo != enums.Recorrente {
			descricaoParcela := fmt.Sprintf("%d/%d", i+1, parcelamento.Quantidade)
			lancamento.Descricao = fmt.Sprintf("%s - %s", lancamento.Descricao, descricaoParcela)
			lancamento.Valor = valorParcela
		}

		if i > 0 {
			lancamento.DataVencimento = lancamento.DataVencimento.AddDate(0, mes, 0)
			mes++
		}
		_, erro = repositorioLancamento.Criar(lancamento)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarLancamentoDoMes(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	layout := "2006-01-02"
	periodo, erro := time.Parse(layout, parametros["periodo"])
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeLancamentos(db)
	lancamentos, erro := repositorio.BuscarLancamentosDoMes(usuarioID, periodo)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, lancamentos)
}

func BuscarDespesasDoMes(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	layout := "2006-01-02"
	periodo, erro := time.Parse(layout, parametros["periodo"])
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeLancamentos(db)
	lancamentos, erro := repositorio.BuscarLancamentosDoMesNaoPagasESemFatura(usuarioID, periodo)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var despesas []modelos.Despesa
	for _, lancamento := range lancamentos {
		var despesa modelos.Despesa
		despesa.ID = lancamento.ID
		despesa.DataVencimento = lancamento.DataVencimento
		despesa.Descricao = lancamento.Descricao
		despesa.Tipo = lancamento.Tipo
		despesa.Valor = lancamento.Valor
		despesas = append(despesas, despesa)
	}

	repositorioFatura := repositorios.NovoRepositorioDeFaturas(db)
	fatura, erro := repositorioFatura.BuscarFaturaAtual()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if (fatura != modelos.Fatura{}) {
		var despesa modelos.Despesa
		despesa.ID = fatura.ID
		despesa.DataVencimento = fatura.DataVencimento
		despesa.Descricao = "CARTÃO DE CRÉDITO"
		despesa.Tipo = "FATURA"
		despesa.Valor = fatura.Valor
		despesas = append(despesas, despesa)
	}

	respostas.JSON(w, http.StatusOK, despesas)
}
