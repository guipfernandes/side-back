package query

// CamposMedicao são campos da tabela
const CamposMedicao = `id,
data_medicao,
potencia_ativa,
potencia_reativa,
medidor_id`

// ObterMedicao consulta de um único registro por chave primária
var ObterMedicao = `
		SELECT ` + CamposMedicao + `
		FROM public.sideufg_medicao 
		WHERE id = $1 `

// ObterMedicoes consulta livre de todos os registros
var ObterMedicoes = `
		SELECT ` + CamposMedicao + `
		FROM public.sideufg_medicao LIMIT 1000`
