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

// InserirTipoMedidor é controller associado a rota responsável por adicionar um elemento da tabela sideufgtipomedidor
func InserirTipoMedidor(c echo.Context) error {
	registro := &model.TipoMedidor{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	registro, err = repository.InserirTipoMedidor(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, registro)
}

// ExcluirTipoMedidor é o controller associado a rota responsável por remover um elemento da tabela sideufgtipomedidor
func ExcluirTipoMedidor(c echo.Context) error {
	var err error
	id := new(model.NullInt64)
	idStr, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Erro na conversao de parametro: ", err)
		return err
	}
	id.Int64 = int64(idStr)
	err = repository.ExcluirTipoMedidor(model.TipoMedidor{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// AtualizarTipoMedidor é o controller associado a rota responsável por atualizar um elemento da tabela sideufgtipomedidor
func AtualizarTipoMedidor(c echo.Context) error {
	registro := &model.TipoMedidor{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	err = repository.AtualizarTipoMedidor(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// ObterTipoMedidor é o controller associado a rota responsável por selecionar um elemento da tabela sideufgtipomedidor
func ObterTipoMedidor(c echo.Context) error {
	id := new(model.NullInt64)
	idStr, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Erro na conversao de parametro: ", err)
		return err
	}
	id.Int64 = int64(idStr)
	retorno, err := repository.ObterTipoMedidor(model.TipoMedidor{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// ObterTiposMedidor é o controller associado a rota responsável por selecionar todos os elementos na tabela sideufgtipomedidor
func ObterTiposMedidor(c echo.Context) error {
	retorno, err := repository.ObterTiposMedidor()
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

