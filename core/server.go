package core

import (
	"errors"
	"io"
	"net/http"
	"os"
	"side/core/config"
	"side/core/erro"
	"side/core/routes"
	"side/core/security"
	"side/core/storage"
	"side/data/service"
	"side/logger"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	glog "github.com/labstack/gommon/log"
	log "github.com/sirupsen/logrus"
)

type API struct {
	*echo.Echo
}

// LoadDefault carrega as informações e faz
// as configurações padrões do serviço echo
func (api *API) LoadDefault() *API {

	api.Debug = true
	api.HideBanner = false
	api.Server.Addr = ":9000"

	api.HTTPErrorHandler = func(err error, c echo.Context) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		} else {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}

		// Issue #1426
		code := he.Code
		message := he.Message
		if m, ok := message.(string); ok {
			message = m
		}

		// Send response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead { // Issue #608
				err = c.NoContent(he.Code)
			} else {
				err = c.JSON(code, erro.Get(errors.New(message.(string)), strconv.Itoa(code)))
			}
		}
	}

	api.Use(
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
		}),
	)

	api.POST("/login", security.Login)
	api.POST("/cadastrar", service.InserirUsuario)

	return api.CreateGroupV1()
}

// CreateGroupV1 cria o grupo autenticado via JWT
func (api *API) CreateGroupV1() *API {

	// Grupo que será autenticado via JWT.
	v1 := api.Group("/api/v1/")

	// Configuração jwt.
	v1.Use(
		middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(security.JwtKey),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSON(http.StatusUnauthorized, erro.Get(err, strconv.Itoa(http.StatusUnauthorized)))
			},
		}),
	)

	routes.Routes(v1)

	return api
}

// StartAPI abre e testa a conexão com o banco de dados assim como inicia o serviço Echo
func (api *API) StartAPI() {
	var cfg config.Config
	if err := cfg.LoadFromJSON("config.json"); err != nil {
		log.Fatal(err)
		return
	}

	logger.CarregaConfiguracoes(cfg.LogConfig)

	defaultLogger := logger.ConfiguracaoPadrao()
	log.SetLevel(logger.Config.ObtemLogLevel())
	mw := io.MultiWriter(os.Stdout, defaultLogger)
	log.SetOutput(mw)

	api.Logger.SetLevel(glog.OFF)
	api.Logger.SetOutput(mw)
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:           mw,
		CustomTimeFormat: logger.Config.Formatter.TimestampFormat,
		Format:           `[${time_custom}] [ECHO ] [core/server.go:0] [id=${id}] [remote_ip=${remote_ip}] ` + `[status=${status}] [error=${error}] [protocol=${protocol}] [agent=${user_agent}] ` + `[latency=${latency}] [latency_human=${latency_human}] [bytes_in=${bytes_in}] ` + `[bytes_out=${bytes_out}] – ${method} ${host}${uri}` + "\n",
	}))

	if err := storage.New(&cfg); err != nil {
		log.Fatal(err)
		return
	}
	if err := api.Start(api.Server.Addr); err != nil {
		log.Error(err)
	}
}
