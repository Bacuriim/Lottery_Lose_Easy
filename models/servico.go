package models

import (
	"database/sql"
	"fmt"
	"lottery-lose-easy/database"
	"strconv"

	"fyne.io/fyne/v2/data/binding"
	_ "github.com/lib/pq"
)

type Servico struct {
	Id          int32
	NomeServico string
}

func (*Servico) Salvar(s Servico) string {
	db, _ := database.GetDbSession()
	_, err := db.Exec(
		"INSERT INTO Servico (id, nome_servico) VALUES ($1, $2)",
		s.Id, s.NomeServico,
	)
	if err != nil {
		return "Erro: " + err.Error()
	}
	return "Novo serviço salvo!"
}

func (*Servico) Remover(s Servico) string {
	db, _ := database.GetDbSession()
	_, err := db.Exec("DELETE FROM Servico WHERE id = $1", s.Id)
	if err != nil {
		return "Erro: " + err.Error()
	}
	return "Serviço removido com sucesso!"
}

func (*Servico) Alterar(rowName string, rowValue interface{}, column string, condition interface{}) string {
	db, _ := database.GetDbSession()

	query := fmt.Sprintf("UPDATE Servico SET %s = $1 WHERE %s = $2", rowName, column)
	_, err := db.Exec(query, rowValue, condition)
	if err != nil {
		return "Erro: " + err.Error()
	}
	return "Serviço alterado com sucesso!"
}

func (*Servico) Pesquisar(searchParameter string, value binding.String, isNumber bool) (*Servico, string) {
	db, _ := database.GetDbSession()
	query := `SELECT id, nome_servico FROM Servico WHERE ` + searchParameter + ` = $1`

	var row *sql.Row
	valueString, _ := value.Get()

	switch {
	case isNumber:
		// Sempre converte para inteiro, pois id é SERIAL
		valueToInt, err := strconv.Atoi(valueString)
		if err != nil {
			return nil, "Erro ao converter valor para inteiro: " + err.Error()
		}
		row = db.QueryRow(query, valueToInt)
	default:
		row = db.QueryRow(query, valueString)
	}

	var s Servico
	err := row.Scan(&s.Id, &s.NomeServico)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &s, "Serviço encontrado!"
}

func (*Servico) BuscarTodos() ([]*Servico, string) {
	db, _ := database.GetDbSession()
	query := `
		SELECT id, nome_servico
		FROM Servico`

	var rows *sql.Rows
	var err error
	rows, err = db.Query(query)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	defer rows.Close()

	var servicos []*Servico
	for rows.Next() {
		var s Servico
		err := rows.Scan(&s.Id, &s.NomeServico)
		if err != nil {
			return nil, "Erro: " + err.Error()
		}
		servicos = append(servicos, &s)
	}
	return servicos, "Serviços encontrados!"
}
