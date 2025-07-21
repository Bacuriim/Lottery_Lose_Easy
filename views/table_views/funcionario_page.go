package table_views

import (
	"fmt"
	"lottery-lose-easy/models"
	"lottery-lose-easy/utils"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func FuncionarioPage(myApp fyne.App, mainPage fyne.Window) {
	win := myApp.NewWindow("Funcionários")
	win.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/LotericaPercaFacil.png")
	if err == nil {
		win.SetIcon(icon)
	}

	entryNome := utils.CriarEntryLetras("Nome")
	entryIdent := utils.CriarEntryNumeros("Número Identificação")
	entrySalario := utils.CriarEntryNumeros("Salário")
	entryCpf := utils.CriarEntryNumeros("CPF")
	entrySexo := utils.CriarEntryLetras("Sexo")
	entryIdade := utils.CriarEntryNumeros("Idade")
	entryHorario := utils.CriarEntryLetrasNumeros("Horário Trabalho")

	resultLabel := widget.NewLabel("")

	addTab := container.NewVBox(
		widget.NewLabel("Adicionar Funcionário"),
		entryNome, entryIdent, entrySalario, entryCpf, entrySexo, entryIdade, entryHorario,
		widget.NewButton("Criar funcionário", func() {
			if utils.AtLeastOneEntryNil(entryNome, entryIdent, entrySalario, entryCpf, entrySexo, entryIdade, entryHorario) {
				resultLabel.SetText("Preencha todos os campos!")
				return
			}
			idade, _ := strconv.Atoi(entryIdade.Text)
			salario, _ := strconv.ParseFloat(entrySalario.Text, 64)
			funcionario := models.Funcionario{
				Id:                  uuid.New(),
				Nome:                entryNome.Text,
				NumeroIdentificacao: entryIdent.Text,
				Salario:             salario,
				Cpf:                 entryCpf.Text,
				Sexo:                entrySexo.Text,
				Idade:               idade,
				HorarioTrabalho:     entryHorario.Text,
			}
			msg := funcionario.Salvar(funcionario)
			resultLabel.SetText(msg)
			utils.LimparCampos(entryNome, entryIdent, entrySalario, entryCpf, entrySexo, entryIdade, entryHorario)
		}),
		resultLabel,
	)

	funcList := container.NewVBox()
	listTab := container.NewVBox(
		widget.NewLabel("Funcionários cadastrados:"),
		funcList,
		widget.NewButton("Listar Funcionários", func() {
			funcList.Objects = nil
			funcionarios, msg := new(models.Funcionario).BuscarTodos()
			if msg != "Funcionários encontrados!" {
				resultLabel.SetText(msg)
			} else {
				for _, f := range funcionarios {
					card := widget.NewCard(
						f.Nome,
						fmt.Sprintf("CPF: %s", f.Cpf),
						widget.NewLabel(fmt.Sprintf(
							"Identificação: %s\nSalário: %.2f\nSexo: %s\nIdade: %d\nHorário: %s",
							f.NumeroIdentificacao, f.Salario, f.Sexo, f.Idade, f.HorarioTrabalho,
						)),
					)
					funcList.Add(card)
				}
				resultLabel.SetText("Funcionários listados com sucesso!")
			}
			funcList.Refresh()
		}),
		resultLabel,
	)

	searchResult := widget.NewLabel("Resultado: nenhum")
	searchCpf := utils.CriarEntryNumeros("CPF")

	searchTab := container.NewVBox(
		widget.NewLabel("Buscar Funcionário por CPF"),
		searchCpf,
		widget.NewButton("Buscar", func() {
			if utils.AtLeastOneEntryNil(searchCpf) {
				resultLabel.SetText("Preencha o cpf a ser buscado!")
				return
			}
			val := binding.NewString()
			val.Set(searchCpf.Text)
			f, msg := new(models.Funcionario).Pesquisar("cpf", val, false)
			if f != nil {
				searchResult.SetText(fmt.Sprintf("Encontrado: %s, CPF: %s", f.Nome, f.Cpf))
			} else {
				searchResult.SetText(msg)
			}
			utils.LimparCampos(entryNome, entryIdent, entrySalario, entryCpf, entrySexo, entryIdade, entryHorario)
		}),
		searchResult,
	)

	updateCpf := utils.CriarEntryNumeros("CPF")

	updateTab := container.NewVBox(
		widget.NewLabel("Atualizar Funcionário"),
		updateCpf,
		entryNome, entryIdent, entrySalario, entrySexo, entryIdade, entryHorario,
		widget.NewButton("Atualizar", func() {
			if utils.AtLeastOneEntryNil(updateCpf) {
				resultLabel.SetText("Preencha o cpf do funcionário a ser alterado")
				return
			}
			msg := ""
			if entryNome.Text != "" {
				msg = new(models.Funcionario).Alterar("nome", entryNome.Text, "cpf", updateCpf.Text)
			}
			if entryIdent.Text != "" {
				msg = new(models.Funcionario).Alterar("numero_identificacao", entryIdent.Text, "cpf", updateCpf.Text)
			}
			if entrySalario.Text != "" {
				salario, _ := strconv.ParseFloat(entrySalario.Text, 64)
				msg = new(models.Funcionario).Alterar("salario", salario, "cpf", updateCpf.Text)
			}
			if entrySexo.Text != "" {
				msg = new(models.Funcionario).Alterar("sexo", entrySexo.Text, "cpf", updateCpf.Text)
			}
			if entryIdade.Text != "" {
				idade, _ := strconv.Atoi(entryIdade.Text)
				msg = new(models.Funcionario).Alterar("idade", idade, "cpf", updateCpf.Text)
			}
			if entryHorario.Text != "" {
				msg = new(models.Funcionario).Alterar("horario_trabalho", entryHorario.Text, "cpf", updateCpf.Text)
			}
			resultLabel.SetText(msg)
			utils.LimparCampos(entryNome, entryIdent, entrySalario, entryCpf, entrySexo, entryIdade, entryHorario)
		}),
		resultLabel,
	)

	removeCpf := utils.CriarEntryNumeros("CPF")

	removeTab := container.NewVBox(
		widget.NewLabel("Remover Funcionário"),
		removeCpf,
		widget.NewButton("Remover", func() {
			if utils.AtLeastOneEntryNil(removeCpf) {
				resultLabel.SetText("Preencha o cpf do funcionário a ser removido")
				return
			}
			val := binding.NewString()
			val.Set(removeCpf.Text)
			f, _ := new(models.Funcionario).Pesquisar("cpf", val, false)
			if f != nil {
				msg := new(models.Funcionario).Remover(*f)
				resultLabel.SetText(msg)
			} else {
				resultLabel.SetText("Funcionário não encontrado para remover.")
			}
			utils.LimparCampos(entryNome, entryIdent, entrySalario, entryCpf, entrySexo, entryIdade, entryHorario)
		}),
		resultLabel,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Adicionar", addTab),
		container.NewTabItem("Listar", listTab),
		container.NewTabItem("Buscar", searchTab),
		container.NewTabItem("Atualizar", updateTab),
		container.NewTabItem("Remover", removeTab),
	)

	win.SetCloseIntercept(func() {
		win.Hide()
		mainPage.Show()
	})

	tabs.SetTabLocation(container.TabLocationTop)
	win.SetContent(tabs)
	win.Show()
}
