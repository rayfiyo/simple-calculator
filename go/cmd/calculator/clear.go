package calculator

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Clear(this js.Value, args []js.Value) interface{} {
	model.Input.Number = "0"
	model.Input.Operation = ""
	model.Log.Number = ""
	model.Log.Operation = ""
	return nil
}
