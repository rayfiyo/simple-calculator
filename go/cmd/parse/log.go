package parse

import (
	"fmt"
	"simple-calculator/go/model"
	"strconv"
)

func Log() (int, error) {
	num, err := strconv.Atoi(model.Log.Number[:len(model.Log.Number)-2])
	if err != nil {
		return num, err
	}

	if num > 9999 {
		return num, fmt.Errorf("large digits of log")
	}

	return num, nil
}
