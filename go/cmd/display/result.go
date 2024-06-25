package display

import (
	"simple-calculator/go/model"
	"strconv"
)

func Result(result model.Result) {
	model.Base.Display = strconv.Itoa(result.Value)
	model.Base.Operation = ""
	model.Log.Display = strconv.Itoa(result.Log) + " " + result.Operation + " " + strconv.Itoa(result.Base) + " ="
}
