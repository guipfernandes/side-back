package model

type MedidorEnel struct {
	ID               NullString `json:"id"`
	UC               string     `json:"uc"`
	Nome             NullString `json:"nome"`
	NumMedidor       string     `json:"num_medidor"`
	Demanda          NullInt64  `json:"demanda"`
	DemandaPonta     NullInt64  `json:"demanda_ponta"`
	DemandaForaPonta NullInt64  `json:"demanda_fora_ponta"`
	Endereco         *Endereco  `json:"endereco"`
}

// Funçao para inicializar parametros da struct
func (medidorEnel *MedidorEnel) newParams() {
	if medidorEnel.Endereco == nil {
		medidorEnel.Endereco = new(Endereco)
	}
}

// GetParams - Funçao para buscar parametros da struct
func (medidorEnel *MedidorEnel) GetParams() (parametrosFinais []interface{}) {
	medidorEnel.newParams()
	parametrosFinais = append(parametrosFinais,
		&medidorEnel.ID,
		&medidorEnel.UC,
		&medidorEnel.Nome,
		&medidorEnel.NumMedidor,
		&medidorEnel.Demanda,
		&medidorEnel.DemandaPonta,
		&medidorEnel.DemandaForaPonta,
		&medidorEnel.Endereco.ID)
	return
}
