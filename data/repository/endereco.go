package repository

import (
	"database/sql"
	"side/core/storage"
	"side/data/model"
	"side/data/query"

	log "github.com/sirupsen/logrus"
)

// Selecionar generico para listas
func selecionarEnderecos(rows *sql.Rows, err error) ([]model.Endereco, error) {
	if err != nil {
		log.Error("> Erro no exec: ", err)
		return nil, err
	}
	registros := make([]model.Endereco, 0)
	for rows.Next() {
		var registro model.Endereco
		err = rows.Scan(registro.GetParams()...)
		if err != nil {
			log.Error("Erro no scan: ", err)
			return nil, err
		}
		registros = append(registros, registro)
	}
	return registros, nil
}

// Metodos do CRUD

// InserirEndereco insere um sideufg_endereco no banco de dados
func InserirEndereco(registro model.Endereco) ( *model.Endereco, error ) {
	tx, err := storage.DB.Begin()
	if err != nil {
		log.Error("Erro ao iniciar transação: ", err)
		return nil, err
	}	
	row := tx.QueryRow(query.InserirEndereco, registro.GetParams()[1:]...)
	retorno := new(model.Endereco)
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

// ExcluirEndereco - deleta um sideufg_endereco do banco de dados
func ExcluirEndereco(registro model.Endereco) (err error) {
	_, err = storage.DB.Exec(query.ExcluirEndereco, registro.ID)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// AtualizarEndereco - Atualizar atualiza as informações de um sideufg_endereco na base de dados
func AtualizarEndereco(registro model.Endereco) (err error) {
	_, err = storage.DB.Exec(query.AtualizarEndereco, registro.GetParams()...)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// ObterEndereco seleciona um sideufg_endereco da base de dados através da PK
func ObterEndereco(registro model.Endereco) (*model.Endereco, error) {
	row := storage.DB.QueryRow(query.ObterEndereco, registro.ID)
	var output model.Endereco
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

// ObterEnderecos seleciona todos os sideufg_enderecos da base de dados
func ObterEnderecos() ([]model.Endereco, error) {
	registros, err := selecionarEnderecos(storage.DB.Query(query.ObterEnderecos))
	return registros, err
}
