package table_views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lottery-lose-easy/models"
)

func RelatoriosPage(myApp fyne.App, mainPage fyne.Window) {
	relatoriosWindow := myApp.NewWindow("Relatórios")
	relatoriosWindow.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/LotericaPercaFacil.png")
	if err != nil {
		fmt.Println("Erro ao carregar o ícone:", err)
	} else {
		relatoriosWindow.SetIcon(icon)
	}

	resultLabel := widget.NewLabel("")

	// === ABA DE SERVIÇOS MAIS UTILIZADOS ===
	listaServicos := container.NewVBox()

	tabServicos := container.NewVBox(
		widget.NewLabel("Serviços mais utilizados:"),
		listaServicos,
		widget.NewButton("Listar Serviços", func() {
			listaServicos.Objects = nil
			servicos, msg := models.ListarServicosMaisUtilizados()
			if msg != "Serviços mais utilizados encontrados!" {
				resultLabel.SetText(msg)
			} else {
				for _, s := range servicos {
					card := widget.NewCard(
						s.Nome,
						fmt.Sprintf("Total de utilizações: %d", s.Total),
						nil,
					)
					listaServicos.Add(card)
				}
				resultLabel.SetText("Serviços listados com sucesso!")
			}
			listaServicos.Refresh()
		}),
		resultLabel,
	)

	// === ABA DE MEIOS DE PAGAMENTO MAIS UTILIZADOS ===
	listaPagamentos := container.NewVBox()

	tabPagamentos := container.NewVBox(
		widget.NewLabel("Meios de pagamento mais utilizados:"),
		listaPagamentos,
		widget.NewButton("Listar Pagamentos", func() {
			listaPagamentos.Objects = nil
			pagamentos, msg := models.ListarPagamentosMaisUtilizados()
			if msg != "Pagamentos mais utilizados encontrados!" {
				resultLabel.SetText(msg)
			} else {
				for _, p := range pagamentos {
					card := widget.NewCard(
						p.Forma,
						fmt.Sprintf("Total de vezes usado: %d", p.Total),
						nil,
					)
					listaPagamentos.Add(card)
				}
				resultLabel.SetText("Pagamentos listados com sucesso!")
			}
			listaPagamentos.Refresh()
		}),
		resultLabel,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Serviços", tabServicos),
		container.NewTabItem("Pagamentos", tabPagamentos),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	relatoriosWindow.SetCloseIntercept(func() {
		relatoriosWindow.Hide()
		mainPage.Show()
	})

	relatoriosWindow.SetContent(tabs)
	relatoriosWindow.Show()
}
