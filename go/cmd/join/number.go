package join

import (
	"simple-calculator/go/model"
	"strings"
	"syscall/js"
)

func Number(this js.Value, args []js.Value) interface{} {
	number := args[0].String()

	if strings.Contains(model.Base.Display, "err:") {
		model.Base.Display = ""
	}

	switch model.Base.Display {
	case "Simple Calculator":
		model.Base.Display = number
	case "0":
		model.Base.Display = number
	default:
		model.Base.Display += number
	}

	return nil
}
