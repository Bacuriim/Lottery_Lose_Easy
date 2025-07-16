package models

import (
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

func Salvar(c Cliente) string {
	db, _ := database.GetDbSession()
	_, exeError := db.Exec("INSERT INTO cliente (id, nome, cpf, sexo, idade, endereco, conta) VALUES (?, ?, ?, ?, ?, ?, ?)", c.Id, c.Nome, c.Cpf, c.Sexo, c.Idade, c.Endereco, c.Conta)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return "Novo Cliente Salvo!"
}
