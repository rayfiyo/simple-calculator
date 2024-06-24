package get

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Log(this js.Value, args []js.Value) interface{} {
	return model.Log.Number
}
