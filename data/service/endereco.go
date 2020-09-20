package service

import (
	"net/http"
	"side/core/erro"
	"side/data/model"
	"side/data/repository"
	
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// InserirEndereco é controller associado a rota responsável por adicionar um elemento da tabela sideufgendereco
func InserirEndereco(c echo.Context) error {
	registro := &model.Endereco{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	registro, err = repository.InserirEndereco(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, registro)
}

// ExcluirEndereco é o controller associado a rota responsável por remover um elemento da tabela sideufgendereco
func ExcluirEndereco(c echo.Context) error {
	var err error
	id := new(model.NullString)
	id.String = c.Param("id")
	err = repository.ExcluirEndereco(model.Endereco{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// AtualizarEndereco é o controller associado a rota responsável por atualizar um elemento da tabela sideufgendereco
func AtualizarEndereco(c echo.Context) error {
	registro := &model.Endereco{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	err = repository.AtualizarEndereco(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// ObterEndereco é o controller associado a rota responsável por selecionar um elemento da tabela sideufgendereco
func ObterEndereco(c echo.Context) error {
	id := new(model.NullString)
	id.String = c.Param("id")
	retorno, err := repository.ObterEndereco(model.Endereco{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

// ObterEnderecos é o controller associado a rota responsável por selecionar todos os elementos na tabela sideufgendereco
func ObterEnderecos(c echo.Context) error {
	retorno, err := repository.ObterEnderecos()
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, retorno)
}

