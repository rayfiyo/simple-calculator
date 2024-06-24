package main

import (
	"simple-calculator/go/cmd"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("add", js.FuncOf(cmd.Add))
	js.Global().Set("subtract", js.FuncOf(cmd.Subtract))
	js.Global().Set("multiply", js.FuncOf(cmd.Multiply))
	js.Global().Set("divide", js.FuncOf(cmd.Divide))
	<-c
}
