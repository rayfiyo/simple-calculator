package model

type input struct {
	Number    string
	Operation string
}

var Input = &input{
	Number:    "Simple Calculator",
	Operation: "",
}

type log struct {
	Number    string
	Operation string
}

var Log = &log{
	Number:    "",
	Operation: "",
}
