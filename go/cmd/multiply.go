package cmd

import (
	"syscall/js"
)

func Multiply(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	return a * b
}
