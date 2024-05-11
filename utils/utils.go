package utils

import "math"

func FloatToInteger_OneDecimal(num float64) int {
	return int(math.Round(num * 10))
}

func FloatToInteger_TwoDecimals(num float64) int {
	return int(math.Round(num * 100))
}

func DegreesToInteger(degree float64) int {
	return int(math.Round(degree * math.Pow(10, 7)))
}

func IntegerToFloat_OneDecimal(num int) float64 {
	return float64(num) / 10
}

func IntegerToFloat_TwoDecimals(num int) float64 {
	return float64(num) / 100
}

func DegreesToFloat(degree int) float64 {
	return float64(degree) / math.Pow(10, 7)
}

func DegreesMultiOnePointFive(degree float64) int {
	return int(math.Round(degree * 1.5))
}

func DegreesDivideOnePointFive(degree int) float64 {
	return float64(degree) / 1.5
}
