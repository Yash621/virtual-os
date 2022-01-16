package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	input := widget.NewLabel("")
	output := ""
	historyBtn := widget.NewButton("History", func() {
		w.SetContent(widget.NewLabel("history"))
	})
	backBtn := widget.NewButton("Back", func() {})
	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	openBtn := widget.NewButton("(", func() {
		output += "("
		input.SetText(output)
	})
	closeBtn := widget.NewButton(")", func() {
		output += ")"
		input.SetText(output)
	})
	divideBtn := widget.NewButton("/", func() {
		output += "/"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output += "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output += "8"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output += "9"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output += "*"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output += "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output += "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output += "6"
		input.SetText(output)
	})
	subtractBtn := widget.NewButton("-", func() {
		output += "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output += "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output += "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output += "3"
		input.SetText(output)
	})
	addBtn := widget.NewButton("+", func() {
		output += "+"
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output += "0"
		input.SetText(output)
	})
	decimalBtn := widget.NewButton(".", func() {
		output += "."
		input.SetText(output)
	})
	equalsBtn := widget.NewButton("=", func() {
		output += "="
		input.SetText(output)
	})

	w.SetContent(container.NewVBox(input, container.NewGridWithColumns(1,
		container.NewGridWithColumns(2, historyBtn, backBtn),
		container.NewGridWithColumns(4, clearBtn, openBtn, closeBtn, divideBtn),
		container.NewGridWithColumns(4, nineBtn, eightBtn, sevenBtn, multiplyBtn),
		container.NewGridWithColumns(4, fourBtn, fiveBtn, sixBtn, subtractBtn),
		container.NewGridWithColumns(4, twoBtn, oneBtn, threeBtn, addBtn),
		container.NewGridWithColumns(2, container.NewGridWithColumns(2, zeroBtn, decimalBtn), equalsBtn),
	)))
	w.ShowAndRun()
}
