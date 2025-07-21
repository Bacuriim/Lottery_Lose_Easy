package table_views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lottery-lose-easy/utils"
)

func ClientePage(myApp fyne.App, mainPage fyne.Window) {
	clientMainPage := myApp.NewWindow("Clientes")
	clientMainPage.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/CRUD_IMAGE.png")
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

	addTab := container.NewVBox(
		widget.NewLabel("Adicionar Cliente"),
		entryName, entryCpf, entryGender, entryAge, entryAddress, entryAccount,
		widget.NewButton("Criar cliente", func() {
			// TODO: Salvar cliente no banco
			fmt.Println("Cliente criado (simulação).")
		}),
	)

	clientList := container.NewVBox()
	listTab := container.NewVBox(
		widget.NewLabel("Clientes cadastrados:"),
		clientList,
		widget.NewButton("Listar Clientes", func() {
			clientList.Objects = nil // limpa a lista
			// TODO: Buscar clientes do banco e preencher
			clientList.Add(widget.NewLabel("-> Cliente Exemplo"))
			clientList.Refresh()
		}),
	)

	searchResult := widget.NewLabel("Resultado: nenhum")
	searchCpf := utils.CriarEntryNumeros("CPF para buscar")

	searchTab := container.NewVBox(
		widget.NewLabel("Buscar Cliente por CPF"),
		searchCpf,
		widget.NewButton("Buscar", func() {
			// TODO: Buscar cliente no banco
			searchResult.SetText("Cliente encontrado (simulação)")
		}),
		searchResult,
	)

	updateCpf := utils.CriarEntryNumeros("CPF para atualizar")

	updateTab := container.NewVBox(
		widget.NewLabel("Atualizar Cliente"),
		updateCpf,
		entryName, entryGender, entryAge, entryAddress, entryAccount,
		widget.NewButton("Atualizar", func() {
			// TODO: Atualizar cliente no banco
			fmt.Println("Cliente atualizado (simulação).")
		}),
	)

	removeCpf := utils.CriarEntryNumeros("CPF para remover")

	removeTab := container.NewVBox(
		widget.NewLabel("Remover Cliente"),
		removeCpf,
		widget.NewButton("Remover", func() {
			// TODO: Remover cliente do banco
			fmt.Println("Cliente removido (simulação).")
		}),
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
