package main

import (
	"errors"
)

func main() {
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}
