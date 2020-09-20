package repository


import (
	"database/sql"
	"side/core/storage"
	"side/data/model"
	"side/data/query"

	log "github.com/sirupsen/logrus"
)

// Selecionar generico para listas
func selecionarTiposMedidor(rows *sql.Rows, err error) ([]model.TipoMedidor, error) {
	if err != nil {
		log.Error("> Erro no exec: ", err)
		return nil, err
	}
	registros := make([]model.TipoMedidor, 0)
	for rows.Next() {
		var registro model.TipoMedidor
		err = rows.Scan(registro.GetParams()...)
		if err != nil {
			log.Error("Erro no scan: ", err)
			return nil, err
		}
		registros = append(registros, registro)
	}
	return registros, nil
}

// InserirTipoMedidor insere um sideufg_tipo_medidor no banco de dados
func InserirTipoMedidor(registro model.TipoMedidor) ( *model.TipoMedidor, error ) {
	tx, err := storage.DB.Begin()
	if err != nil {
		log.Error("Erro ao iniciar transação: ", err)
		return nil, err
	}
	row := tx.QueryRow(query.InserirTipoMedidor, registro.GetParams()[1:]...)
	retorno := new(model.TipoMedidor)
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

// ExcluirTipoMedidor - deleta um sideufg_tipo_medidor do banco de dados
func ExcluirTipoMedidor(registro model.TipoMedidor) (err error) {
	_, err = storage.DB.Exec(query.ExcluirTipoMedidor, registro.ID)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// AtualizarTipoMedidor - Atualizar atualiza as informações de um sideufg_tipo_medidor na base de dados
func AtualizarTipoMedidor(registro model.TipoMedidor) (err error) {
	_, err = storage.DB.Exec(query.AtualizarTipoMedidor, registro.GetParams()...)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// ObterTipoMedidor seleciona um sideufg_tipo_medidor da base de dados através da PK
func ObterTipoMedidor(registro model.TipoMedidor) (*model.TipoMedidor, error) {
	row := storage.DB.QueryRow(query.ObterTipoMedidor, registro.ID)
	var output model.TipoMedidor
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

// ObterTiposMedidor seleciona todos os sideufg_tipo_medidor da base de dados
func ObterTiposMedidor() ([]model.TipoMedidor, error) {
	registros, err := selecionarTiposMedidor(storage.DB.Query(query.ObterTiposMedidor))
	return registros, err
}
