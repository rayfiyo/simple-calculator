package calculate

import (
	"simple-calculator/go/cmd/display"
	"simple-calculator/go/cmd/parse"
	"simple-calculator/go/model"
	"syscall/js"
)

func Do(this js.Value, args []js.Value) interface{} {
	if model.Base.Operation == "" || model.Log.Number == "" {
		return nil
	}

	base, err := parse.Base()
	if err != nil {
		display.Error(err)
	}

	log, err := parse.Log()
	if err != nil {
		display.Error(err)
	}

	var result int
	switch model.Base.Operation {
	case "add":
		result = base + log
	case "subtract":
		result = base - log
	case "multiply":
		result = base * log
	case "divide":
		if log == 0 {
			display.Error("Division by zero")
			return nil
		}
		result = base / log
	}

	display.Result(result)

	return nil
}
