package models

import (
	"fmt"
	_ "github.com/lib/pq"
	"lottery-lose-easy/database"
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

func (*Servico) PesquisarPorId(id int32) (*Servico, string) {
	db, _ := database.GetDbSession()
	row := db.QueryRow("SELECT id, nome_servico FROM Servico WHERE id = $1", id)

	var s Servico
	err := row.Scan(&s.Id, &s.NomeServico)
	if err != nil {
		return nil, "Erro: " + err.Error()
	}
	return &s, "Serviço encontrado!"
}
