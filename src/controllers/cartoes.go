package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
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

func CriarCartao(w http.ResponseWriter, r *http.Request) {
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

	var cartao modelos.Cartao
	if erro = json.Unmarshal(corpoRequest, &cartao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = cartao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	cartao.UsuarioID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeCartoes(db)
	cartao.ID, erro = repositorio.Criar(cartao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, cartao)
}

func BuscarCartoesPorUsuario(w http.ResponseWriter, r *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeCartoes(db)
	cartoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, cartoes)
}

func BuscarFaturasPorCartao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	cartaoID, erro := strconv.ParseUint(parametros["cartaoId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeFaturas(db)
	faturas, erro := repositorio.BuscarPorCartao(cartaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, faturas)
}

func LancarDespesaNaFatura(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	codigoFatura, erro := strconv.ParseUint(parametros["codigoFatura"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

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

	if erro = parcelamento.Preparar(); erro != nil {
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

	quantidade, erro := strconv.Atoi(parcelamento.Quantidade)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	codigos := make([]uint64, quantidade)
	for i := 0; i < int(quantidade); i++ {
		codigos[i] = codigoFatura + uint64(i)
	}
	// FATURA
	repositorioFatura := repositorios.NovoRepositorioDeFaturas(db)
	faturas, erro := repositorioFatura.BuscarFaturasPorCodigo(codigos)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// CRIAR LANCAMENTO
	repositorioLancamento := repositorios.NovoRepositorioDeLancamentos(db)

	quantidadeParcelas := decimal.NewFromInt32(int32(quantidade))
	valorParcela := parcelamento.Lancamento.Valor.Div(quantidadeParcelas)
	parcelamento.Lancamento.Valor = valorParcela
	mes := 1

	cartaoID := faturas[0].CartaoID
	dataVencimento := faturas[0].DataVencimento
	for i := 0; i < quantidade; i++ {
		codigoFatura := codigos[i]
		fatura := modelos.Fatura{}
		lancamento := parcelamento.Lancamento

		if len(faturas) > i {
			fatura = faturas[i]
			if i == 0 {
				fatura.Valor = fatura.Valor.Add(valorParcela)

				erro = repositorioFatura.Atualizar(fatura.ID, fatura)
				if erro != nil {
					respostas.Erro(w, http.StatusInternalServerError, erro)
					return
				}

			} else {
				lancamento.DataVencimento = lancamento.DataVencimento.AddDate(0, mes, 0)
				fatura.Valor = fatura.Valor.Add(valorParcela)
				erro = repositorioFatura.Atualizar(fatura.ID, fatura)
				if erro != nil {
					respostas.Erro(w, http.StatusInternalServerError, erro)
					return
				}
				mes++
			}
		} else {
			lancamento.DataVencimento = lancamento.DataVencimento.AddDate(0, mes, 0)
			if fatura == (modelos.Fatura{}) {
				dataVencimento = dataVencimento.AddDate(0, mes, 0)
				fmt.Println(lancamento.DataVencimento)
				fmt.Println(mes)

				fatura = modelos.Fatura{
					0,
					cartaoID,
					dataVencimento,
					modelos.NullTime{},
					false,
					valorParcela,
					codigoFatura,
				}

				fmt.Println("new fatura")
				fmt.Println(fatura)

				fatura.ID, erro = repositorioFatura.Criar(fatura)
				if erro != nil {
					respostas.Erro(w, http.StatusInternalServerError, erro)
					return
				}
			}
			mes++
		}
		lancamento.FaturaID = fatura.ID
		lancamento.ID, erro = repositorioLancamento.Criar(lancamento)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
