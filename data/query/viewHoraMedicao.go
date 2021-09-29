package query

// CamposViewHoraMedicao são campos da view
const CamposViewHoraMedicao = `
id_medidor,
nome_medidor,
data_medicao,
potencia_ativa,
potencia_reativa`

// ObterViewHoraMedicaoByMedidorEDataMedicao view que consulta registros por medidor e data medição
var ObterViewHoraMedicaoByMedidorEDataMedicao = `
		SELECT ` + CamposViewHoraMedicao + `
		FROM public.vw_hora_medicao
		WHERE nome_medidor = $1 and data_medicao between $2 and $3`
