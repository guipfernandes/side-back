package query

// Consultas do CRUD

// CamposMedidor são campos da tabela
const CamposMedidor = `id,
denominacao,
tipo,
latitude,
longitude,
data_primeira_leitura,
data_ultima_leitura,
ultima_data_sincronizada,
medidor_enel_id`

// ObterTipoMedidor consulta de um único registro por chave primária
var	ObterMedidor = `
		SELECT ` + CamposMedidor + `
		FROM public.sideufg_medidor 
		WHERE id = $1 `

// ObterMedidores consulta livre de todos os registros
var ObterMedidores = `
		SELECT ` + CamposMedidor + `
		FROM public.sideufg_medidor LIMIT 1000`

// AtualizarMedidor sql de atualização
var AtualizarMedidor = `
		UPDATE public.sideufg_medidor SET 
	   	(` + CamposMedidor + `) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
       	WHERE id = $1 `

// Consultas do Específicas
	