package models

import (
	"fmt"
	_ "github.com/lib/pq"
	"lottery-lose-easy/database"
	"time"
)

type Atendimento struct {
	Id              int32
	NumeroTransacao string
	DataHora        time.Time
	ClienteId       string
	FuncionarioId   string
	ServicoId       int32
	MeioPagamento   string
}

func (*Atendimento) Salvar(a Atendimento) string {
	db, _ := database.GetDbSession()
	_, err := db.Exec(
		`INSERT INTO Atendimento 
		(id, numero_transacao, data_hora, cliente_id, funcionario_id, servico_id, meio_pagamento)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		a.Id, a.NumeroTransacao, a.DataHora, a.ClienteId, a.FuncionarioId, a.ServicoId, a.MeioPagamento,
	)
	if err != nil {
		return "Erro: " + err.Error()
	}
	return "Novo atendimento salvo!"
}

func (*Atendimento) Remover(a Atendimento) string {
	db, _ := database.GetDbSession()
	_, err := db.Exec("DELETE FROM Atendimento WHERE id = $1", a.Id)
	if err != nil {
		return "Erro: " + err.Error()
	}
	return "Atendimento removido com sucesso!"
}

func (*Atendimento) Alterar(rowName string, rowValue interface{}, column string, condition interface{}) string {
	db, _ := database.GetDbSession()

	query := fmt.Sprintf("UPDATE Atendimento SET %s = $1 WHERE %s = $2", rowName, column)
	_, err := db.Exec(query, rowValue, condition)
	if err != nil {
		return "Erro: " + err.Error()
	}
	return "Atendimento alterado com sucesso!"
}

func (*Atendimento) PesquisarPorId(id int32) (*Atendimento, string) {
	db, _ := database.GetDbSession()
	row := db.QueryRow("SELECT id, numero_transacao, data_hora, cliente_id, funcionario_id, servico_id, meio_pagamento FROM Atendimento WHERE id = $1", id)

	var a Atendimento
	err := row.Scan(&a.Id, &a.NumeroTransacao, &a.DataHora, &a.ClienteId, &a.FuncionarioId, &a.ServicoId, &a.MeioPagamento)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &a, "Atendimento encontrado!"
}
