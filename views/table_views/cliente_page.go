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

func ClientePage(myApp fyne.App, mainPage fyne.Window) {
	clientMainPage := myApp.NewWindow("Clientes")
	clientMainPage.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/LotericaPercaFacil.png")
	if err != nil {
		fmt.Println("Erro ao carregar o ícone:", err)
	} else {
		clientMainPage.SetIcon(icon)
	}

	entryName := utils.CriarEntryLetras("Nome")
	entryCpf := utils.CriarEntryNumeros("CPF")
	entryGender := utils.CriarEntryLetras("Sexo")
	entryAge := utils.CriarEntryNumeros("Idade")
	entryAddress := utils.CriarEntryLetrasNumeros("Endereço")
	entryAccount := utils.CriarEntryNumeros("Conta")

	resultLabel := widget.NewLabel("")

	addTab := container.NewVBox(
		widget.NewLabel("Adicionar Cliente"),
		entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount,
		widget.NewButton("Criar cliente", func() {
			if utils.AtLeastOneEntryNil(entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount) {
				resultLabel.SetText("Preencha todos os campos!")
				return
			}
			idade, _ := strconv.Atoi(entryAge.Text)
			cliente := models.Cliente{
				Id:       uuid.New(),
				Nome:     entryName.Text,
				Cpf:      entryCpf.Text,
				Sexo:     entryGender.Text,
				Idade:    idade,
				Endereco: entryAddress.Text,
				Conta:    entryAccount.Text,
			}
			msg := cliente.Salvar(cliente)
			resultLabel.SetText(msg)
			utils.LimparCampos(entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount)
		}),
		resultLabel,
	)

	clientList := container.NewVBox()
	listTab := container.NewVBox(
		widget.NewLabel("Clientes cadastrados:"),
		clientList,
		widget.NewButton("Listar Clientes", func() {
			clientList.Objects = nil
			clientes, msg := new(models.Cliente).BuscarTodos()
			if msg != "Clientes encontrados!" {
				resultLabel.SetText(msg)
			} else {
				for _, c := range clientes {
					card := widget.NewCard(
						c.Nome,
						fmt.Sprintf("CPF: %s", c.Cpf),
						widget.NewLabel(fmt.Sprintf(
							"Sexo: %s\nIdade: %d\nEndereço: %s\nConta: %s",
							c.Sexo, c.Idade, c.Endereco, c.Conta,
						)),
					)
					clientList.Add(card)
				}
				resultLabel.SetText("Clientes listados com sucesso!")
			}
			clientList.Refresh()
		}),
		resultLabel,
	)

	searchResult := widget.NewLabel("Resultado: nenhum")
	searchCpf := utils.CriarEntryNumeros("CPF")

	searchTab := container.NewVBox(
		widget.NewLabel("Buscar Cliente por CPF"),
		searchCpf,
		widget.NewButton("Buscar", func() {
			if utils.AtLeastOneEntryNil(searchCpf) {
				resultLabel.SetText("Preencha todos os campos!")
				return
			}
			val := binding.NewString()
			val.Set(searchCpf.Text)
			c, msg := new(models.Cliente).Pesquisar("cpf", val, false)
			if c != nil {
				searchResult.SetText(fmt.Sprintf("Encontrado: %s, CPF: %s", c.Nome, c.Cpf))
			} else {
				searchResult.SetText(msg)
			}
			utils.LimparCampos(entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount)
		}),
		searchResult,
	)

	updateCpf := utils.CriarEntryNumeros("CPF")

	updateTab := container.NewVBox(
		widget.NewLabel("Atualizar Cliente"),
		updateCpf,
		entryName, entryGender, entryAge, entryAddress, entryAccount,
		widget.NewButton("Atualizar", func() {
			if utils.AtLeastOneEntryNil(updateCpf) {
				resultLabel.SetText("Preencha o cpf a ser alterado!")
				return
			}
			msg := ""
			if entryName.Text != "" {
				msg = new(models.Cliente).Alterar("nome", entryName.Text, "cpf", updateCpf.Text)
			}
			if entryGender.Text != "" {
				msg = new(models.Cliente).Alterar("sexo", entryGender.Text, "cpf", updateCpf.Text)
			}
			if entryAge.Text != "" {
				idade, _ := strconv.Atoi(entryAge.Text)
				msg = new(models.Cliente).Alterar("idade", idade, "cpf", updateCpf.Text)
			}
			if entryAddress.Text != "" {
				msg = new(models.Cliente).Alterar("endereco", entryAddress.Text, "cpf", updateCpf.Text)
			}
			if entryAccount.Text != "" {
				msg = new(models.Cliente).Alterar("conta", entryAccount.Text, "cpf", updateCpf.Text)
			}
			resultLabel.SetText(msg)
			utils.LimparCampos(entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount)
		}),
		resultLabel,
	)

	removeCpf := utils.CriarEntryNumeros("CPF")

	removeTab := container.NewVBox(
		widget.NewLabel("Remover Cliente"),
		removeCpf,
		widget.NewButton("Remover", func() {
			if utils.AtLeastOneEntryNil(removeCpf) {
				resultLabel.SetText("Preencha o cpf do usuário a ser removido")
				return
			}
			val := binding.NewString()
			val.Set(removeCpf.Text)
			c, _ := new(models.Cliente).Pesquisar("cpf", val, false)
			if c != nil {
				msg := new(models.Cliente).Remover(*c)
				resultLabel.SetText(msg)
			} else {
				resultLabel.SetText("Cliente não encontrado para remover.")
			}
			utils.LimparCampos(entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount)
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

	clientMainPage.SetCloseIntercept(func() {
		clientMainPage.Hide()
		mainPage.Show()
	})

	tabs.SetTabLocation(container.TabLocationTop)

	clientMainPage.SetContent(tabs)
	clientMainPage.Show()
}
