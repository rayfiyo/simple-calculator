package join

import (
	"simple-calculator/go/model"
	"strings"
	"syscall/js"
)

func Number(this js.Value, args []js.Value) interface{} {
	number := args[0].String()

	if strings.Contains(model.Base.Number, "err:") {
		model.Base.Number = ""
	}

	switch model.Base.Number {
	case "Simple Calculator":
		model.Base.Number = ""
	case "0":
		model.Base.Number = number
	default:
		model.Base.Number += number
	}

	return nil
}
