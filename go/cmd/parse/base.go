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

	if num > +3037000499 {
		return num, fmt.Errorf("positive overflow (base) %d", num)
	}
	if num < -3037000500 {
		return num, fmt.Errorf("negative overflow (base)")
	}

	return num, nil
}
