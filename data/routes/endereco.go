package routes

import (
	"github.com/labstack/echo/v4"

	"side/data/service"
)

// EnderecoRoutes configura as rotas do módulo
func EnderecoRoutes(g *echo.Group) {
	g.POST("endereco", service.InserirEndereco)
	g.GET("endereco/:id", service.ObterEndereco)
	g.GET("endereco", service.ObterEnderecos)
	g.PUT("endereco", service.AtualizarEndereco)
	g.DELETE("endereco/:id", service.ExcluirEndereco)
}
