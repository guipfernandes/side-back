package service

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"side/core/erro"
	"side/data/model"
	"side/data/repository"
)

// InserirUsuario é controller associado a rota responsável por adicionar um elemento da tabela sideufg_usuario
func InserirUsuario(c echo.Context) error {
	registro := &model.Usuario{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	registro, err = repository.InserirUsuario(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, registro)
}

// ExcluirUsuario é o controller associado a rota responsável por remover um elemento da tabela sideufg_usuario
func ExcluirUsuario(c echo.Context) error {
	var err error
	id := new(model.NullString)
	id.String = c.Param("id")
	err = repository.ExcluirUsuario(model.Usuario{ID: *id})
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

// AtualizarUsuario é o controller associado a rota responsável por atualizar um elemento da tabela sideufg_usuario
func AtualizarUsuario(c echo.Context) error {
	registro := &model.Usuario{}
	err := c.Bind(registro)
	if err != nil {
		log.Error("Erro no Bind: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	err = repository.AtualizarUsuario(*registro)
	if err != nil {
		log.Error("Erro no Repository: ", err)
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.NoContent(http.StatusOK)
}

