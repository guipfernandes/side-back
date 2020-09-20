package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"side/core/erro"
	"side/core/utils"
	"side/data/repository"
	"time"
)

const (
	JwtKey = "CHAVEJWTDEVIASERSEGURA"
)
var errLoginInvalido = errors.New("login incorreto")

// Login é o controller responsável pela validaçao de login
func Login(c echo.Context) error {
	acesso := new(LoginInput)

	err := c.Bind(acesso)
	if err != nil {
		log.Info(err)
		return c.JSON(http.StatusUnauthorized, erro.Get(errLoginInvalido))
	}
	log.Debug(acesso)
	usuario, err := repository.ObterUsuarioAtivoPorLogin(acesso.Login)
	if err != nil {
		return c.HTML(http.StatusUnauthorized, "Login incorreto")
	}
	if !utils.CheckPasswordHash(acesso.Senha, usuario.Senha) {
		return c.HTML(http.StatusUnauthorized, "Login incorreto")
	}
	// Set claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = usuario.Email
	claims["idUsuario"] = usuario.ID
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, erro.Get(err))
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

type LoginInput struct {
	Login string `json:"login"`
	Senha string `json:"senha"`
}