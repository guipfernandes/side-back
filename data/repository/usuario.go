package repository

import (
	"database/sql"
	"side/core/storage"
	"side/core/utils"
	"side/data/model"
	"side/data/query"

	log "github.com/sirupsen/logrus"
)

// InserirUsuario insere um sideufg_usuario no banco de dados
func InserirUsuario(registro model.Usuario) ( *model.Usuario, error ) {
	tx, err := storage.DB.Begin()
	if err != nil {
		log.Error("Erro ao iniciar transação: ", err)
		return nil, err
	}
	registro.Situacao = 1
	registro.Senha, _ = utils.HashPassword(registro.Senha)
	row := tx.QueryRow(query.InserirUsuario, registro.GetParams()[1:]...)
	retorno := new(model.Usuario)
	err = row.Scan(retorno.GetParams()...)
	if err != nil {
		log.Error("Erro no scan: ", err)
		_ = tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		log.Error("Erro ao commitar: ", err)
		return nil, err
	}
	return retorno, nil
}

// ExcluirUsuario - deleta um sideufg_usuario do banco de dados
func ExcluirUsuario(registro model.Usuario) (err error) {
	_, err = storage.DB.Exec(query.ExcluirUsuario, registro.ID)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// AtualizarUsuario - Atualizar atualiza as informações de um sideufg_usuario na base de dados
func AtualizarUsuario(registro model.Usuario) (err error) {
	_, err = storage.DB.Exec(query.AtualizarUsuario, registro.GetParams()...)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// ObterUsuario seleciona um sideufg_usuario da base de dados através da PK
func ObterUsuarioAtivoPorLogin(login string) (*model.Usuario, error) {
	row := storage.DB.QueryRow(query.ObterUsuarioAtivoPorLogin, login)
	var output model.Usuario
	err := row.Scan(output.GetParams()...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error("Erro no scan: ", err)
		return nil, err
	}
	return &output, nil
}