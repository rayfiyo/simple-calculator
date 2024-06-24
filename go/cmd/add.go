package cmd

import (
	"syscall/js"
)

func Add(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	return a + b
}
