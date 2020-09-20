package repository


import (
	"database/sql"
	"side/core/storage"
	"side/data/model"
	"side/data/query"

	log "github.com/sirupsen/logrus"
)

// Selecionar generico para listas
func SelecionarMedicoes(rows *sql.Rows, err error) ([]model.Medicao, error) {
	if err != nil {
		log.Error("> Erro no exec: ", err)
		return nil, err
	}
	registros := make([]model.Medicao, 0)
	for rows.Next() {
		var registro model.Medicao
		err = rows.Scan(registro.GetParams()...)
		if err != nil {
			log.Error("Erro no scan: ", err)
			return nil, err
		}
		registros = append(registros, registro)
	}
	return registros, nil
}

// ObterMedicao seleciona um sideufg_medicao da base de dados através da PK
func ObterMedicao(registro model.Medicao) (*model.Medicao, error) {
	row := storage.DB.QueryRow(query.ObterMedicao, registro.ID)
	var output model.Medicao
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

// ObterMedicoes seleciona todos os sideufg_medicaos da base de dados
func ObterMedicoes() ([]model.Medicao, error) {
	registros, err := SelecionarMedicoes(storage.DB.Query(query.ObterMedicoes))
	return registros, err
}
	