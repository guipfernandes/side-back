package query

// CamposEndereco são campos da tabela
const CamposEndereco = `id,
cep,
logradouro,
numero,
complemento,
bairro,
cidade,
estado`

// ObterSideufgEndereco consulta de um único registro por chave primária
var ObterEndereco = `
		SELECT ` + CamposEndereco + `
		FROM public.sideufg_endereco 
		WHERE id = $1 `

// ObterEnderecos consulta livre de todos os registros
var ObterEnderecos = `
		SELECT ` + CamposEndereco + `
		FROM public.sideufg_endereco LIMIT 1000`

// InserirEndereco sql de inclusão
var InserirEndereco = `INSERT INTO public.sideufg_endereco (` + CamposEndereco + `) 
		VALUES(newid(), $1, $2, $3, $4, $5, $6, $7) 
		RETURNING ` + CamposEndereco

// AtualizarEndereco sql de atualização
var AtualizarEndereco = `
		UPDATE public.sideufg_endereco SET 
	   	(` + CamposEndereco + `) = ($1, $2, $3, $4, $5, $6, $7, $8)
       	WHERE id = $1`

// ExcluirEndereco sql de deleção por chave primária
var ExcluirEndereco = `
		DELETE FROM public.sideufg_endereco 
       	WHERE id = $1 `
