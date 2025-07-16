package table_views

import "fyne.io/fyne/v2"

func ClientePage(myApp fyne.App, mainPage fyne.Window) {
	clienteMainPage := myApp.NewWindow("Clientes")
	clienteMainPage.Resize(fyne.NewSize(800, 600))

	clienteSearchPage := myApp.NewWindow("Buscar Clientes")
	clienteSearchPage.Resize(fyne.NewSize(800, 600))

	clienteInsertPage := myApp.NewWindow("Adicionar Clientes")
	clienteInsertPage.Resize(fyne.NewSize(800, 600))

	clienteUpdatePage := myApp.NewWindow("Atualizar Clientes")
	clienteUpdatePage.Resize(fyne.NewSize(800, 600))

	clienteDeletePage := myApp.NewWindow("Remover Clientes")
	clienteDeletePage.Resize(fyne.NewSize(800, 600))
}
