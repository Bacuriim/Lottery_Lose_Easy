package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	arquivoChefesDepartamentoJSON = "data/json/chefes_departamento.json"
	arquivoChefesDepartamentoTXT  = "data/txt/chefes_departamento.txt"
)

type ChefeDepartamento struct {
	ID            int    `json:"id"`
	FuncionarioID string `json:"funcionario_id"` // Alterado para string
}

// Create
func (c *ChefeDepartamento) Salvar(funcionarios []string) string {
	chefes, err := CarregarChefesDepartamento()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar chefes de departamento: %v\n", err)
	}

	for _, chefe := range chefes {
		if c.ID == 0 || c.FuncionarioID == "" {
			return "Erro: todos os campos devem ser preenchidos. Ação não realizada.\n"
		}

		if chefe.ID == c.ID {
			return fmt.Sprintf("Erro: já existe um chefe de departamento com o ID %d. Ação não realizada.\n", c.ID)
		}
	}

	funcionarioValido := false
	for _, funcionario := range funcionarios {
		if c.FuncionarioID == funcionario {
			funcionarioValido = true
			break
		}
	}

	if !funcionarioValido {
		return fmt.Sprintf("Erro: FuncionarioID %s não encontrado. Ação não realizada.\n", c.FuncionarioID)
	}

	chefes = append(chefes, *c)
	if err := salvarChefesDepartamento(chefes); err != nil {
		return fmt.Sprintf("Erro ao salvar chefe de departamento: %v\n", err)
	}

	sincronizarChefesDepartamentoTxt(chefes)
	return fmt.Sprintf("Chefe de departamento com ID %d salvo com sucesso.\n", c.ID)
}

// Update
func (c *ChefeDepartamento) Atualizar() string {
	chefes, err := CarregarChefesDepartamento()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar chefes de departamento: %v\n", err)
	}

	for i, chefe := range chefes {
		if chefe.ID == c.ID {
			chefes[i] = *c
			if err := salvarChefesDepartamento(chefes); err != nil {
				return fmt.Sprintf("Erro ao atualizar chefe de departamento: %v\n", err)
			}
			sincronizarChefesDepartamentoTxt(chefes)
			return fmt.Sprintf("Chefe de departamento com ID %d atualizado com sucesso.\n", c.ID)
		}
	}

	return fmt.Sprintf("Erro: chefe de departamento com ID %d não encontrado para atualização. Ação não realizada.\n", c.ID)
}

// Delete
func (c *ChefeDepartamento) Deletar() string {
	chefes, err := CarregarChefesDepartamento()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar chefes de departamento: %v\n", err)
	}

	for i, chefe := range chefes {
		if chefe.ID == c.ID {
			chefes = append(chefes[:i], chefes[i+1:]...)
			if err := salvarChefesDepartamento(chefes); err != nil {
				return fmt.Sprintf("Erro ao deletar chefe de departamento: %v\n", err)
			}
			sincronizarChefesDepartamentoTxt(chefes)
			return fmt.Sprintf("Chefe de departamento com ID %d deletado com sucesso.\n", c.ID)
		}
	}

	return fmt.Sprintf("Erro: chefe de departamento com ID %d não encontrado para exclusão. Ação não realizada.\n", c.ID)
}

// Read
func ListarChefesDepartamento() {
	chefes, err := CarregarChefesDepartamento()
	if err != nil {
		fmt.Printf("Erro ao carregar chefes de departamento: %v\n", err)
		return
	}

	fmt.Println("Lista de chefes de departamento:")
	for _, c := range chefes {
		fmt.Printf("ID: %d, FuncionarioID: %s\n", c.ID, c.FuncionarioID)
	}
}

// Funções utilitárias

func CarregarChefesDepartamento() ([]ChefeDepartamento, error) {
	var chefes []ChefeDepartamento

	file, err := os.Open(arquivoChefesDepartamentoJSON)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return chefes, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&chefes); err != nil {
		return nil, err
	}

	return chefes, nil
}

func salvarChefesDepartamento(chefes []ChefeDepartamento) error {
	file, err := os.Create(arquivoChefesDepartamentoJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(chefes)
}

func sincronizarChefesDepartamentoTxt(chefes []ChefeDepartamento) {
	file, err := os.Create(arquivoChefesDepartamentoTXT)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo TXT: %v\n", err)
		return
	}
	defer file.Close()

	for _, c := range chefes {
		linha := fmt.Sprintf("ID: %d, FuncionarioID: %s\n", c.ID, c.FuncionarioID)
		_, err := file.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo TXT: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo TXT sincronizado com sucesso.")
}
