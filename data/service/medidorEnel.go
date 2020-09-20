package service

import (
	"net/http"
	"side/core/erro"
	"side/data/model"
	"side/data/repository"
	
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// InserirTipoMedidor é controller associado a rota responsável por adicionar um elemento da tabela sideufgmedidorenel
func InserirMedidorEnel(c echo.Context) error {
	registro := &model.MedidorEnel{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	registro, err = repository.InserirMedidorEnel(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, registro)
}

// ExcluirTipoMedidor é o controller associado a rota responsável por remover um elemento da tabela sideufgmedidorenel
func ExcluirMedidorEnel(c echo.Context) error {
	var err error
	id := new(model.NullString)
	id.String = c.Param("id")
	err = repository.ExcluirMedidorEnel(model.MedidorEnel{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// AtualizarMedidorEnel é o controller associado a rota responsável por atualizar um elemento da tabela sideufgmedidorenel
func AtualizarMedidorEnel(c echo.Context) error {
	registro := &model.MedidorEnel{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	err = repository.AtualizarMedidorEnel(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// ObterMedidorEnel é o controller associado a rota responsável por selecionar um elemento da tabela sideufgmedidorenel
func ObterMedidorEnel(c echo.Context) error {
	id := new(model.NullString)
	id.String = c.Param("id")
	retorno, err := repository.ObterMedidorEnel(model.MedidorEnel{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// ObterMedidoresEnel é o controller associado a rota responsável por selecionar todos os elementos na tabela sideufgmedidorenel
func ObterMedidoresEnel(c echo.Context) error {
	retorno, err := repository.ObterMedidoresEnel()
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

