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
var ObterMedidor = `
		SELECT med.id,
			med.denominacao,
			med.latitude,
			med.longitude,
			med.data_primeira_leitura,
			med.data_ultima_leitura,
			med.ultima_data_sincronizada,
			med_enel.id,
			med_enel.nome,
			tp_medidor.id,
			tp_medidor.nome,
			tp_medicao.id,
			tp_medicao.denominacao
		FROM public.sideufg_medidor med 
		LEFT JOIN sideufg_medidor_enel med_enel ON med.medidor_enel_id = med_enel.id 
		LEFT JOIN sideufg_tipo_medidor tp_medidor ON med.tipo = tp_medidor.id
		LEFT JOIN sideufg_tipo_medicao tp_medicao ON med.tipo_medicao = tp_medicao.id
		WHERE med.id = $1 `

// ObterMedidores consulta livre de todos os registros
var ObterMedidores = `
		SELECT med.id,
			med.denominacao,
			med.latitude,
			med.longitude,
			med.data_primeira_leitura,
			med.data_ultima_leitura,
			med.ultima_data_sincronizada,
			med_enel.id,
			med_enel.nome,
			tp_medidor.id,
			tp_medidor.nome,
			tp_medicao.id,
			tp_medicao.denominacao
		FROM public.sideufg_medidor med 
		LEFT JOIN sideufg_medidor_enel med_enel ON med.medidor_enel_id = med_enel.id 
		LEFT JOIN sideufg_tipo_medidor tp_medidor ON med.tipo = tp_medidor.id
		LEFT JOIN sideufg_tipo_medicao tp_medicao ON med.tipo_medicao = tp_medicao.id
		ORDER BY med.denominacao ASC`

// AtualizarMedidor sql de atualização
var AtualizarMedidor = `
		UPDATE public.sideufg_medidor SET 
	   	(` + CamposMedidor + `) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
       	WHERE id = $1 `

// AtualizarMedidorLocalizacao sql de atualização de localização
var AtualizarMedidorLocalizacao = `
		UPDATE public.sideufg_medidor SET latitude = $2, longitude = $3
			WHERE id = $1 `
