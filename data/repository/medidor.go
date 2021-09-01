package repository

import (
	"database/sql"
	"side/core/storage"
	"side/data/model"
	"side/data/query"

	log "github.com/sirupsen/logrus"
)

// Selecionar generico para listas
func selecionarMedidores(rows *sql.Rows, err error) ([]model.Medidor, error) {
	if err != nil {
		log.Error("> Erro no exec: ", err)
		return nil, err
	}
	registros := make([]model.Medidor, 0)
	for rows.Next() {
		var registro model.Medidor
		err = rows.Scan(registro.GetParams()...)
		if err != nil {
			log.Error("Erro no scan: ", err)
			return nil, err
		}
		registros = append(registros, registro)
	}
	return registros, nil
}

// AtualizarMedidor - Atualizar atualiza as informações de um sideufg_medidor na base de dados
func AtualizarMedidor(registro model.Medidor) (err error) {
	_, err = storage.DB.Exec(query.AtualizarMedidor, registro.GetParams()...)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// ObterMedidor seleciona um sideufg_medidor da base de dados através da PK
func ObterMedidor(registro model.Medidor) (*model.Medidor, error) {
	row := storage.DB.QueryRow(query.ObterMedidor, registro.ID.Int64)
	var output model.Medidor
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

// ObterMedidores seleciona todos os sideufg_medidors da base de dados
func ObterMedidores() ([]model.Medidor, error) {
	registros, err := selecionarMedidores(storage.DB.Query(query.ObterMedidores))
	return registros, err
}
