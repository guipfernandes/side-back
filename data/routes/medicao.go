package routes

import (
	"github.com/labstack/echo/v4"

	"side/data/service"
)

// MedicaoRoutes configura as rotas do módulo
func MedicaoRoutes(g *echo.Group) {
	g.GET("medicao/:id", service.ObterMedicao)
	g.GET("medicao", service.ObterMedicoes)
}
