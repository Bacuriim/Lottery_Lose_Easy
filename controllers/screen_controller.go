package controllers

import (
	"lottery-lose-easy/views/main_view"
	"lottery-lose-easy/views/table_views"
)

func Init() {
	myApp := main_view.Init()

	main_view.MainPage()
	table_views.ClientePage(myApp, main_view.GetMainScreen())
	table_views.FuncionarioPage(myApp, main_view.GetMainScreen())
	table_views.ServicoPage(myApp, main_view.GetMainScreen())
	table_views.AtendimentoPage(myApp, main_view.GetMainScreen())
	table_views.RelatoriosPage(myApp, main_view.GetMainScreen())
}
