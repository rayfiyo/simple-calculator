package display

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Clear(this js.Value, args []js.Value) interface{} {
	model.Base.Display = "0"
	model.Base.Operation = ""
	model.Log.Display = ""

	return nil
}
