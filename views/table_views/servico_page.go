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
)

func ServicoPage(myApp fyne.App, mainPage fyne.Window) {
	win := myApp.NewWindow("Serviços")
	win.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/LotericaPercaFacil.png")
	if err == nil {
		win.SetIcon(icon)
	}

	entryId := utils.CriarEntryNumeros("ID")
	servicoOptions := []string{
		"pagamento de boletos",
		"apostas em loterias",
		"saques",
		"depósitos",
	}
	selectNome := widget.NewSelect(servicoOptions, func(string) {})
	selectNome.PlaceHolder = "Selecione o serviço"

	resultLabel := widget.NewLabel("")

	addTab := container.NewVBox(
		widget.NewLabel("Adicionar Serviço"),
		entryId, selectNome,
		widget.NewButton("Criar serviço", func() {
			id, _ := strconv.Atoi(entryId.Text)
			servico := models.Servico{
				Id:          int32(id),
				NomeServico: selectNome.Selected,
			}
			msg := servico.Salvar(servico)
			resultLabel.SetText(msg)
		}),
		resultLabel,
	)

	servList := container.NewVBox()
	listTab := container.NewVBox(
		widget.NewLabel("Serviços cadastrados:"),
		servList,
		widget.NewButton("Listar Serviços", func() {
			servList.Objects = nil
			servicos, msg := new(models.Servico).BuscarTodos()
			if msg != "Serviços encontrados!" {
				resultLabel.SetText(msg)
			} else {
				for _, s := range servicos {
					card := widget.NewCard(
						s.NomeServico,
						fmt.Sprintf("ID: %d", s.Id),
						widget.NewLabel(""),
					)
					servList.Add(card)
				}
				resultLabel.SetText("Serviços listados com sucesso!")
			}
			servList.Refresh()
		}),
		resultLabel,
	)

	searchResult := widget.NewLabel("Resultado: nenhum")
	searchId := utils.CriarEntryNumeros("ID")

	searchTab := container.NewVBox(
		widget.NewLabel("Buscar Serviço por ID"),
		searchId,
		widget.NewButton("Buscar", func() {
			val := binding.NewString()
			val.Set(searchId.Text)
			s, msg := new(models.Servico).Pesquisar("id", val, true)
			if s != nil {
				searchResult.SetText(fmt.Sprintf("Encontrado: %s, ID: %d", s.NomeServico, s.Id))
			} else {
				searchResult.SetText(msg)
			}
		}),
		searchResult,
	)

	updateId := utils.CriarEntryNumeros("ID")
	updateSelectNome := widget.NewSelect([]string{
		"pagamento de boletos",
		"apostas em loterias",
		"saques",
		"depósitos",
	}, func(string) {})
	updateSelectNome.PlaceHolder = "Selecione o serviço"

	updateTab := container.NewVBox(
		widget.NewLabel("Atualizar Serviço"),
		updateId, updateSelectNome,
		widget.NewButton("Atualizar", func() {
			msg := ""
			if updateSelectNome.Selected != "" {
				msg = new(models.Servico).Alterar("nome_servico", updateSelectNome.Selected, "id", updateId.Text)
			}
			resultLabel.SetText(msg)
		}),
		resultLabel,
	)

	removeId := utils.CriarEntryNumeros("ID")

	removeTab := container.NewVBox(
		widget.NewLabel("Remover Serviço"),
		removeId,
		widget.NewButton("Remover", func() {
			val := binding.NewString()
			val.Set(removeId.Text)
			s, _ := new(models.Servico).Pesquisar("id", val, true)
			if s != nil {
				msg := new(models.Servico).Remover(*s)
				resultLabel.SetText(msg)
			} else {
				resultLabel.SetText("Serviço não encontrado para remover.")
			}
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
