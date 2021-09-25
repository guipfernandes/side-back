package model

// MedidorLocalizacao - MedidorLocalizacao
type MedidorLocalizacao struct {
	ID        NullInt64   `json:"id"`
	Latitude  NullFloat64 `json:"latitude,omitempty"`
	Longitude NullFloat64 `json:"longitude,omitempty"`
}

// GetParams - Funçao para buscar parametros da struct
func (medidor *MedidorLocalizacao) GetParams() (parametrosFinais []interface{}) {
	parametrosFinais = append(parametrosFinais,
		&medidor.ID,
		&medidor.Latitude,
		&medidor.Longitude)
	return
}
