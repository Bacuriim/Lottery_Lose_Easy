package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	arquivoFuncionariosJSON = "data/json/funcionarios.json"
	arquivoFuncionariosTXT  = "data/txt/funcionarios.txt"
)

type Funcionario struct {
	ID             string `json:"id"`
	Nome           string `json:"nome"`
	CPF            string `json:"cpf"`
	CEP            string `json:"cep"`
	Salario        string `json:"salario"`
	DataNascimento string `json:"data_nascimento"`
	Sexo           string `json:"sexo"`
	DepartamentoID int    `json:"departamento_id"`
}

// Create
func (f *Funcionario) Salvar(departamentos []int) string {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar funcionários: %v\n", err)
	}

	for _, funcionario := range funcionarios {
		if f.Nome == "" || f.CPF == "" || f.CEP == "" || f.Salario == "" || f.DataNascimento == "" || f.Sexo == "" {
			return "Erro: todos os campos devem ser preenchidos. Ação não realizada.\n"
		}
		if funcionario.ID == f.ID {
			return fmt.Sprintf("Erro: já existe um funcionário com o ID %s. Ação não realizada.\n", f.ID)
		}

		if funcionario.CPF == f.CPF {
			return fmt.Sprintf("Erro: já existe um funcionário com o CPF %s. Ação não realizada.\n", f.CPF)
		}
	}

	departamentoValido := false
	for _, departamentoID := range departamentos {
		if f.DepartamentoID == departamentoID {
			departamentoValido = true
			break
		}
	}

	if !departamentoValido {
		return fmt.Sprintf("DepartamentoID %d não encontrado", f.DepartamentoID)
	}

	funcionarios = append(funcionarios, *f)
	if err := salvarFuncionarios(funcionarios); err != nil {
		return fmt.Sprintf("Erro ao salvar funcionário: %v\n", err)
	}

	sincronizarFuncionariosTxt(funcionarios)
	return fmt.Sprintf("Funcionário com ID %s salvo com sucesso.\n", f.ID)
}

// Update
func (f *Funcionario) Atualizar() string {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar funcionários: %v\n", err)
	}

	for i, funcionario := range funcionarios {
		if funcionario.ID == f.ID {
			funcionarios[i] = *f
			if err := salvarFuncionarios(funcionarios); err != nil {
				return fmt.Sprintf("Erro ao atualizar funcionário: %v\n", err)
			}
			sincronizarFuncionariosTxt(funcionarios)
			return fmt.Sprintf("Funcionário com ID %s atualizado com sucesso.\n", f.ID)
		}
	}

	return fmt.Sprintf("Erro: funcionário com ID %s não encontrado para atualização. Ação não realizada.\n", f.ID)
}

// Delete
func (f *Funcionario) Deletar() string {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		return fmt.Sprintf("Erro ao carregar funcionários: %v\n", err)
	}

	for i, funcionario := range funcionarios {
		if funcionario.ID == f.ID {
			funcionarios = append(funcionarios[:i], funcionarios[i+1:]...)
			if err := salvarFuncionarios(funcionarios); err != nil {
				return fmt.Sprintf("Erro ao deletar funcionário: %v\n", err)
			}
			sincronizarFuncionariosTxt(funcionarios)
			return fmt.Sprintf("Funcionário com ID %s deletado com sucesso.\n", f.ID)
		}
	}

	return fmt.Sprintf("Erro: funcionário com ID %s não encontrado para exclusão. Ação não realizada.\n", f.ID)
}

// Read
func ListarFuncionarios() []Funcionario {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionários: %v\n", err)
		return nil
	}

	fmt.Println("Lista de funcionários:")
	for _, f := range funcionarios {
		fmt.Printf("ID: %s, Nome: %s, CPF: %s, CEP: %s, Salário: %s, Data de Nascimento: %s, Sexo: %s, DepartamentoID: %d\n",
			f.ID, f.Nome, f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo, f.DepartamentoID)
	}
	listaFuncionarios := []Funcionario{}
	listaFuncionarios = append(listaFuncionarios, funcionarios...)
	return listaFuncionarios
}

// Funções utilitárias

func CarregarFuncionarios() ([]Funcionario, error) {
	var funcionarios []Funcionario

	file, err := os.Open(arquivoFuncionariosJSON)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return funcionarios, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&funcionarios); err != nil {
		return nil, err
	}

	return funcionarios, nil
}

func salvarFuncionarios(funcionarios []Funcionario) error {
	file, err := os.Create(arquivoFuncionariosJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(funcionarios)
}

// Sincronizar o estado do JSON com o arquivo TXT
func sincronizarFuncionariosTxt(funcionarios []Funcionario) {
	file, err := os.Create(arquivoFuncionariosTXT)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo TXT: %v\n", err)
		return
	}
	defer file.Close()

	for _, f := range funcionarios {
		linha := fmt.Sprintf("ID: %s, Nome: %s, CPF: %s, CEP: %s, Salário: %s, Data de Nascimento: %s, Sexo: %s, DepartamentoID: %d\n",
			f.ID, f.Nome, f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo, f.DepartamentoID)
		_, err := file.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo TXT: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo TXT sincronizado com sucesso.")
}

func GetFuncionariosIDs() []string {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionarios: %v\n", err)
		return nil
	}

	var ids []string
	for _, f := range funcionarios {
		ids = append(ids, f.ID)
	}
	return ids
}
