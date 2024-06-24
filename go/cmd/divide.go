package cmd

import (
	"fmt"
	"syscall/js"
)

func Divide(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	if b == 0 {
		return fmt.Errorf("error: Division by zero")
	}
	return a / b
}
