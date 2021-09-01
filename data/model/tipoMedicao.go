package model

type TipoMedicao struct {
	ID          NullInt64  `json:"id"`
	Denominacao NullString `json:"denominacao"`
}

// GetParams - Funçao para buscar parametros da struct
func (tipoMedicao *TipoMedicao) GetParams() (parametrosFinais []interface{}) {
	parametrosFinais = append(parametrosFinais,
		&tipoMedicao.ID,
		&tipoMedicao.Denominacao)
	return
}
