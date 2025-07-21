package models

import "lottery-lose-easy/database"

type ServicosUtilizados struct {
	Nome  string
	Total int
}

type PagamentosUtilizados struct {
	Forma string
	Total int
}

func ListarPagamentosMaisUtilizados() ([]*PagamentosUtilizados, string) {
	db, err := database.GetDbSession()
	if err != nil {
		return nil, "Erro ao conectar ao banco: " + err.Error()
	}

	query := `
		SELECT meio_pagamento, COUNT(*) as total
		FROM Atendimento
		WHERE meio_pagamento IS NOT NULL
		GROUP BY meio_pagamento
		ORDER BY total DESC;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, "Erro ao executar consulta: " + err.Error()
	}
	defer rows.Close()

	var pagamentos []*PagamentosUtilizados
	for rows.Next() {
		var p PagamentosUtilizados
		if err := rows.Scan(&p.Forma, &p.Total); err != nil {
			return nil, "Erro ao ler resultados: " + err.Error()
		}
		pagamentos = append(pagamentos, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, "Erro após leitura dos resultados: " + err.Error()
	}

	return pagamentos, "Pagamentos mais utilizados encontrados!"
}

func ListarServicosMaisUtilizados() ([]*ServicosUtilizados, string) {
	db, err := database.GetDbSession()
	if err != nil {
		return nil, "Erro ao conectar ao banco: " + err.Error()
	}

	query := `
		SELECT s.nome_servico, COUNT(a.servico_id) AS total
		FROM Atendimento a
		INNER JOIN Servico s ON a.servico_id = s.id
		GROUP BY s.nome_servico
		ORDER BY total DESC;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, "Erro ao executar consulta: " + err.Error()
	}
	defer rows.Close()

	var servicos []*ServicosUtilizados
	for rows.Next() {
		var s ServicosUtilizados
		if err := rows.Scan(&s.Nome, &s.Total); err != nil {
			return nil, "Erro ao ler resultados: " + err.Error()
		}
		servicos = append(servicos, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, "Erro após leitura dos resultados: " + err.Error()
	}

	return servicos, "Serviços mais utilizados encontrados!"
}
