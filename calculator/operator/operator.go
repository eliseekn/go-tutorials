package operator

import "calculator/utils"

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

func Divide(args []string) float32 {
	result := utils.Float(args[1]) / utils.Float(args[3])
	return result
}
