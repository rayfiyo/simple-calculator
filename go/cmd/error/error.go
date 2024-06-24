package error

import (
	"fmt"
	"simple-calculator/go/model"
)

func Print(err interface{}) {
	model.Input.Number = "err:" + fmt.Sprint(err)
	model.Input.Operation = ""
	model.Log.Number = ""
	model.Log.Operation = ""
}
