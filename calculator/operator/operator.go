package operator

import (
	"calculator/utils"
	"errors"
)

func Sum(args []string) float32 {
	result := utils.Float(args[1]) + utils.Float(args[3])
	return result
}

func Substract(args []string) float32 {
	result := utils.Float(args[1]) - utils.Float(args[3])
	return result
}

func Multiply(args []string) float32 {
	result := utils.Float(args[1]) * utils.Float(args[3])
	return result
}

func Divide(args []string) (float32, error) {
	number1 := utils.Float(args[1])
	number2 := utils.Float(args[3])

	if number2 == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	result := number1 / number2
	return result, nil
}
