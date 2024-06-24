package main

import (
	calc "simple-calculator/go/cmd/calculator"
	"simple-calculator/go/cmd/get"
	"simple-calculator/go/cmd/join"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	js.Global().Set("appendNumber", js.FuncOf(join.Number))
	js.Global().Set("appendOperation", js.FuncOf(join.Operation))
	js.Global().Set("calculate", js.FuncOf(calc.Calculate))
	js.Global().Set("clearCalculator", js.FuncOf(calc.Clear))
	js.Global().Set("getInput", js.FuncOf(get.Input))
	js.Global().Set("getLog", js.FuncOf(get.Log))

	<-c
}
