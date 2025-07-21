package main_view

import (
	"fmt"
	"lottery-lose-easy/database"
	"lottery-lose-easy/views/table_views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App
var mainPage fyne.Window

func Init() fyne.App {
	if myApp == nil {
		myApp = app.New()
		return myApp
	}
	return myApp // cria com ID fixo
}

func GetMainScreen() fyne.Window {
	if mainPage == nil {
		mainPage = myApp.NewWindow("Loteria Perca Fácil")
		return mainPage
	}
	return mainPage
}

func MainPage() {
	// Inicializa o aplicativo
	myApp := Init()

	// Cria a janela principal
	mainPage = GetMainScreen()

	// Define o ícone da janela
	icon, err := fyne.LoadResourceFromPath("assets/imgs/LotericaPercaFacil.png")
	if err != nil {
		fmt.Println("Erro ao carregar o ícone:", err)
	} else {
		mainPage.SetIcon(icon) // Set the icon for the main window
	}

	// Define o tamanho da janela
	mainPage.Resize(fyne.NewSize(800, 600))

	// Botões para navegação
	btClients := widget.NewButton("Clientes", func() {
		mainPage.Hide()
		table_views.ClientePage(myApp, mainPage)
	})

	btFuncionarios := widget.NewButton("Funcionários", func() {
		mainPage.Hide()
		table_views.FuncionarioPage(myApp, mainPage)
	})

	btServicos := widget.NewButton("Serviços", func() {
		mainPage.Hide()
		table_views.ServicoPage(myApp, mainPage)
	})

	btAtendimentos := widget.NewButton("Atendimentos", func() {
		mainPage.Hide()
		table_views.AtendimentoPage(myApp, mainPage)
	})

	// Layout da barra superior
	barraSuperior := container.NewHBox(btClients, btFuncionarios, btServicos, btAtendimentos)

	// Layout principal
	mainPage.SetContent(
		container.NewBorder(
			barraSuperior, // Topo: Botões de navegação
			nil,           // Rodapé: Botão de voltar
			nil,           // Esquerda: vazio
			nil,           // Direita: vazio
			nil,           // Centro: Árvore
		),
	)

	mainPage.SetCloseIntercept(func() {
		db, _ := database.GetDbSession()
		_ = db.Close()
		fmt.Println("Database connection closed")
		mainPage.Close()
	})

	// Exibe a janela
	mainPage.ShowAndRun()
}
