package routes

import (
	"github.com/labstack/echo/v4"
	"side/data/service"
)

// TipoMedidorRoutes configura as rotas do módulo
func TipoMedidorRoutes(g *echo.Group) {
	g.POST("tipomedidor", service.InserirTipoMedidor)
	g.GET("tipomedidor/:id", service.ObterTipoMedidor)
	g.GET("tipomedidor", service.ObterTiposMedidor)
	g.PUT("tipomedidor", service.AtualizarTipoMedidor)
	g.DELETE("tipomedidor/:id", service.ExcluirTipoMedidor)
}
