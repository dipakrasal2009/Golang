package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	//"fyne.io/fyne/v2/theme"
	"strconv"

	"fyne.io/fyne/v2"

	//"fine.io/fyne/v2/fyne"

	"github.com/Knetic/govaluate"
)

var input string
var result string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Enter numbers")
	inputEntry.Disable()

	// Function to handle button clicks
	updateDisplay := func(text string) {
		input += text
		inputEntry.SetText(input)
	}

	// Function to handle the result
	calculateResult := func() {
		if res, err := eval(input); err == nil {
			result = strconv.FormatFloat(res, 'f', -1, 64)
			inputEntry.SetText(result)
			input = result // To continue calculations
		} else {
			inputEntry.SetText("Error")
			input = ""
		}
	}

	// Creating buttons
	buttons := []string{
		"7", "8", "9", "/", "C",
		"4", "5", "6", "*", "(",
		"1", "2", "3", "-", ")",
		"0", ".", "=", "+", "CE",
	}

	grid := container.NewGridWithColumns(5)
	for _, btn := range buttons {
		text := btn
		button := widget.NewButton(text, func() {
			if text == "C" {
				input = ""
				inputEntry.SetText("")
			} else if text == "=" {
				calculateResult()
			} else if text == "CE" {
				if len(input) > 0 {
					input = input[:len(input)-1]
					inputEntry.SetText(input)
				}
			} else {
				updateDisplay(text)
			}
		})
		grid.Add(button)
	}

	myWindow.SetContent(container.NewVBox(
		inputEntry,
		grid,
	))

	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

// Simple evaluator for arithmetic expressions
func eval(expression string) (float64, error) {
	// Use a third-party library or write a simple parser to evaluate the expression
	// For simplicity, let's use a package that evaluates expressions

	exp, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	result, err := exp.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	return result.(float64), nil
}
