package join

import (
	calc "simple-calculator/go/cmd/calculate"
	"simple-calculator/go/cmd/display"
	"simple-calculator/go/cmd/get"
	"simple-calculator/go/model"
	"strings"
	"syscall/js"
)

func Operation(this js.Value, args []js.Value) interface{} {
	operation := args[0].String()

	if strings.Contains(model.Log.Display, "Error") {
		display.Clear(this, args)
		return nil
	}

	if model.Base.Operation != "" {
		calc.Do(this, nil)
		return nil
	}

	model.Base.Operation = operation
	model.Log.Display = model.Base.Display + " " + get.OperationSymbol(operation)
	model.Base.Display = "0"

	return nil
}
