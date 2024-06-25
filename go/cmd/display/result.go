package display

import (
	"simple-calculator/go/model"
	"strconv"
)

func Result(result int) {
	model.Base.Number = strconv.Itoa(result)
	model.Base.Operation = "="
	model.Log.Number = ""
	model.Log.Operation = ""
}
