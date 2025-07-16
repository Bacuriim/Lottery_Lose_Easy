package models

import (
	"fmt"
	_ "github.com/lib/pq"
	"lottery-lose-easy/database"
)

type Cliente struct {
	Id       int
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

func (*Cliente) PesquisarPorId(id string) (*Cliente, string) {
	db, _ := database.GetDbSession()
	result := db.QueryRow("SELECT id, nome, cpf, sexo, idade, endereco, conta FROM Cliente WHERE id = $1", id)
	var c Cliente
	exeError := result.Scan(&c.Id, &c.Nome, &c.Cpf, &c.Sexo, &c.Idade, &c.Endereco, &c.Conta)
	if exeError != nil {
		return nil, "Erro: " + exeError.Error()
	}
	return &c, "Cliente encontrado!"
}
