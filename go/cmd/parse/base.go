package parse

import (
	"fmt"
	"simple-calculator/go/model"
	"strconv"
)

func Base() (int, error) {
	num, err := strconv.Atoi(model.Base.Display)
	if err != nil {
		return num, err
	}

	if num > 4294967296-1 {
		return num, fmt.Errorf("too big (base)")
	}

	return num, nil
}
