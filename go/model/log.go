package model

type log struct {
	Number    string
	Operation string
}

var Log = &log{
	Number:    "",
	Operation: "",
}
