package service

import (
	"net/http"
	"side/core/erro"
	"side/data/model"
	"side/data/repository"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ObterMedicao é o controller associado a rota responsável por selecionar um elemento da tabela sideufgmedicao
func ObterMedicao(c echo.Context) error {
	id := new(model.NullString)
	id.String = c.Param("id")
	retorno, err := repository.ObterMedicao(model.Medicao{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// ObterMedicoes é o controller associado a rota responsável por selecionar todos os elementos na tabela sideufgmedicao
func ObterMedicoes(c echo.Context) error {
	retorno, err := repository.ObterMedicoes()
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// ObterHoraMedicoesByMedidorEDataMedicao é o controller associado a rota responsável por selecionar os elementos na view vw_hora_medicao
func ObterHoraMedicoesByMedidorEDataMedicao(c echo.Context) error {
	nomeMedidor := c.QueryParam("nomeMedidor")
	dataMedicaoInicio, err := time.Parse(time.RFC3339, c.QueryParam("dataMedicaoInicio"))
	dataMedicaoFim, err := time.Parse(time.RFC3339, c.QueryParam("dataMedicaoFim"))
	if err != nil {
		log.Error("Erro na conversao de parametro: ", err)
		return err
	}

	retorno, err := repository.ObterHoraMedicoesByMedidorEDataMedicao(nomeMedidor, dataMedicaoInicio, dataMedicaoFim)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}
