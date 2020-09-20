package erro

//Erro - estrutura global de erro da API
type Erro struct {
	Mensagem string `json:"mensagem,omitempty"`
	Codigo   string `json:"codigo,omitempty"`
}

func Get(err error, variaveis ...string) (erro Erro) {
	erro.Mensagem = err.Error()
	if len(variaveis) > 0 {
		erro.Codigo = variaveis[0]
	}
	return erro
}
