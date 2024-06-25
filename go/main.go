package main

import (
	calc "simple-calculator/go/cmd/calculate"
	"simple-calculator/go/cmd/display"
	"simple-calculator/go/cmd/get"
	"simple-calculator/go/cmd/join"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	js.Global().Set("joinNumber", js.FuncOf(join.Number))
	js.Global().Set("joinOperation", js.FuncOf(join.Operation))

	js.Global().Set("calculate", js.FuncOf(calc.Do))

	js.Global().Set("clearCalculator", js.FuncOf(display.Clear))

	js.Global().Set("getBase", js.FuncOf(get.Base))
	js.Global().Set("getLog", js.FuncOf(get.Log))

	<-c
}
