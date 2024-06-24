package join

import (
	calc "simple-calculator/go/cmd/calculator"
	"simple-calculator/go/cmd/get"
	"simple-calculator/go/model"
	"syscall/js"
)

func Operation(this js.Value, args []js.Value) interface{} {
	operation := args[0].String()

	if model.Input.Operation != "" {
		calc.Calculate(this, nil)
	}

	model.Input.Operation = operation
	model.Log.Number = model.Input.Number + " " + get.OperationSymbol(operation)
	model.Input.Number = "0"
	return nil
}
