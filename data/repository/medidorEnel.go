package repository


import (
	"database/sql"
	"side/core/storage"
	"side/data/model"
	"side/data/query"

	log "github.com/sirupsen/logrus"
)

// Selecionar generico para listas
func selecionarMedidoresEnel(rows *sql.Rows, err error) ([]model.MedidorEnel, error) {
	if err != nil {
		log.Error("> Erro no exec: ", err)
		return nil, err
	}
	registros := make([]model.MedidorEnel, 0)
	for rows.Next() {
		var registro model.MedidorEnel
		err = rows.Scan(registro.GetParams()...)
		if err != nil {
			log.Error("Erro no scan: ", err)
			return nil, err
		}
		registros = append(registros, registro)
	}
	return registros, nil
}

// InserirUsuario insere um sideufg_medidor_enel no banco de dados
func InserirMedidorEnel(registro model.MedidorEnel) ( *model.MedidorEnel, error ) {
	tx, err := storage.DB.Begin()
	if err != nil {
		log.Error("Erro ao iniciar transação: ", err)
		return nil, err
	}	
	row := tx.QueryRow(query.InserirMedidorEnel, registro.GetParams()[1:]...)
	retorno := new(model.MedidorEnel)
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

// ExcluirUsuario - deleta um sideufg_medidor_enel do banco de dados
func ExcluirMedidorEnel(registro model.MedidorEnel) (err error) {
	_, err = storage.DB.Exec(query.ExcluirMedidorEnel, registro.ID)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// AtualizarTipoMedidor - Atualizar atualiza as informações de um sideufg_medidor_enel na base de dados
func AtualizarMedidorEnel(registro model.MedidorEnel) (err error) {
	_, err = storage.DB.Exec(query.AtualizarMedidorEnel, registro.GetParams()...)
	if err != nil {
		log.Error("Erro no exec: ", err)
	}
	return
}

// ObterTipoMedidor seleciona um sideufg_medidor_enel da base de dados através da PK
func ObterMedidorEnel(registro model.MedidorEnel) (*model.MedidorEnel, error) {
	row := storage.DB.QueryRow(query.ObterMedidorEnel, registro.ID)
	var output model.MedidorEnel
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

// ObterTiposMedidor seleciona todos os sideufg_medidor_enels da base de dados
func ObterMedidoresEnel() ([]model.MedidorEnel, error) {
	registros, err := selecionarMedidoresEnel(storage.DB.Query(query.ObterMedidoresEnel))
	return registros, err
}
	