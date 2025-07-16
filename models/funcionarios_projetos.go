package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	arquivoFuncionariosProjetosJSON = "data/json/funcionarios_projetos.json"
	arquivoFuncionariosProjetosTXT  = "data/txt/funcionarios_projetos.txt"
)

type FuncionarioProjeto struct {
	ID            string `json:"id"`             // Alterado para string
	FuncionarioID string `json:"funcionario_id"` // Alterado para string
	ProjetoID     int    `json:"projeto_id"`
}

// Create
func (fp *FuncionarioProjeto) Salvar() {
	relacoes, err := carregarFuncionariosProjetos()
	if err != nil {
		fmt.Printf("Erro ao carregar relações funcionário-projeto: %v\n", err)
		return
	}

	for _, relacao := range relacoes {
		if relacao.ID == fp.ID {
			fmt.Printf("Erro: já existe uma relação com o ID %s. Ação não realizada.\n", fp.ID)
			return
		}
	}

	relacoes = append(relacoes, *fp)
	if err := salvarFuncionariosProjetos(relacoes); err != nil {
		fmt.Printf("Erro ao salvar relação funcionário-projeto: %v\n", err)
		return
	}

	fmt.Printf("Relação funcionário-projeto com ID %s salva com sucesso.\n", fp.ID)
	sincronizarFuncionariosProjetosTxt(relacoes)
}

// Update
func (fp *FuncionarioProjeto) Atualizar() {
	relacoes, err := carregarFuncionariosProjetos()
	if err != nil {
		fmt.Printf("Erro ao carregar relações funcionário-projeto: %v\n", err)
		return
	}

	for i, relacao := range relacoes {
		if relacao.ID == fp.ID {
			relacoes[i] = *fp
			if err := salvarFuncionariosProjetos(relacoes); err != nil {
				fmt.Printf("Erro ao atualizar relação funcionário-projeto: %v\n", err)
				return
			}
			fmt.Printf("Relação funcionário-projeto com ID %s atualizada com sucesso.\n", fp.ID)
			sincronizarFuncionariosProjetosTxt(relacoes)
			return
		}
	}

	fmt.Printf("Erro: relação funcionário-projeto com ID %s não encontrada para atualização. Ação não realizada.\n", fp.ID)
}

// Delete
func (fp *FuncionarioProjeto) Deletar() {
	relacoes, err := carregarFuncionariosProjetos()
	if err != nil {
		fmt.Printf("Erro ao carregar relações funcionário-projeto: %v\n", err)
		return
	}

	for i, relacao := range relacoes {
		if relacao.ID == fp.ID {
			relacoes = append(relacoes[:i], relacoes[i+1:]...)
			if err := salvarFuncionariosProjetos(relacoes); err != nil {
				fmt.Printf("Erro ao deletar relação funcionário-projeto: %v\n", err)
				return
			}
			fmt.Printf("Relação funcionário-projeto com ID %s deletada com sucesso.\n", fp.ID)
			sincronizarFuncionariosProjetosTxt(relacoes)
			return
		}
	}

	fmt.Printf("Erro: relação funcionário-projeto com ID %s não encontrada para exclusão. Ação não realizada.\n", fp.ID)
}

// Read
func ListarFuncionariosProjetos() {
	relacoes, err := carregarFuncionariosProjetos()
	if err != nil {
		fmt.Printf("Erro ao carregar relações funcionário-projeto: %v\n", err)
		return
	}

	fmt.Println("Lista de relações funcionário-projeto:")
	for _, fp := range relacoes {
		fmt.Printf("ID: %s, FuncionarioID: %s, ProjetoID: %d\n", fp.ID, fp.FuncionarioID, fp.ProjetoID)
	}
}

// Funções utilitárias

func carregarFuncionariosProjetos() ([]FuncionarioProjeto, error) {
	var relacoes []FuncionarioProjeto

	file, err := os.Open(arquivoFuncionariosProjetosJSON)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return relacoes, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&relacoes); err != nil {
		return nil, err
	}

	return relacoes, nil
}

func salvarFuncionariosProjetos(relacoes []FuncionarioProjeto) error {
	file, err := os.Create(arquivoFuncionariosProjetosJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(relacoes)
}

// Sincronizar o estado do JSON com o arquivo TXT
func sincronizarFuncionariosProjetosTxt(relacoes []FuncionarioProjeto) {
	file, err := os.Create(arquivoFuncionariosProjetosTXT)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo TXT: %v\n", err)
		return
	}
	defer file.Close()

	for _, fp := range relacoes {
		linha := fmt.Sprintf("ID: %s, FuncionarioID: %s, ProjetoID: %d\n", fp.ID, fp.FuncionarioID, fp.ProjetoID)
		_, err := file.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo TXT: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo TXT sincronizado com sucesso.")
}
