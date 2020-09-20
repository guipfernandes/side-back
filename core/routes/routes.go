
package routes

import (
	"github.com/labstack/echo/v4"
	"side/data/routes"
)

// Routes configura as rotas do módulos
func Routes(g *echo.Group) {
	routes.EnderecoRoutes(g)
	routes.MedicaoRoutes(g)
	routes.MedidorRoutes(g)
	routes.MedidorEnelRoutes(g)
	routes.TipoMedidorRoutes(g)
	routes.UsuarioRoutes(g)
}
