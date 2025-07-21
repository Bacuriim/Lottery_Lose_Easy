package main

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"lottery-lose-easy/controllers"
	"lottery-lose-easy/models"
)

func main() {
	uuidCLient := uuid.New()
	cliente := models.Cliente{
		Id:       uuidCLient,
		Nome:     "Conrado",
		Cpf:      "111.222.333-00",
		Sexo:     "M",
		Idade:    20,
		Endereco: "Rua Monte Serrat 19",
		Conta:    "1234567890",
	}

	cliente.Salvar(cliente)
	controllers.Init()
}
