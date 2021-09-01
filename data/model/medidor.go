package model

// Medidor - Medidor
type Medidor struct {
	ID                     NullInt64    `json:"id"`
	Denominacao            string       `json:"denominacao,omitempty"`
	Tipo                   *TipoMedidor `json:"tipo,omitempty"`
	Latitude               NullFloat64  `json:"latitude,omitempty"`
	Longitude              NullFloat64  `json:"longitude,omitempty"`
	DataPrimeiraLeitura    NullTime     `json:"data_primeira_leitura,omitempty"`
	DataUltimaLeitura      NullTime     `json:"data_ultima_leitura,omitempty"`
	UltimaDataSincronizada NullTime     `json:"ultima_data_sincronizada,omitempty"`
	MedidorEnel            *MedidorEnel `json:"medidor_enel,omitempty"`
	TipoMedicao            *TipoMedicao `json:"tipo_medicao,omitempty"`
}

// Funçao para inicializar parametros da struct
func (medidor *Medidor) newParams() {
	if medidor.Tipo == nil {
		medidor.Tipo = new(TipoMedidor)
	}
	if medidor.MedidorEnel == nil {
		medidor.MedidorEnel = new(MedidorEnel)
	}
	if medidor.TipoMedicao == nil {
		medidor.TipoMedicao = new(TipoMedicao)
	}
}

// GetParams - Funçao para buscar parametros da struct
func (medidor *Medidor) GetParams() (parametrosFinais []interface{}) {
	medidor.newParams()
	parametrosFinais = append(parametrosFinais,
		&medidor.ID,
		&medidor.Denominacao,
		&medidor.Latitude,
		&medidor.Longitude,
		&medidor.DataPrimeiraLeitura,
		&medidor.DataUltimaLeitura,
		&medidor.UltimaDataSincronizada,
		&medidor.MedidorEnel.ID,
		&medidor.MedidorEnel.Nome,
		&medidor.Tipo.ID,
		&medidor.Tipo.Nome,
		&medidor.TipoMedicao.ID,
		&medidor.TipoMedicao.Denominacao)
	return
}
