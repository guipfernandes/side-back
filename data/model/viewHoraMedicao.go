package model

import (
	"time"
)

// Medicao - Gerado Automaticamente
type ViewHoraMedicao struct {
	IdMedidor       NullInt64   `json:"id_medidor,omitempty"`
	NomeMedidor     NullString  `json:"nome_medidor"`
	DataMedicao     time.Time   `json:"data_medicao"`
	PotenciaAtiva   NullFloat64 `json:"potencia_ativa,omitempty"`
	PotenciaReativa NullFloat64 `json:"potencia_reativa,omitempty"`
}

// GetParams - Funçao para buscar parametros da struct
func (medicao *ViewHoraMedicao) GetParams() (parametrosFinais []interface{}) {
	parametrosFinais = append(parametrosFinais,
		&medicao.IdMedidor,
		&medicao.NomeMedidor,
		&medicao.DataMedicao,
		&medicao.PotenciaAtiva,
		&medicao.PotenciaReativa)
	return
}
