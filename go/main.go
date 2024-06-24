package main

import (
	"strconv"
	"syscall/js"
)

type Calculator struct {
	currentInput     string
	previousInput    string
	currentOperation string
}

var calc = &Calculator{
	currentInput:     "0",
	previousInput:    "",
	currentOperation: "",
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("appendNumber", js.FuncOf(appendNumber))
	js.Global().Set("setOperation", js.FuncOf(setOperation))
	js.Global().Set("calculate", js.FuncOf(calculate))
	js.Global().Set("clearCalculator", js.FuncOf(clearCalculator))
	js.Global().Set("getCurrentInput", js.FuncOf(getCurrentInput))
	js.Global().Set("getPreviousInput", js.FuncOf(getPreviousInput))

	<-c
}

func appendNumber(this js.Value, args []js.Value) interface{} {
	number := args[0].String()
	if calc.currentInput == "0" {
		calc.currentInput = number
	} else {
		calc.currentInput += number
	}
	return nil
}

func setOperation(this js.Value, args []js.Value) interface{} {
	operation := args[0].String()
	if calc.currentOperation != "" {
		calculate(this, nil)
	}
	calc.currentOperation = operation
	calc.previousInput = calc.currentInput + " " + getOperationSymbol(operation)
	calc.currentInput = "0"
	return nil
}

func calculate(this js.Value, args []js.Value) interface{} {
	if calc.currentOperation == "" || calc.previousInput == "" {
		return nil
	}

	prev, _ := strconv.Atoi(calc.previousInput[:len(calc.previousInput)-2])
	current, _ := strconv.Atoi(calc.currentInput)

	var result int
	switch calc.currentOperation {
	case "add":
		result = prev + current
	case "subtract":
		result = prev - current
	case "multiply":
		result = prev * current
	case "divide":
		if current == 0 {
			calc.currentInput = "Error"
			calc.previousInput = ""
			calc.currentOperation = ""
			return nil
		}
		result = prev / current
	}

	calc.currentInput = strconv.Itoa(result)
	calc.previousInput = ""
	calc.currentOperation = ""
	return nil
}

func clearCalculator(this js.Value, args []js.Value) interface{} {
	calc.currentInput = "0"
	calc.previousInput = ""
	calc.currentOperation = ""
	return nil
}

func getCurrentInput(this js.Value, args []js.Value) interface{} {
	return calc.currentInput
}

func getPreviousInput(this js.Value, args []js.Value) interface{} {
	return calc.previousInput
}

func getOperationSymbol(operation string) string {
	switch operation {
	case "add":
		return "+"
	case "subtract":
		return "-"
	case "multiply":
		return "×"
	case "divide":
		return "÷"
	default:
		return ""
	}
}
