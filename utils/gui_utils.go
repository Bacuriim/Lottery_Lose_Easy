package utils

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func LimparCampos(entries ...*widget.Entry) {
	for _, entry := range entries {
		entry.SetText("")
		entry.Refresh()
	}
}

func CriarEntryLetras(valor string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Digite o " + valor + " aqui")
	entry.Validator = validation.NewRegexp("^[a-zA-Z]*$", "Somente letras são permitidas")
	return entry
}

func CriarEntryNumeros(valor string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Digite o " + valor + " aqui")
	if valor == "CPF" {
		entry.Validator = validation.NewRegexp("^[0-9]{3}.[0-9]{3}.[0-9]{3}-[0-9]{2}$", "Somente números são permitidos e deve ter 11 dígitos")
	} else if valor == "CEP" {
		entry.Validator = validation.NewRegexp("^[0-9]{5}-[0-9]{3}$", "Somente números são permitidos e deve ter 8 dígitos")
	} else {
		entry.Validator = validation.NewRegexp("^[0-9]*$", "Somente números são permitidos")
	}
	return entry
}

func CriarEntryLetrasNumeros(valor string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Digite o " + valor + " aqui")
	entry.Validator = validation.NewRegexp("^[a-zA-Z0-9]*$", "Somente letras e números são permitidos")
	return entry
}

func CriarEntryData(valor string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Digite a " + valor + " aqui")
	entry.Validator = validation.NewRegexp("^[0-9]{2}/[0-9]{2}/[0-9]{4}$", "Formato de data inválido. Use DD-MM-AAAA")
	return entry
}

func FormatarData(data string) string {
	if len(data) != 10 {
		return data
	}
	return data[6:10] + "/" + data[3:5] + "/" + data[0:2]
}

func FormatarCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf
	}
	return cpf[0:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:11]
}

func FormatarCEP(cep string) string {
	if len(cep) != 8 {
		return cep
	}
	return cep[0:5] + "-" + cep[5:8]
}
