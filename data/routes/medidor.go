package routes

import (
	"github.com/labstack/echo/v4"

	"side/data/service"
)

// MedidorRoutes configura as rotas do módulo
func MedidorRoutes(g *echo.Group) {
	g.GET("medidor/:id", service.ObterMedidor)
	g.GET("medidor", service.ObterMedidores)
	g.PUT("medidor", service.AtualizarMedidor)
}
