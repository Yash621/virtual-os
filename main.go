package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	historyBtn := widget.NewButton("history", func() {
		w.SetContent(widget.NewLabel("history"))
	})
	backBtn := widget.NewButton("back", func() {})
	clearBtn := widget.NewButton("clear", func() {})
	openBtn := widget.NewButton("open", func() {})
	closeBtn := widget.NewButton("close", func() {})
	divideBtn := widget.NewButton("divide", func() {})
	sevenBtn := widget.NewButton("7", func() {})
	eightBtn := widget.NewButton("8", func() {})
	nineBtn := widget.NewButton("9", func() {})
	multiplyBtn := widget.NewButton("multiply", func() {})
	fourBtn := widget.NewButton("4", func() {})
	fiveBtn := widget.NewButton("5", func() {})
	sixBtn := widget.NewButton("6", func() {})
	subtractBtn := widget.NewButton("subtract", func() {})
	oneBtn := widget.NewButton("1", func() {})
	twoBtn := widget.NewButton("2", func() {})
	threeBtn := widget.NewButton("3", func() {})
	addBtn := widget.NewButton("add", func() {})
	zeroBtn := widget.NewButton("0", func() {})
	decimalBtn := widget.NewButton(".", func() {})
	equalsBtn := widget.NewButton("=", func() {})
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
