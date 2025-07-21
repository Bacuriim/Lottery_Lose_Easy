package table_views

import (
	"fyne.io/fyne/v2"
)

var isShowing = false

func FuncionariosPage(myApp fyne.App, mainPage fyne.Window) {
	//employeesMainPage := myApp.NewWindow("Tabela de Funcionários")
	//employeesMainPage.Resize(fyne.NewSize(800, 600))
	//
	//employeesListPage := myApp.NewWindow("Funcionários Existentes")
	//employeesListPage.Resize(fyne.NewSize(800, 600))
	//
	//icon, err := fyne.LoadResourceFromPath("assets/imgs/CRUD_IMAGE.png") // Load the icon from a file
	//if err != nil {
	//	fmt.Println("Erro ao carregar o ícone:", err)
	//} else {
	//	employeesMainPage.SetIcon(icon)
	//	employeesListPage.SetIcon(icon) // Set the icon for the main window
	//}
	//
	//entryID := utils.CriarEntryLetrasNumeros("ID")
	//entryName := utils.CriarEntryLetrasNumeros("Nome")
	//entryCPF := utils.CriarEntryNumeros("CPF")
	//entryCEP := utils.CriarEntryNumeros("CEP")
	//entrySalario := utils.CriarEntryNumeros("Salário")
	//entryDataNascimento := utils.CriarEntryData("Data de Nascimento")
	//entrySexo := utils.CriarEntryLetras("Sexo")
	//entryDepartamentoID := utils.CriarEntryNumeros("ID do Departamento")
	//
	//lbResultado := widget.NewLabel("Resultado: nenhum")
	//
	//listaFuncionarios := container.NewVBox()
	//btnCriar := widget.NewButton("Criar funcionário", func() {
	//	depID, err := strconv.Atoi(entryDepartamentoID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryDepartamentoID.SetText(strconv.Itoa(depID))
	//	}
	//	funcionario := models.Funcionario{
	//		ID:             entryID.Text,
	//		Nome:           entryName.Text,
	//		CPF:            entryCPF.Text,
	//		CEP:            entryCEP.Text,
	//		Salario:        entrySalario.Text,
	//		DataNascimento: entryDataNascimento.Text,
	//		Sexo:           entrySexo.Text,
	//		DepartamentoID: depID,
	//	}
	//	lbResultado.SetText("Resultado: " + funcionario.Salvar(models.GetDepartamentosIDs()))
	//	utils.LimparCampos(entryID, entryName, entryCPF, entryCEP,
	//		entrySalario, entryDataNascimento, entrySexo, entryDepartamentoID)
	//})
	//
	//btnAlterar := widget.NewButton("Alterar funcionários", func() {
	//	depID, err := strconv.Atoi(entryDepartamentoID.Text)
	//	if err != nil {
	//		fmt.Println("Erro ao converter ID do departamento:", err)
	//		return
	//	} else {
	//		entryDepartamentoID.SetText(strconv.Itoa(depID))
	//	}
	//	funcionario := models.Funcionario{
	//		ID:             entryID.Text,
	//		Nome:           entryName.Text,
	//		CPF:            entryCPF.Text,
	//		CEP:            entryCEP.Text,
	//		Salario:        entrySalario.Text,
	//		DataNascimento: entryDataNascimento.Text,
	//		Sexo:           entrySexo.Text,
	//		DepartamentoID: depID,
	//	}
	//	lbResultado.SetText("Resultado: " + funcionario.Atualizar())
	//	utils.LimparCampos(entryID, entryName, entryCPF, entryCEP,
	//		entrySalario, entryDataNascimento, entrySexo, entryDepartamentoID)
	//})
	//
	//btnListar := widget.NewButton("Listar Funcionários", func() {
	//	funcionarios, err := models.CarregarFuncionarios()
	//	if err != nil {
	//		fmt.Println("Erro:", err)
	//		return
	//	}
	//
	//	listaFuncionarios.RemoveAll()
	//
	//	for _, f := range funcionarios {
	//		card := widget.NewCard(f.Nome,
	//			"",
	//			widget.NewLabel(fmt.Sprintf("ID: %s\nDepartamento: %d\nCPF: %s\nCEP: %s\nSalário: %s\nNascimento: %s\nSexo: %s",
	//				f.ID, f.DepartamentoID, f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo)))
	//		listaFuncionarios.Add(card)
	//	}
	//	employeesListPage.Show()
	//
	//	listaFuncionarios.Refresh()
	//	employeesListPage.SetCloseIntercept(func() {
	//		employeesListPage.Hide()
	//	})
	//})
	//
	//btnDeletar := widget.NewButton("Deletar funcionário", func() {
	//	funcionarioID := entryID.Text
	//	if funcionarioID == "" {
	//		lbResultado.SetText("Erro: ID do funcionário não pode ser vazio.")
	//		return
	//	}
	//
	//	funcionario := models.Funcionario{
	//		ID:             funcionarioID,
	//		Nome:           entryName.Text,
	//		CPF:            entryCPF.Text,
	//		CEP:            entryCEP.Text,
	//		Salario:        entrySalario.Text,
	//		DataNascimento: entryDataNascimento.Text,
	//		Sexo:           entrySexo.Text,
	//		DepartamentoID: 0,
	//	}
	//	lbResultado.SetText("Resultado: " + funcionario.Deletar())
	//	utils.LimparCampos(entryID, entryName, entryCPF, entryCEP,
	//		entrySalario, entryDataNascimento, entrySexo, entryDepartamentoID)
	//})
	//
	//employeesListPage.SetContent(container.NewScroll(
	//	listaFuncionarios,
	//))
	//
	//rodape := widget.NewButton("Voltar", func() {
	//	employeesMainPage.Hide()
	//	employeesListPage.Hide()
	//	mainPage.Show()
	//})
	//
	//employeesMainPage.SetContent(container.NewBorder(
	//	container.NewVBox(
	//		entryID,
	//		entryName,
	//		entryCPF,
	//		entryCEP,
	//		entrySalario,
	//		entryDataNascimento,
	//		entrySexo,
	//		entryDepartamentoID,
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
	//employeesMainPage.SetCloseIntercept(func() {
	//	employeesMainPage.Close()
	//	mainPage.Show()
	//})
	//
	//employeesMainPage.Show()
}
