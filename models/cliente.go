package models

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
	"lottery-lose-easy/database"
	"strings"
)

type Cliente struct {
	Id       uuid.UUID
	Nome     string
	Cpf      string
	Sexo     string
	Idade    int
	Endereco string
	Conta    string
}

func (*Cliente) Salvar(c Cliente) string {
	db, _ := database.GetDbSession()
	_, exeError := db.Exec("INSERT INTO Cliente (id, nome, cpf, sexo, idade, endereco, conta) VALUES ($1, $2, $3, $4, $5, $6, $7)", c.Id, c.Nome, c.Cpf, c.Sexo, c.Idade, c.Endereco, c.Conta)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return fmt.Sprintf("Novo cliente salvo!")
}

func (*Cliente) Remover(c Cliente) string {
	db, _ := database.GetDbSession()
	_, exeError := db.Exec("DELETE FROM Cliente WHERE id = $1", c.Id)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return "Novo cliente salvo!"
}

func (*Cliente) Alterar(rowName string, rowValue interface{}, column string, condition interface{}) string {
	db, _ := database.GetDbSession()
	query := fmt.Sprintf("UPDATE Cliente SET %s = $1 WHERE %s = $2", rowName, column)
	_, exeError := db.Exec(query, rowValue, condition)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return "Novo cliente salvo!"
}

func (*Cliente) Pesquisar(searchParameter string, value binding.String, isNumber bool) (*Cliente, string) {
	db, _ := database.GetDbSession()
	query := `
		SELECT id, nome, cpf, sexo, idade, endereco, conta
		FROM Cliente WHERE ` + searchParameter + ` = $1`

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

	var c Cliente
	err := row.Scan(&c.Id, &c.Nome, &c.Cpf, &c.Sexo, &c.Idade, &c.Endereco, &c.Conta)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &c, "Cliente encontrado!"
}
