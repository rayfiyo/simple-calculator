package get

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Input(this js.Value, args []js.Value) interface{} {
	return model.Input.Number
}
