package main

import (
	_ "github.com/lib/pq"
	"lottery-lose-easy/controllers"
	"lottery-lose-easy/models"
)

func main() {
	cliente := models.Cliente{
		Id:       1,
		Nome:     "Conrado",
		Cpf:      "111.222.333-00",
		Sexo:     "M",
		Idade:    20,
		Endereco: "Rua Monte Serrat 19",
		Conta:    "1234567890",
	}

	models.Salvar(cliente)
	controllers.Init()
}
