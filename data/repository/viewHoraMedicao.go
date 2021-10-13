package repository

import (
	"database/sql"
	"side/core/storage"
	"side/data/model"
	"side/data/query"
	"time"

	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// Selecionar generico para listas
func SelecionarViewHoraMedicoes(rows *sql.Rows, err error) ([]model.ViewHoraMedicao, error) {
	if err != nil {
		log.Error("> Erro no exec: ", err)
		return nil, err
	}
	registros := make([]model.ViewHoraMedicao, 0)
	for rows.Next() {
		var registro model.ViewHoraMedicao
		err = rows.Scan(registro.GetParams()...)
		if err != nil {
			log.Error("Erro no scan: ", err)
			return nil, err
		}
		registros = append(registros, registro)
	}
	return registros, nil
}

// ObterHoraMedicoesByMedidorEDataMedicao seleciona vw_hora_medicao a partir dos medidores e data de medicao da base de dados
func ObterHoraMedicoesByMedidorEDataMedicao(idMedidores []int, dataMedicaoInicio, dataMedicaoFim time.Time) ([]model.ViewHoraMedicao, error) {
	layout := "2006-01-02 15:04:05"
	registros, err := SelecionarViewHoraMedicoes(
		storage.DB.Query(
			query.ObterViewHoraMedicaoByMedidorEDataMedicao,
			pq.Array(idMedidores),
			dataMedicaoInicio.Format(layout),
			dataMedicaoFim.Format(layout)))
	return registros, err
}
