package model

// SideufgEndereco - Gerado Automaticamente
type Endereco struct {
	ID          NullString `json:"id,omitempty"`
	Cep         string     `json:"cep,omitempty"`
	Logradouro  string     `json:"logradouro,omitempty"`
	Numero      int        `json:"numero,omitempty"`
	Complemento NullString `json:"complemento,omitempty"`
	Bairro      NullString `json:"bairro,omitempty"`
	Cidade      string     `json:"cidade,omitempty"`
	Estado      string     `json:"estado,omitempty"`
}

// GetParams - Funçao para buscar parametros da struct
func (endereco *Endereco) GetParams() (parametrosFinais []interface{}) {
	parametrosFinais = append(parametrosFinais,
		&endereco.ID,
		&endereco.Cep,
		&endereco.Logradouro,
		&endereco.Numero,
		&endereco.Complemento,
		&endereco.Bairro,
		&endereco.Cidade,
		&endereco.Estado)
	return
}
