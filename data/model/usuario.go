package model

// Usuario - Usuario do Sistema
type Usuario struct {
	ID       NullString `json:"id"`
	Login    string     `json:"login"`
	Senha    string     `json:"senha"`
	Nome     string     `json:"nome"`
	Email    string     `json:"email"`
	Situacao int64      `json:"ativo" `
}

// GetParams - Funçao para buscar parametros da struct
func (usuario *Usuario) GetParams() (parametrosFinais []interface{}) {
	parametrosFinais = append(parametrosFinais,
		&usuario.ID,
		&usuario.Login,
		&usuario.Senha,
		&usuario.Nome,
		&usuario.Email,
		&usuario.Situacao)
	return
}
