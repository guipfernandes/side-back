package model

import (
	"time"
)

// Medicao - Gerado Automaticamente
type Medicao struct {
	ID              NullString `json:"id,omitempty"`
	DataMedicao     time.Time      `json:"data_medicao"`
	PotenciaAtiva   NullString `json:"potencia_ativa,omitempty"`
	PotenciaReativa NullString `json:"potencia_reativa,omitempty"`
	Medidor         *Medidor       `json:"medidor,omitempty"`
}

// Funçao para inicializar parametros da struct
func (medicao *Medicao) newParams() {
	if medicao.Medidor == nil {
		medicao.Medidor = new(Medidor)
	}
}


// GetParams - Funçao para buscar parametros da struct
func (medicao *Medicao) GetParams() (parametrosFinais []interface{}) {
	medicao.newParams()
	parametrosFinais = append(parametrosFinais,
		&medicao.ID,
		&medicao.DataMedicao,
		&medicao.PotenciaAtiva,
		&medicao.PotenciaReativa,
		&medicao.Medidor.ID)
	return
}
