package get

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Base(this js.Value, args []js.Value) interface{} {
	return model.Base.Number
}
