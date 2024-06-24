package calculator

import (
	"simple-calculator/go/model"
	"strconv"
	"syscall/js"
)

func Calculate(this js.Value, args []js.Value) interface{} {
	if model.Input.Operation == "" || model.Log.Number == "" {
		return nil
	}

	prev, _ := strconv.Atoi(model.Log.Number[:len(model.Log.Number)-2])
	current, _ := strconv.Atoi(model.Input.Number)

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
			model.Input.Number = "Error"
			model.Input.Operation = ""
			model.Log.Number = ""
			model.Log.Operation = ""
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
