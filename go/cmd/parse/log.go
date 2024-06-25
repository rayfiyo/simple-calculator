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

	if num > 4294967296-1 {
		return num, fmt.Errorf("too big (log)")
	}

	return num, nil
}
