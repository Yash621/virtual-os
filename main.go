package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	input := widget.NewLabel("")
	output := ""
	historyStr := ""
	var historyArr []string
	ans := ""
	history := widget.NewLabel(historyStr)
	isHistory := true
	historyBtn := widget.NewButton("History", func() {
		if isHistory {
			historyStr = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyStr += historyArr[i] + "\n"
			}

		}
		isHistory = !isHistory
		history.SetText(historyStr)
	})
	backBtn := widget.NewButton("Back", func() {
		if output != "" {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
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
		expression, err := govaluate.NewEvaluableExpression(output)
		if err != nil {
			output = "Incorrect Expression Please try again"
			input.SetText(output)
		} else {
			result, err := expression.Evaluate(nil)
			if err != nil {
				output = "Error"
			} else {
				ans = strconv.FormatFloat(result.(float64), 'f', -1, 64)
				historyArr = append(historyArr, output+"="+ans)
				output = ans
			}
			input.SetText(output)
		}

	})

	w.SetContent(container.NewVBox(input, history, container.NewGridWithColumns(1,
		container.NewGridWithColumns(2, historyBtn, backBtn),
		container.NewGridWithColumns(4, clearBtn, openBtn, closeBtn, divideBtn),
		container.NewGridWithColumns(4, nineBtn, eightBtn, sevenBtn, multiplyBtn),
		container.NewGridWithColumns(4, fourBtn, fiveBtn, sixBtn, subtractBtn),
		container.NewGridWithColumns(4, twoBtn, oneBtn, threeBtn, addBtn),
		container.NewGridWithColumns(2, container.NewGridWithColumns(2, zeroBtn, decimalBtn), equalsBtn),
	)))
	w.ShowAndRun()
}
