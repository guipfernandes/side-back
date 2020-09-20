package query

const CamposTipoMedidor = `id,
nome`

// ObterTipoMedidor consulta de um único registro por chave primária
var	ObterTipoMedidor = `
		SELECT ` + CamposTipoMedidor + `
		FROM public.sideufg_tipo_medidor 
		WHERE id = $1 `

// ObterTiposMedidor consulta livre de todos os registros
var ObterTiposMedidor = `
		SELECT ` + CamposTipoMedidor + `
		FROM public.sideufg_tipo_medidor LIMIT 1000`

// InserirTipoMedidor sql de inclusão
var InserirTipoMedidor = `INSERT INTO public.sideufg_tipo_medidor (` + CamposTipoMedidor + `) 
		VALUES(NEXTVAL('sideufg_tipo_medidor_id_seq'), $1) 
		RETURNING ` + CamposTipoMedidor

// AtualizarTipoMedidor sql de atualização
var AtualizarTipoMedidor = `
		UPDATE public.sideufg_tipo_medidor SET 
	   	(` + CamposTipoMedidor + `) = ($1, $2)
       	WHERE id = $1 `

// ExcluirTipoMedidor sql de deleção por chave primária
var ExcluirTipoMedidor = `
		DELETE FROM public.sideufg_tipo_medidor 
       	WHERE id = $1 `
