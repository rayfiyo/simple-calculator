package display

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Clear(this js.Value, args []js.Value) interface{} {
	model.Base.Number = "0"
	model.Base.Operation = ""
	model.Log.Number = ""
	model.Log.Operation = ""
	return nil
}
