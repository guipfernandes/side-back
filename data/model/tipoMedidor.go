package model

type TipoMedidor struct {
	ID   NullInt64  `json:"id"`
	Nome NullString `json:"nome"`
}

// GetParams - Funçao para buscar parametros da struct
func (tipoMedidor *TipoMedidor) GetParams() (parametrosFinais []interface{}) {
	parametrosFinais = append(parametrosFinais,
		&tipoMedidor.ID,
		&tipoMedidor.Nome)
	return
}