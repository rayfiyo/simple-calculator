package calculate

import (
	"simple-calculator/go/cmd/display"
	"simple-calculator/go/cmd/parse"
	"simple-calculator/go/model"
	"syscall/js"
)

func Do(this js.Value, args []js.Value) interface{} {
	if model.Base.Display == "" || model.Log.Display == "" {
		return nil
	}

	base, err := parse.Base()
	if err != nil {
		display.Error(err)
		return nil
	}

	log, err := parse.Log()
	if err != nil {
		display.Error(err)
		return nil
	}

	var result model.Result
	switch model.Base.Operation {
	case "add":
		result.Value = log + base
		result.Log = log
		result.Base = base
		result.Operation = "+"
	case "subtract":
		result.Value = log - base
		result.Log = log
		result.Base = base
		result.Operation = "-"
	case "multiply":
		result.Value = log * base
		result.Log = log
		result.Base = base
		result.Operation = "*"
	case "divide":
		if base == 0 {
			display.Error("Division by zero")
			return nil
		}
		result.Value = log / base
		result.Log = log
		result.Base = base
		result.Operation = "/"
	}

	display.Result(result)

	return nil
}
