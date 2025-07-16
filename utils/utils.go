package utils

import (
	"fmt"
	"os"
)

// Função para esvaziar um arquivo JSON
func EsvaziarArquivoJSON(caminho string) error {
	file, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escreve uma lista vazia no arquivo
	_, err = file.WriteString("[]")
	if err != nil {
		return err
	}

	fmt.Printf("Arquivo %s esvaziado com sucesso.\n", caminho)
	return nil
}
