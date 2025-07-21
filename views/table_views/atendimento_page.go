package table_views

import (
	"fmt"
	"lottery-lose-easy/models"
	"lottery-lose-easy/utils"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func AtendimentoPage(myApp fyne.App, mainPage fyne.Window) {
	win := myApp.NewWindow("Atendimentos")
	win.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/LotericaPercaFacil.png")
	if err == nil {
		win.SetIcon(icon)
	}

	entryId := utils.CriarEntryNumeros("ID")
	entryTransacao := utils.CriarEntryLetrasNumeros("Número Transação")
	entryData := utils.CriarEntryLetrasNumeros("Data/Hora (YYYY-MM-DD HH:MM:SS)")
	entryClienteCpf := utils.CriarEntryNumeros("CPF do Cliente")
	entryFuncionarioCpf := utils.CriarEntryNumeros("CPF do Funcionário")
	entryServicoId := utils.CriarEntryNumeros("ID Serviço")
	entryMeioPagamento := utils.CriarEntryLetrasNumeros("Meio de Pagamento")

	meiosPagamento := []string{"dinheiro", "boleto", "cartão", "cheque", "PIX", "TED"}
	selectMeioPagamento := widget.NewSelect(meiosPagamento, func(string) {})
	selectMeioPagamento.PlaceHolder = "Selecione o meio de pagamento"

	resultLabel := widget.NewLabel("")

	addTab := container.NewVBox(
		widget.NewLabel("Adicionar Atendimento"),
		entryId, entryData, entryClienteCpf, entryFuncionarioCpf, entryServicoId, selectMeioPagamento,
		widget.NewButton("Criar atendimento", func() {
			id, _ := strconv.Atoi(entryId.Text)
			servicoId, _ := strconv.Atoi(entryServicoId.Text)
			clienteUUIDstr, err := utils.BuscarClienteUUIDPorCPF(entryClienteCpf.Text)
			if err != nil {
				resultLabel.SetText("Cliente não encontrado: " + err.Error())
				return
			}
			clienteId, _ := uuid.Parse(clienteUUIDstr)

			funcionarioUUIDstr, err := utils.BuscarFuncionarioUUIDPorCPF(entryFuncionarioCpf.Text)
			if err != nil {
				resultLabel.SetText("Funcionário não encontrado: " + err.Error())
				return
			}
			funcionarioId, _ := uuid.Parse(funcionarioUUIDstr)

			dataHora, _ := time.Parse("2006-01-02 15:04:05", entryData.Text)
			numeroTransacao := uuid.New().String() // Gera automaticamente

			atendimento := models.Atendimento{
				Id:              int32(id),
				NumeroTransacao: numeroTransacao,
				DataHora:        dataHora,
				ClienteId:       clienteId,
				FuncionarioId:   funcionarioId,
				ServicoId:       int32(servicoId),
				MeioPagamento:   selectMeioPagamento.Selected,
			}
			msg := atendimento.Salvar(atendimento)
			resultLabel.SetText(msg)
		}),
		resultLabel,
	)

	atendList := container.NewVBox()
	listTab := container.NewVBox(
		widget.NewLabel("Atendimentos cadastrados:"),
		atendList,
		widget.NewButton("Listar Atendimentos", func() {
			atendList.Objects = nil
			atendimentos, msg := new(models.Atendimento).BuscarTodos()
			if msg != "Atendimentos encontrados!" {
				resultLabel.SetText(msg)
			} else {
				for _, a := range atendimentos {
					card := widget.NewCard(
						fmt.Sprintf("Transação: %s", a.NumeroTransacao),
						fmt.Sprintf("ID: %d", a.Id),
						widget.NewLabel(fmt.Sprintf(
							"Data/Hora: %s\nCliente: %s\nFuncionário: %s\nServiço: %d\nMeio Pagamento: %s",
							a.DataHora.Format("2006-01-02 15:04:05"), a.ClienteId, a.FuncionarioId, a.ServicoId, a.MeioPagamento,
						)),
					)
					atendList.Add(card)
				}
				resultLabel.SetText("Atendimentos listados com sucesso!")
			}
			atendList.Refresh()
		}),
		resultLabel,
	)

	searchResult := widget.NewLabel("Resultado: nenhum")
	searchId := utils.CriarEntryNumeros("ID")

	searchTab := container.NewVBox(
		widget.NewLabel("Buscar Atendimento por ID"),
		searchId,
		widget.NewButton("Buscar", func() {
			val := binding.NewString()
			val.Set(searchId.Text)
			a, msg := new(models.Atendimento).Pesquisar("id", val, true)
			if a != nil {
				searchResult.SetText(fmt.Sprintf("Encontrado: %s, ID: %d", a.NumeroTransacao, a.Id))
			} else {
				searchResult.SetText(msg)
			}
		}),
		searchResult,
	)

	updateId := utils.CriarEntryNumeros("ID")

	updateTab := container.NewVBox(
		widget.NewLabel("Atualizar Atendimento"),
		updateId, entryTransacao, entryData, entryClienteCpf, entryFuncionarioCpf, entryServicoId, entryMeioPagamento,
		widget.NewButton("Atualizar", func() {
			msg := ""
			if entryTransacao.Text != "" {
				msg = new(models.Atendimento).Alterar("numero_transacao", entryTransacao.Text, "id", updateId.Text)
			}
			if entryData.Text != "" {
				dataHora, _ := time.Parse("2006-01-02 15:04:05", entryData.Text)
				msg = new(models.Atendimento).Alterar("data_hora", dataHora, "id", updateId.Text)
			}
			if entryClienteCpf.Text != "" {
				clienteUUIDstr, err := utils.BuscarClienteUUIDPorCPF(entryClienteCpf.Text)
				if err == nil {
					clienteId, _ := uuid.Parse(clienteUUIDstr)
					msg = new(models.Atendimento).Alterar("cliente_id", clienteId, "id", updateId.Text)
				} else {
					msg = "Cliente não encontrado: " + err.Error()
				}
			}
			if entryFuncionarioCpf.Text != "" {
				funcionarioUUIDstr, err := utils.BuscarFuncionarioUUIDPorCPF(entryFuncionarioCpf.Text)
				if err == nil {
					funcionarioId, _ := uuid.Parse(funcionarioUUIDstr)
					msg = new(models.Atendimento).Alterar("funcionario_id", funcionarioId, "id", updateId.Text)
				} else {
					msg = "Funcionário não encontrado: " + err.Error()
				}
			}
			if entryServicoId.Text != "" {
				servicoId, _ := strconv.Atoi(entryServicoId.Text)
				msg = new(models.Atendimento).Alterar("servico_id", int32(servicoId), "id", updateId.Text)
			}
			if entryMeioPagamento.Text != "" {
				msg = new(models.Atendimento).Alterar("meio_pagamento", entryMeioPagamento.Text, "id", updateId.Text)
			}
			resultLabel.SetText(msg)
		}),
		resultLabel,
	)

	removeId := utils.CriarEntryNumeros("ID")

	removeTab := container.NewVBox(
		widget.NewLabel("Remover Atendimento"),
		removeId,
		widget.NewButton("Remover", func() {
			val := binding.NewString()
			val.Set(removeId.Text)
			a, _ := new(models.Atendimento).Pesquisar("id", val, true)
			if a != nil {
				msg := new(models.Atendimento).Remover(*a)
				resultLabel.SetText(msg)
			} else {
				resultLabel.SetText("Atendimento não encontrado para remover.")
			}
		}),
		resultLabel,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Adicionar", addTab),
		container.NewTabItem("Listar", listTab),
		container.NewTabItem("Buscar", searchTab),
		container.NewTabItem("Atualizar", updateTab),
		container.NewTabItem("Remover", removeTab),
	)

	win.SetCloseIntercept(func() {
		win.Hide()
		mainPage.Show()
	})

	tabs.SetTabLocation(container.TabLocationTop)
	win.SetContent(tabs)
	win.Show()
}
