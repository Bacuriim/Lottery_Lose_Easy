package models

import (
	"database/sql"
	"fmt"
	"lottery-lose-easy/database"
	"strings"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Atendimento struct {
	Id              int32
	NumeroTransacao string
	DataHora        time.Time
	ClienteId       uuid.UUID
	FuncionarioId   uuid.UUID
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

func (*Atendimento) Pesquisar(searchParameter string, value binding.String, isNumber bool) (*Atendimento, string) {
	db, _ := database.GetDbSession()
	query := `
		SELECT id, numero_transacao, data_hora, cliente_id, funcionario_id, servico_id, meio_pagamento
		FROM Atendimento WHERE ` + searchParameter + ` = $1`

	var row *sql.Row
	valueString, _ := value.Get()

	if isNumber {
		if strings.Contains(valueString, ".") {
			valueToFloat := binding.StringToFloat(value)
			row = db.QueryRow(query, valueToFloat)
		} else {
			valueToInt := binding.StringToInt(value)
			row = db.QueryRow(query, valueToInt)
		}
	} else {
		row = db.QueryRow(query, valueString)
	}

	var a Atendimento
	err := row.Scan(&a.Id, &a.NumeroTransacao, &a.DataHora, &a.ClienteId, &a.FuncionarioId, &a.ServicoId, &a.MeioPagamento)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &a, "Atendimento encontrado!"
}

func (*Atendimento) BuscarTodos() ([]*Atendimento, string) {
	db, _ := database.GetDbSession()
	query := `
		SELECT id, numero_transacao, data_hora, cliente_id, funcionario_id, servico_id, meio_pagamento
		FROM Atendimento`

	rows, err := db.Query(query)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	defer rows.Close()

	var atendimentos []*Atendimento
	for rows.Next() {
		var a Atendimento
		err := rows.Scan(&a.Id, &a.NumeroTransacao, &a.DataHora, &a.ClienteId, &a.FuncionarioId, &a.ServicoId, &a.MeioPagamento)
		if err != nil {
			return nil, "Erro: " + err.Error()
		}
		atendimentos = append(atendimentos, &a)
	}

	return atendimentos, "Atendimentos encontrados!"
}
