package parse

import (
	"fmt"
	"simple-calculator/go/model"
	"strconv"
)

func Log() (int, error) {
	num, err := strconv.Atoi(model.Log.Display[:len(model.Log.Display)-2])
	if err != nil {
		return num, err
	}

	if num > +3037000499 {
		return num, fmt.Errorf("positive overflow (log)")
	}
	if num < -3037000500 {
		return num, fmt.Errorf("negative overflow (log)")
	}

	return num, nil
}
