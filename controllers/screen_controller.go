package controllers

import (
	"lottery-lose-easy/views/main_view"
	"lottery-lose-easy/views/table_views"
)

func Init() {
	myApp := main_view.Init()

	main_view.MainPage()
	table_views.FuncionariosPage(myApp, main_view.GetMainScreen())
	table_views.DepartamentosPage(myApp, main_view.GetMainScreen())
	table_views.ChefeDepartamentoPage(myApp, main_view.GetMainScreen())
	table_views.ProjetosPage(myApp, main_view.GetMainScreen())
}
