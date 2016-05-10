package pkg2

import (
	"errors"
)

func SumInt(nums ...int) int {
	var sum int

	for _, v := range nums {
		sum += v
	}

	return sum
}

func DivInt(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为 0 ")
	}

	return a / b, nil
}
