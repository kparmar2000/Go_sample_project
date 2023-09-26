package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64) (float64, error) {
	defer func() {
		fmt.Println(recover())
	}()
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	fmt.Println(value)
	return math.Sqrt(value), nil
}
func DivideByZero(n1 int, n2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	result := n1 / n2
	return result
}
