package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connStr = "user=postgres password=admin dbname=LotteryLoseEasy sslmode=disable"
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Erro ao abrir conexão:", err)
		db = nil
		return
	}

	// Testa conexão
	if err = db.Ping(); err != nil {
		fmt.Println("Erro ao conectar com o banco:", err)
		db = nil
	}
}

func GetDbSession() (*sql.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("conexão com o banco de dados não foi estabelecida")
	}
	return db, nil
}
