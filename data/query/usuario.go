package query

const CamposUsuario = `id,
login,
senha,
nome,
email,
situacao`

// ObterUsuario consulta livre de todos os registros
var ObterUsuarioAtivoPorLogin = `
		SELECT ` + CamposUsuario + `
		FROM public.sideufg_usuario 
		WHERE login = $1 AND situacao = 1`

// InserirUsuario sql de inclusão
var InserirUsuario = `INSERT INTO public.sideufg_usuario (` + CamposUsuario + `) 
		VALUES(newid(), $1, $2, $3, $4, $5) 
		RETURNING ` + CamposUsuario

// AtualizarUsuario sql de atualização
var AtualizarUsuario = `
		UPDATE public.sideufg_usuario SET 
	   	(` + CamposUsuario + `) = ($1, $2, $3, $4, $5, $6)
       	WHERE id = $1 `

// ExcluirUsuario sql de deleção por chave primária
var ExcluirUsuario = `
		DELETE FROM public.sideufg_usuario 
       	WHERE id = $1 `
