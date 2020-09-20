package routes

import (
	"github.com/labstack/echo/v4"
	"side/data/service"
)

// UsuarioRoutes configura as rotas do módulo
func UsuarioRoutes(g *echo.Group) {
	g.PUT("usuario", service.AtualizarUsuario)
	g.DELETE("usuario/:id", service.ExcluirUsuario)
}
