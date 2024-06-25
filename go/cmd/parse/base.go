package parse

import (
	"fmt"
	"simple-calculator/go/model"
	"strconv"
)

func Base() (int, error) {
	num, err := strconv.Atoi(model.Base.Number)
	if err != nil {
		return num, err
	}

	if num > 9999 {
		return num, fmt.Errorf("large digits of base")
	}

	return num, nil
}
