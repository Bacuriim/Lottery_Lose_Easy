package main_view

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"image/color"
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
		mainPage = myApp.NewWindow("Lotérica Perca Fácil")
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

	// Label do trabalho
	txLottery := canvas.NewText("Lotérica Perca Fácil", color.RGBA{R: 173, G: 216, B: 240, A: 255})
	txLottery.Alignment = fyne.TextAlignCenter
	txLottery.TextSize = 30
	txLottery.Resize(fyne.NewSize(800, 300))

	txParticipants := canvas.NewText("Conrado Einstein - Hiel Saraiva - João Marcelo Pimenta", color.RGBA{R: 255, G: 255, B: 255, A: 255})
	txParticipants.Alignment = fyne.TextAlignCenter
	txParticipants.Resize(fyne.NewSize(800, 200))

	// Botões para navegação
	btClients := widget.NewButton("Clientes", func() {
		mainPage.Hide()
		table_views.ClientePage(myApp, mainPage)
	})
	btClients.Resize(fyne.NewSize(400, 100))

	btFuncionarios := widget.NewButton("Funcionários", func() {
		mainPage.Hide()
		table_views.FuncionarioPage(myApp, mainPage)
	})
	btFuncionarios.Resize(fyne.NewSize(400, 100))

	btServicos := widget.NewButton("Serviços", func() {
		mainPage.Hide()
		table_views.ServicoPage(myApp, mainPage)
	})
	btServicos.Resize(fyne.NewSize(400, 100))

	btAtendimentos := widget.NewButton("Atendimentos", func() {
		mainPage.Hide()
		table_views.AtendimentoPage(myApp, mainPage)
	})
	btAtendimentos.Resize(fyne.NewSize(400, 100))

	// Layout da barra superior
	barraSuperior := container.NewVBox(txLottery, txParticipants, btClients, btFuncionarios, btServicos, btAtendimentos)

	// Layout principal
	mainPage.SetContent(
		container.NewBorder(
			barraSuperior, // Topo: Botões de navegação
			nil,           // Rodapé: Botão de voltar
			nil,           // Esquerda: vazio
			nil,           // Direita: vazio
			nil,
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
