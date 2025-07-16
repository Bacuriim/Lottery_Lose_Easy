package table_views

import (
	//"fmt"
	//"lottery-lose-easy/models"
	//"lottery-lose-easy/utils"
	//"strconv"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
)

func ChefeDepartamentoPage(myApp fyne.App, mainPage fyne.Window) {
	//chefeDepartamentoMainPage := myApp.NewWindow("Tabela de Chefes de Departamento")
	//chefeDepartamentoMainPage.Resize(fyne.NewSize(800, 600))
	//
	//chefeDepartamentoListPage := myApp.NewWindow("Chefes de Departamento Existentes")
	//chefeDepartamentoListPage.Resize(fyne.NewSize(800, 600))
	//
	//icon, err := fyne.LoadResourceFromPath("assets/imgs/CRUD_IMAGE.png") // Load the icon from a file
	//if err != nil {
	//	fmt.Println("Erro ao carregar o ícone:", err)
	//} else {
	//	chefeDepartamentoMainPage.SetIcon(icon)
	//	chefeDepartamentoListPage.SetIcon(icon) // Set the icon for the main window
	//}
	//
	//entryID := utils.CriarEntryNumeros("ID")
	//entryFuncionarioID := utils.CriarEntryLetrasNumeros("ID do Funcionário")
	//
	//lbResultado := widget.NewLabel("Resultado: nenhum")
	//
	//listaChefesDepartamentos := container.NewVBox()
	//btnCriar := widget.NewButton("Criar chefe de departamento", func() {
	//	chefeID, err := strconv.Atoi(entryID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do Chefe:", err)
	//		return
	//	} else {
	//		entryID.SetText(strconv.Itoa(chefeID))
	//	}
	//	chefeDepartamento := models.ChefeDepartamento{
	//		ID:            chefeID,
	//		FuncionarioID: entryFuncionarioID.Text,
	//	}
	//	lbResultado.SetText("Resultado: " + chefeDepartamento.Salvar(models.GetFuncionariosIDs()))
	//	utils.LimparCampos(entryID, entryFuncionarioID)
	//})
	//
	//btnAlterar := widget.NewButton("Alterar chefe de departamento", func() {
	//	chefeID, err := strconv.Atoi(entryID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do Chefe:", err)
	//		return
	//	} else {
	//		entryID.SetText(strconv.Itoa(chefeID))
	//	}
	//
	//	chefeDepartamento := models.ChefeDepartamento{
	//		ID:            chefeID,
	//		FuncionarioID: entryFuncionarioID.Text,
	//	}
	//
	//	lbResultado.SetText("Resultado: " + chefeDepartamento.Atualizar())
	//	utils.LimparCampos(entryID, entryFuncionarioID)
	//})
	//
	//btnListar := widget.NewButton("Listar Chefes de Departamentos", func() {
	//	chefeDepartamentos, err := models.CarregarChefesDepartamento()
	//	if err != nil {
	//		fmt.Println("Erro:", err)
	//		return
	//	}
	//
	//	listaChefesDepartamentos.RemoveAll()
	//
	//	for _, c := range chefeDepartamentos {
	//		card := widget.NewCard("ID do Chefe: "+strconv.Itoa(c.ID),
	//			"",
	//			widget.NewLabel(fmt.Sprintf("FuncionárioID: %s",
	//				c.FuncionarioID)))
	//		listaChefesDepartamentos.Add(card)
	//	}
	//	chefeDepartamentoListPage.Show()
	//
	//	listaChefesDepartamentos.Refresh()
	//	chefeDepartamentoListPage.SetCloseIntercept(func() {
	//		chefeDepartamentoListPage.Hide()
	//	})
	//})
	//
	//btnDeletar := widget.NewButton("Deletar chefe de departamento", func() {
	//	chefeID, err := strconv.Atoi(entryID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do Chefe:", err)
	//		return
	//	} else {
	//		entryID.SetText(strconv.Itoa(chefeID))
	//	}
	//	chefeDepartamento := models.ChefeDepartamento{
	//		ID: chefeID,
	//	}
	//	lbResultado.SetText("Resultado: " + chefeDepartamento.Deletar())
	//	utils.LimparCampos(entryID, entryFuncionarioID)
	//})
	//
	//chefeDepartamentoListPage.SetContent(container.NewScroll(
	//	listaChefesDepartamentos,
	//))
	//
	//rodape := widget.NewButton("Voltar", func() {
	//	chefeDepartamentoMainPage.Hide()
	//	chefeDepartamentoListPage.Hide()
	//	mainPage.Show()
	//})
	//
	//chefeDepartamentoMainPage.SetContent(container.NewBorder(
	//	container.NewVBox(
	//		entryID,
	//		entryFuncionarioID,
	//	),
	//	rodape,
	//	nil,
	//	nil,
	//	container.NewVBox(
	//		btnCriar,
	//		btnAlterar,
	//		btnListar,
	//		btnDeletar,
	//		lbResultado,
	//	),
	//))
	//
	//chefeDepartamentoMainPage.SetCloseIntercept(func() {
	//	utils.EsvaziarArquivoJSON("data/Departamentos.json")
	//	utils.EsvaziarArquivoJSON("data/departamentos.json")
	//	utils.EsvaziarArquivoJSON("data/chefes_departamento.json")
	//	utils.EsvaziarArquivoJSON("data/Departamentos_projetos.json")
	//	utils.EsvaziarArquivoJSON("data/projetos.json")
	//	chefeDepartamentoMainPage.Close()
	//	mainPage.Show()
	//})
	//
	//chefeDepartamentoMainPage.Show()
}
