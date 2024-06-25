package display

import (
	"fmt"
	"simple-calculator/go/model"
)

func Error(err interface{}) interface{} {
	model.Base.Number = "err:" + fmt.Sprint(err)
	model.Base.Operation = ""
	model.Log.Number = ""
	model.Log.Operation = ""

	return nil
}
