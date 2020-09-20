package routes

import (
	"github.com/labstack/echo/v4"

	"side/data/service"
)

// MedidorEnelRoutes configura as rotas do módulo
func MedidorEnelRoutes(g *echo.Group) {
	g.POST("medidorenel", service.InserirMedidorEnel)
	g.GET("medidorenel/:id", service.ObterMedidorEnel)
	g.GET("medidorenel", service.ObterMedidoresEnel)
	g.PUT("medidorenel", service.AtualizarMedidorEnel)
	g.DELETE("medidorenel/:id", service.ExcluirMedidorEnel)
}
