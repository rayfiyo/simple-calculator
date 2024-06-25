package display

import (
	"fmt"
	"simple-calculator/go/model"
)

func Error(err interface{}) {
	model.Base.Display = fmt.Sprint(err)
	model.Base.Operation = ""
	model.Log.Display = "Error"
}
