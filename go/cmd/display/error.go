package display

import (
	"fmt"
	"simple-calculator/go/model"
)

func Error(err interface{}) interface{} {
	model.Base.Display = "err:" + fmt.Sprint(err)
	model.Base.Operation = ""
	model.Log.Display = ""

	return nil
}
