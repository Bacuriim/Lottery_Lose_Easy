package models

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	_ "github.com/lib/pq"
	"lottery-lose-easy/database"
	"strings"
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
	query := `
		SELECT id, nome_servico
		FROM Servico WHERE ` + searchParameter + ` = $1`

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

	var s Servico
	err := row.Scan(&s.Id, &s.NomeServico)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &s, "Serviço encontrado!"
}
