package operator

import (
	"calculator/utils"
	"fmt"
)

func Sum(args []string) {
	result := utils.Float(args[1]) + utils.Float(args[3])
	fmt.Println(result)
}

func Substract(args []string) {
	result := utils.Float(args[1]) - utils.Float(args[3])
	fmt.Println(result)
}

func Multiply(args []string) {
	result := utils.Float(args[1]) * utils.Float(args[3])
	fmt.Println(result)
}

func Divide(args []string) {
	result := utils.Float(args[1]) / utils.Float(args[3])
	fmt.Println(result)
}
