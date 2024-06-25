package model

type base struct {
	Display   string
	Operation string
}

var Base = &base{
	Display:   "Simple Calculator",
	Operation: "",
}
