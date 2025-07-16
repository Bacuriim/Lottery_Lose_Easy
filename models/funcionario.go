package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shopspring/decimal"
	"lottery-lose-easy/database"
	"math"
)

type Funcionario struct {
	Id                  int
	Nome                string
	NumeroIdentificacao string
	Salario             decimal.Decimal
	Cpf                 string
	Sexo                string
	Idade               int
	HorarioTrabalho     string
}

func (*Funcionario) Salvar(f Funcionario) string {
	db, _ := database.GetDbSession()
	_, exeError := db.Exec(`
		INSERT INTO Funcionario (id, nome, numero_identificacao, salario, cpf, sexo, idade, horario_trabalho)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		f.Id, f.Nome, f.NumeroIdentificacao, f.Cpf, f.Sexo, f.Idade, f.HorarioTrabalho,
	)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return "Novo funcionário salvo!"
}

func (*Funcionario) Remover(f Funcionario) string {
	db, _ := database.GetDbSession()
	_, exeError := db.Exec("DELETE FROM Funcionario WHERE id = $1", f.Id)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return "Funcionário removido com sucesso!"
}

func (*Funcionario) Alterar(rowName string, rowValue interface{}, column string, condition interface{}) string {
	db, _ := database.GetDbSession()
	query := fmt.Sprintf("UPDATE Funcionario SET %s = $1 WHERE %s = $2", rowName, column)
	_, exeError := db.Exec(query, rowValue, condition)
	if exeError != nil {
		return "Erro: " + exeError.Error()
	}
	return "Funcionário alterado com sucesso!"
}

func (*Funcionario) PesquisarPorId(id string) (*Funcionario, string) {
	db, _ := database.GetDbSession()
	row := db.QueryRow(`
		SELECT id, nome, numero_identificacao, cpf, sexo, idade, horario_trabalho
		FROM Funcionario WHERE id = $1`, id)

	var f Funcionario
	err := row.Scan(&f.Id, &f.Nome, &f.NumeroIdentificacao, &f.Cpf, &f.Sexo, &f.Idade, &f.HorarioTrabalho)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &f, "Funcionário encontrado!"
}
