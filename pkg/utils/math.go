package utils

import "math"

func SafeDiv(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
