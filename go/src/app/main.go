package main

import (
	"fmt"
	"go/src/cmd"
)

func main() {
	var a, b int
	var mark string
	fmt.Scanln(&a, &mark, &b)

	switch mark {
	case "+":
		fmt.Println(cmd.Add(a, b))
	case "-":
		fmt.Println(cmd.Sub(a, b))
	case "*":
		fmt.Println(cmd.Mul(a, b))
	case "/":
		fmt.Println(cmd.Div(a, b))
	}
}
