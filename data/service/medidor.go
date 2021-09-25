package service

import (
	"net/http"
	"side/core/erro"
	"side/data/model"
	"side/data/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// AtualizarMedidor é o controller associado a rota responsável por atualizar um elemento da tabela sideufgmedidor
func AtualizarMedidor(c echo.Context) error {
	registro := &model.Medidor{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	err = repository.AtualizarMedidor(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// ObterMedidor é o controller associado a rota responsável por selecionar um elemento da tabela sideufgmedidor
func ObterMedidor(c echo.Context) error {
	id := new(model.NullInt64)
	idStr, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Erro na conversao de parametro: ", err)
		return err
	}
	id.Int64 = int64(idStr)
	retorno, err := repository.ObterMedidor(model.Medidor{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// ObterMedidores é o controller associado a rota responsável por selecionar todos os elementos na tabela sideufgmedidor
func ObterMedidores(c echo.Context) error {
	retorno, err := repository.ObterMedidores()
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// AtualizarMedidorLocalizacao é o controller associado a rota responsável por atualizar latitude e longitude da tabela sideufgmedidor
func AtualizarMedidorLocalizacao(c echo.Context) error {
	registro := &model.MedidorLocalizacao{}

	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}

	err = repository.AtualizarMedidorLocalizacao(*registro)

	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}
