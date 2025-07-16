package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	arquivoProjetosJSON = "data/json/projetos.json" // Caminho para JSON
	arquivoProjetosTXT  = "data/txt/projetos.txt"   // Caminho para TXT
)

type Projeto struct {
	ID                     int    `json:"id"`
	Nome                   string `json:"nome"`
	Local                  string `json:"local"`
	DepartamentoID         int    `json:"departamento_id"`
	FuncionariosProjetosID string `json:"funcionarios_projetos_id"` // Alterado para string
}

// Create
func (p *Projeto) Salvar(departamentos []int, funcionariosProjetos []string) string {
	projetos, err := CarregarProjetos()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar projetos: %v\n", err)
	}

	for _, projeto := range projetos {
		if projeto.ID == p.ID {
			return fmt.Sprintf("Erro: já existe um projeto com o ID %d. Ação não realizada.\n", p.ID)
		}
	}

	departamentoValido := false
	for _, departamentoID := range departamentos {
		if p.DepartamentoID == departamentoID {
			departamentoValido = true
			break
		}
	}

	if !departamentoValido {
		return fmt.Sprintf("Erro: DepartamentoID %d não encontrado. Ação não realizada.\n", p.DepartamentoID)
	}

	funcionarioValido := false
	for _, funcionarioID := range funcionariosProjetos {
		if p.FuncionariosProjetosID == funcionarioID {
			funcionarioValido = true
			break
		}
	}

	if !funcionarioValido {
		return fmt.Sprintf("Erro: FuncionarioID %s não encontrado. Ação não realizada.\n", p.FuncionariosProjetosID)
	}

	projetos = append(projetos, *p)
	if err := salvarProjetos(projetos); err != nil {
		return fmt.Sprintf("Erro ao salvar projeto: %v\n", err)
	}

	sincronizarProjetosTxt(projetos)
	return fmt.Sprintf("Projeto com ID %d salvo com sucesso.\n", p.ID)
}

// Update
func (p *Projeto) Atualizar() string {
	projetos, err := CarregarProjetos()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar projetos: %v\n", err)
	}

	for i, projeto := range projetos {
		if projeto.ID == p.ID {
			projetos[i] = *p
			if err := salvarProjetos(projetos); err != nil {
				return fmt.Sprintf("Erro ao atualizar projeto: %v\n", err)
			}
			sincronizarProjetosTxt(projetos)
			return fmt.Sprintf("Projeto com ID %d atualizado com sucesso.\n", p.ID)
		}
	}

	return fmt.Sprintf("Erro: projeto com ID %d não encontrado para atualização. Ação não realizada.\n", p.ID)
}

// Delete
func (p *Projeto) Deletar() string {
	projetos, err := CarregarProjetos()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar projetos: %v\n", err)
	}

	for i, projeto := range projetos {
		if projeto.ID == p.ID {
			projetos = append(projetos[:i], projetos[i+1:]...)
			if err := salvarProjetos(projetos); err != nil {
				return fmt.Sprintf("Erro ao deletar projeto: %v\n", err)
			}
			sincronizarProjetosTxt(projetos)
			return fmt.Sprintf("Projeto com ID %d deletado com sucesso.\n", p.ID)
		}
	}

	return fmt.Sprintf("Erro: projeto com ID %d não encontrado para exclusão. Ação não realizada.\n", p.ID)
}

// Read
func ListarProjetos() {
	projetos, err := CarregarProjetos()
	if err != nil {
		fmt.Printf("Erro ao carregar projetos: %v\n", err)
		return
	}

	fmt.Println("Lista de projetos:")
	for _, p := range projetos {
		fmt.Printf("ID: %d, Nome: %s, Local: %s, DepartamentoID: %d, FuncionariosProjetosID: %s\n",
			p.ID, p.Nome, p.Local, p.DepartamentoID, p.FuncionariosProjetosID)
	}
}

// Funções utilitárias

func CarregarProjetos() ([]Projeto, error) {
	var projetos []Projeto

	file, err := os.Open(arquivoProjetosJSON)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return projetos, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&projetos); err != nil {
		return nil, err
	}

	return projetos, nil
}

func salvarProjetos(projetos []Projeto) error {
	file, err := os.Create(arquivoProjetosJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(projetos)
}

// Sincronizar o estado do JSON com o arquivo TXT
func sincronizarProjetosTxt(projetos []Projeto) {
	file, err := os.Create(arquivoProjetosTXT)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo TXT: %v\n", err)
		return
	}
	defer file.Close()

	for _, p := range projetos {
		linha := fmt.Sprintf("ID: %d, Nome: %s, Local: %s, DepartamentoID: %d, FuncionariosProjetosID: %s\n",
			p.ID, p.Nome, p.Local, p.DepartamentoID, p.FuncionariosProjetosID)
		_, err := file.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo TXT: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo TXT sincronizado com sucesso.")
}
