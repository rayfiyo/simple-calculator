package join

import (
	"simple-calculator/go/model"
	"syscall/js"
)

func Number(this js.Value, args []js.Value) interface{} {
	number := args[0].String()

	if model.Input.Number == "Simple Calculator" {
		model.Input.Number = ""
	}

	if model.Input.Number == "0" {
		model.Input.Number = number
	} else {
		model.Input.Number += number
	}
	return nil
}
