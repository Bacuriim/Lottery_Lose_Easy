package table_views

import (
	"fyne.io/fyne/v2"
)

func ProjetosPage(myApp fyne.App, mainPage fyne.Window) {
	//projectsMainPage := myApp.NewWindow("Tabela de Projetos")
	//projectsMainPage.Resize(fyne.NewSize(800, 600))
	//
	//projectsListPage := myApp.NewWindow("Projetos Existentes")
	//projectsListPage.Resize(fyne.NewSize(800, 600))
	//
	//icon, err := fyne.LoadResourceFromPath("assets/imgs/CRUD_IMAGE.png") // Load the icon from a file
	//if err != nil {
	//	fmt.Println("Erro ao carregar o ícone:", err)
	//} else {
	//	projectsMainPage.SetIcon(icon)
	//	projectsListPage.SetIcon(icon) // Set the icon for the main window
	//}
	//
	//entryID := utils.CriarEntryNumeros("ID")
	//entryName := utils.CriarEntryLetrasNumeros("Nome do Projeto")
	//entryLocal := utils.CriarEntryLetrasNumeros("Local")
	//entryDepartmentID := utils.CriarEntryNumeros("ID do Departamento")
	//entryEmployeesProjectsID := utils.CriarEntryLetrasNumeros("ID dos Funcionários do Projeto")
	//
	//lbResultado := widget.NewLabel("Resultado: nenhum")
	//
	//listaProjetos := container.NewVBox()
	//btnCriar := widget.NewButton("Criar projeto", func() {
	//	projectID, err := strconv.Atoi(entryID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryID.SetText(strconv.Itoa(projectID))
	//	}
	//
	//	depID, err := strconv.Atoi(entryDepartmentID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryDepartmentID.SetText(strconv.Itoa(depID))
	//	}
	//
	//	projeto := models.Projeto{
	//		ID:                     projectID,
	//		Nome:                   entryName.Text,
	//		Local:                  entryLocal.Text,
	//		DepartamentoID:         depID,
	//		FuncionariosProjetosID: entryEmployeesProjectsID.Text,
	//	}
	//	lbResultado.SetText("Resultado: " + projeto.Salvar(models.GetDepartamentosIDs(), models.GetFuncionariosIDs()))
	//	utils.LimparCampos(entryID, entryName, entryLocal, entryDepartmentID,
	//		entryEmployeesProjectsID)
	//})
	//
	//btnAlterar := widget.NewButton("Alterar projetos", func() {
	//	projectID, err := strconv.Atoi(entryID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryID.SetText(strconv.Itoa(projectID))
	//	}
	//
	//	depID, err := strconv.Atoi(entryDepartmentID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryDepartmentID.SetText(strconv.Itoa(depID))
	//	}
	//
	//	projeto := models.Projeto{
	//		ID:                     projectID,
	//		Nome:                   entryName.Text,
	//		Local:                  entryLocal.Text,
	//		DepartamentoID:         depID,
	//		FuncionariosProjetosID: entryEmployeesProjectsID.Text,
	//	}
	//	lbResultado.SetText("Resultado: " + projeto.Atualizar())
	//	utils.LimparCampos(entryID, entryName, entryLocal, entryDepartmentID,
	//		entryEmployeesProjectsID)
	//})
	//
	//btnListar := widget.NewButton("Listar Projetos", func() {
	//	projetos, err := models.CarregarProjetos()
	//	if err != nil {
	//		fmt.Println("Erro:", err)
	//		return
	//	}
	//
	//	listaProjetos.RemoveAll()
	//
	//	for _, p := range projetos {
	//		card := widget.NewCard(p.Nome,
	//			"",
	//			widget.NewLabel(fmt.Sprintf("ID: %d\nLocal: %s\nID do Departamento: %d\nID dos Funcionários do Projeto: %s",
	//				p.ID, p.Local, p.DepartamentoID, p.FuncionariosProjetosID)))
	//		listaProjetos.Add(card)
	//	}
	//	projectsListPage.Show()
	//
	//	listaProjetos.Refresh()
	//	projectsListPage.SetCloseIntercept(func() {
	//		projectsListPage.Hide()
	//	})
	//})
	//
	//btnDeletar := widget.NewButton("Deletar projeto", func() {
	//	projectID, err := strconv.Atoi(entryID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryID.SetText(strconv.Itoa(projectID))
	//	}
	//
	//	projeto := models.Projeto{
	//		ID:                     projectID,
	//		Nome:                   entryName.Text,
	//		Local:                  entryLocal.Text,
	//		DepartamentoID:         0,
	//		FuncionariosProjetosID: entryEmployeesProjectsID.Text,
	//	}
	//	lbResultado.SetText("Resultado: " + projeto.Deletar())
	//	utils.LimparCampos(entryID, entryName, entryLocal, entryDepartmentID,
	//		entryEmployeesProjectsID)
	//})
	//
	//projectsListPage.SetContent(container.NewScroll(
	//	listaProjetos,
	//))
	//
	//rodape := widget.NewButton("Voltar", func() {
	//	projectsMainPage.Hide()
	//	projectsListPage.Hide()
	//	mainPage.Show()
	//})
	//
	//projectsMainPage.SetContent(container.NewBorder(
	//	container.NewVBox(
	//		entryID,
	//		entryName,
	//		entryLocal,
	//		entryDepartmentID,
	//		entryEmployeesProjectsID,
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
	//projectsMainPage.SetCloseIntercept(func() {
	//	utils.EsvaziarArquivoJSON("data/funcionarios.json")
	//	utils.EsvaziarArquivoJSON("data/departamentos.json")
	//	utils.EsvaziarArquivoJSON("data/chefes_departamento.json")
	//	utils.EsvaziarArquivoJSON("data/funcionarios_projetos.json")
	//	utils.EsvaziarArquivoJSON("data/projetos.json")
	//	projectsMainPage.Close()
	//	mainPage.Show()
	//})
	//
	//projectsMainPage.Show()
}
