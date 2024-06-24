package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("multiply", js.FuncOf(multiply))
	js.Global().Set("divide", js.FuncOf(divide))
	<-c
}

func add(this js.Value, args []js.Value) interface{} {
	a, _ := strconv.Atoi(args[0].String())
	b, _ := strconv.Atoi(args[1].String())
	return a + b
}

func subtract(this js.Value, args []js.Value) interface{} {
	a, _ := strconv.Atoi(args[0].String())
	b, _ := strconv.Atoi(args[1].String())
	return a - b
}

func multiply(this js.Value, args []js.Value) interface{} {
	a, _ := strconv.Atoi(args[0].String())
	b, _ := strconv.Atoi(args[1].String())
	return a * b
}

func divide(this js.Value, args []js.Value) interface{} {
	a, _ := strconv.Atoi(args[0].String())
	b, _ := strconv.Atoi(args[1].String())
	if b == 0 {
		return "Error: Division by zero"
	}
	return a / b
}
