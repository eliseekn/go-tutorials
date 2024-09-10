package utils

import "strconv"

func Float(str string) float32 {
	result, err := strconv.ParseFloat(str, 32)
	if err != nil {
		panic(err)
	}

	return float32(result)
}
