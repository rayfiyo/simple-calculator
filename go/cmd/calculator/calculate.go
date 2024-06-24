package calculator

import (
	"simple-calculator/go/cmd/error"
	"simple-calculator/go/model"
	"strconv"
	"syscall/js"
)

func Calculate(this js.Value, args []js.Value) interface{} {
	if model.Input.Operation == "" || model.Log.Number == "" {
		return nil
	}

	prev, err := strconv.Atoi(model.Log.Number[:len(model.Log.Number)-2])
	if err != nil {
		error.Print(err)
		return nil
	}
	if prev > 9999 {
		error.Print("Large digits prev")
		return nil
	}

	current, err := strconv.Atoi(model.Input.Number)
	if err != nil {
		error.Print(err)
		return nil
	}
	if current > 9999 {
		error.Print("Large digits current")
		return nil
	}

	var result int
	switch model.Input.Operation {
	case "add":
		result = prev + current
	case "subtract":
		result = prev - current
	case "multiply":
		result = prev * current
	case "divide":
		if current == 0 {
			error.Print("Division by zero")
			return nil
		}
		result = prev / current
	}

	model.Input.Number = strconv.Itoa(result)
	model.Input.Operation = ""
	model.Log.Number = ""
	model.Log.Operation = ""
	return nil
}
