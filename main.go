package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	historyBtn := widget.NewButton("History", func() {
		w.SetContent(widget.NewLabel("history"))
	})
	backBtn := widget.NewButton("Back", func() {})
	clearBtn := widget.NewButton("Clear", func() {})
	openBtn := widget.NewButton("(", func() {})
	closeBtn := widget.NewButton(")", func() {})
	divideBtn := widget.NewButton("/", func() {})
	sevenBtn := widget.NewButton("7", func() {})
	eightBtn := widget.NewButton("8", func() {})
	nineBtn := widget.NewButton("9", func() {})
	multiplyBtn := widget.NewButton("*", func() {})
	fourBtn := widget.NewButton("4", func() {})
	fiveBtn := widget.NewButton("5", func() {})
	sixBtn := widget.NewButton("6", func() {})
	subtractBtn := widget.NewButton("-", func() {})
	oneBtn := widget.NewButton("1", func() {})
	twoBtn := widget.NewButton("2", func() {})
	threeBtn := widget.NewButton("3", func() {})
	addBtn := widget.NewButton("+", func() {})
	zeroBtn := widget.NewButton("0", func() {})
	decimalBtn := widget.NewButton(".", func() {})
	equalsBtn := widget.NewButton("=", func() {})
	hello := widget.NewLabel("Hello")
	w.SetContent(container.NewVBox(hello, container.NewGridWithColumns(1,
		container.NewGridWithColumns(2, historyBtn, backBtn),
		container.NewGridWithColumns(4, clearBtn, openBtn, closeBtn, divideBtn),
		container.NewGridWithColumns(4, nineBtn, eightBtn, sevenBtn, multiplyBtn),
		container.NewGridWithColumns(4, fourBtn, fiveBtn, sixBtn, subtractBtn),
		container.NewGridWithColumns(4, twoBtn, oneBtn, threeBtn, addBtn),
		container.NewGridWithColumns(2, container.NewGridWithColumns(2, zeroBtn, decimalBtn), equalsBtn),
	)))
	w.ShowAndRun()
}
