package main_view

import (
	"fmt"
	"lottery-lose-easy/database"
	"lottery-lose-easy/views/table_views"
	"os"
	"strings"

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
		mainPage = myApp.NewWindow("CRUD Empresarial")
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
	icon, err := fyne.LoadResourceFromPath("assets/imgs/CRUD_IMAGE.png") // Load the icon from a file
	if err != nil {
		fmt.Println("Erro ao carregar o ícone:", err)
	} else {
		mainPage.SetIcon(icon) // Set the icon for the main window
	}

	// Define o tamanho da janela
	mainPage.Resize(fyne.NewSize(800, 600))

	// Lê os dados estruturados da árvore
	data := loadTreeData()

	// Cria a árvore
	tree := widget.NewTree(
		func(uid string) []string {
			// Retorna os filhos do nó atual
			return data[uid]
		},
		func(uid string) bool {
			// Verifica se o nó é um ramo (tem filhos)
			_, ok := data[uid]
			return ok
		},
		func(branch bool) fyne.CanvasObject {
			// Cria o componente visual para cada nó
			return widget.NewLabel("")
		},
		func(uid string, branch bool, obj fyne.CanvasObject) {
			// Atualiza o conteúdo do nó
			obj.(*widget.Label).SetText(uid)
		},
	)

	data = loadTreeData()

	// Define o nó raiz da árvore
	tree.Root = "CRUD Empresarial"

	tree.Resize(fyne.NewSize(400, 300))

	// Botões para navegação
	btFuncionarios := widget.NewButton("Funcionários", func() {
		mainPage.Hide()
		table_views.FuncionariosPage(myApp, mainPage)
	})

	btDepartamentos := widget.NewButton("Departamentos", func() {
		mainPage.Hide()
		table_views.DepartamentosPage(myApp, mainPage)
	})

	btChefes := widget.NewButton("Chefes de Departamentos", func() {
		mainPage.Hide()
		table_views.ChefeDepartamentoPage(myApp, mainPage)
	})

	btProjetos := widget.NewButton("Projetos", func() {
		mainPage.Hide()
		table_views.ProjetosPage(myApp, mainPage)
	})

	// Layout da barra superior
	barraSuperior := container.NewHBox(btFuncionarios, btDepartamentos, btChefes, btProjetos)

	// Layout principal
	mainPage.SetContent(
		container.NewBorder(
			barraSuperior, // Topo: Botões de navegação
			nil,           // Rodapé: Botão de voltar
			nil,           // Esquerda: vazio
			nil,           // Direita: vazio
			tree,          // Centro: Árvore
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

func loadTreeData() map[string][]string {
	data := make(map[string][]string)

	// Caminho da pasta dos arquivos .txt
	dir := "./data/txt"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Erro ao ler o diretório:", err)
		return data
	}

	// Adiciona a raiz
	data["CRUD Empresarial"] = []string{}

	// Processa cada arquivo .txt
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			content, err := os.ReadFile(dir + "/" + file.Name())
			if err != nil {
				fmt.Println("Erro ao ler o arquivo:", file.Name(), err)
				continue
			}

			// Adiciona o arquivo como filho da raiz
			data["CRUD Empresarial"] = append(data["CRUD Empresarial"], file.Name())

			// Processa cada linha do arquivo
			lines := strings.Split(strings.TrimSpace(string(content)), "\n")
			for _, line := range lines {
				// Usa a linha inteira como um nó filho
				data[file.Name()] = append(data[file.Name()], line)
			}
		}
	}

	return data
}
