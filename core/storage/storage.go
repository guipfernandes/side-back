package storage

import (
	"database/sql"
	"side/core/config"
	"strings"
)

// DB - é a variável global de acesso ao banco
var DB *sql.DB

// New cria uma instância do objeto de acesso e manipuloção da base de dados.
func New(config *config.Config) (err error) {
	if DB, err = sql.Open("postgres", config.ToString()); err != nil {
		return
	}
	if err = DB.Ping(); err != nil {
		return
	}
	return
}

// AdicionaAlias - Adiciona um alias nos campos passados (id_usuario > u.id_usuario)
func AdicionaAlias(campos string, alias string) string {
	arrCampos := strings.Split(campos, ",")
	var resCampos []string
	for _, campo := range arrCampos {
		resCampos = append(resCampos, alias + "." + strings.TrimSpace(campo))
	}
	return strings.Join(resCampos, ",")
}
