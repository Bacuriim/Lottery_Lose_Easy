package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	arquivoDepartamentosJSON = "data/json/departamentos.json"
	arquivoDepartamentosTXT  = "data/txt/departamentos.txt"
)

type Departamento struct {
	ID      int    `json:"id"`
	Nome    string `json:"nome"`
	ChefeID int    `json:"chefe_id"` // Novo campo para associar o chefe ao departamento
}

// Create
func (d *Departamento) Salvar() string {
	departamentos, err := CarregarDepartamentos()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar departamentos: %v\n", err)
	}

	for _, departamento := range departamentos {
		if departamento.ID == d.ID {
			return fmt.Sprintf("Erro: já existe um departamento com o ID %d. Ação não realizada.\n", d.ID)
		}
	}

	departamentos = append(departamentos, *d)
	if err := salvarDepartamentos(departamentos); err != nil {
		return fmt.Sprintf("Erro ao salvar departamento: %v\n", err)
	}

	sincronizarDepartamentosTxt(departamentos)
	return fmt.Sprintf("Departamento com ID %d salvo com sucesso.\n", d.ID)
}

// Update
func (d *Departamento) Atualizar() string {
	departamentos, err := CarregarDepartamentos()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar departamentos: %v\n", err)
	}

	for i, departamento := range departamentos {
		if departamento.ID == d.ID {
			departamentos[i] = *d
			if err := salvarDepartamentos(departamentos); err != nil {
				return fmt.Sprintf("Erro ao atualizar departamento: %v\n", err)
			}
			sincronizarDepartamentosTxt(departamentos)
			return fmt.Sprintf("Departamento com ID %d atualizado com sucesso.\n", d.ID)
		}
	}

	return fmt.Sprintf("Erro: departamento com ID %d não encontrado para atualização. Ação não realizada.\n", d.ID)
}

// Delete
func (d *Departamento) Deletar() string {
	departamentos, err := CarregarDepartamentos()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar departamentos: %v\n", err)
	}

	for i, departamento := range departamentos {
		if departamento.ID == d.ID {
			departamentos = append(departamentos[:i], departamentos[i+1:]...)
			if err := salvarDepartamentos(departamentos); err != nil {
				return fmt.Sprintf("Erro ao deletar departamento: %v\n", err)
			}
			sincronizarDepartamentosTxt(departamentos)
			return fmt.Sprintf("Departamento com ID %d deletado com sucesso.\n", d.ID)
		}
	}

	return fmt.Sprintf("Erro: departamento com ID %d não encontrado para exclusão. Ação não realizada.\n", d.ID)
}

// Funções utilitárias

func CarregarDepartamentos() ([]Departamento, error) {
	var departamentos []Departamento

	file, err := os.Open(arquivoDepartamentosJSON)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return departamentos, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&departamentos); err != nil {
		return nil, err
	}

	return departamentos, nil
}

func salvarDepartamentos(departamentos []Departamento) error {
	file, err := os.Create(arquivoDepartamentosJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(departamentos)
}

// Sincronizar o estado do JSON com o arquivo TXT
func sincronizarDepartamentosTxt(departamentos []Departamento) {
	file, err := os.Create(arquivoDepartamentosTXT)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo TXT: %v\n", err)
		return
	}
	defer file.Close()

	for _, d := range departamentos {
		linha := fmt.Sprintf("ID: %d, Nome: %s, ChefeID: %d\n", d.ID, d.Nome, d.ChefeID)
		_, err := file.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo TXT: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo TXT sincronizado com sucesso.")
}

// Read
func ListarDepartamentos() []Departamento {
	departamentos, err := CarregarDepartamentos()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionários: %v\n", err)
		return nil
	}

	fmt.Println("Lista de funcionários:")
	for _, d := range departamentos {
		fmt.Printf("ID: %d, Nome: %s, ChefeID: %d\n",
			d.ID, d.Nome, d.ChefeID)
	}
	listaDepartamentos := []Departamento{}
	listaDepartamentos = append(listaDepartamentos, departamentos...)
	return listaDepartamentos
}

func GetDepartamentosIDs() []int {
	departamentos, err := CarregarDepartamentos()
	if err != nil {
		fmt.Printf("Erro ao carregar departamentos: %v\n", err)
		return nil
	}

	var ids []int
	for _, d := range departamentos {
		ids = append(ids, d.ID)
	}
	return ids
}
