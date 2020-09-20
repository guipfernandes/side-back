package query

// CamposMedidorEnel são campos da tabela
const CamposMedidorEnel = `id,
uc,
nome,
num_medidor,
demanda,
demanda_ponta,
demanda_fora_ponta,
endereco_id`

// ObterMedidorEnel consulta de um único registro por chave primária
var ObterMedidorEnel = `
		SELECT ` + CamposMedidorEnel + `
		FROM public.sideufg_medidor_enel 
		WHERE id = $1 `

// ObterMedidoresEnel consulta livre de todos os registros
var ObterMedidoresEnel = `
		SELECT ` + CamposMedidorEnel + `
		FROM public.sideufg_medidor_enel LIMIT 1000`

// InserirMedidorEnel sql de inclusão
var InserirMedidorEnel = `INSERT INTO public.sideufg_medidor_enel (` + CamposMedidorEnel + `) 
		VALUES(newid(), $1, $2, $3, $4, $5, $6, $7) 
		RETURNING ` + CamposMedidorEnel

// AtualizarMedidorEnel sql de atualização
var AtualizarMedidorEnel = `
		UPDATE public.sideufg_medidor_enel SET 
	   	(` + CamposMedidorEnel + `) = ($1, $2, $3, $4, $5, $6, $7, $8)
       	WHERE id = $1 `

// ExcluirMedidorEnel sql de deleção por chave primária
var ExcluirMedidorEnel = `
		DELETE FROM public.sideufg_medidor_enel 
       	WHERE id = $1 `

