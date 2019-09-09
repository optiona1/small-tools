package round

import "math"

func RoundFixed(num float64, precision int) float64 {
	output := math.Pow10(precision)
	return float64(math.Round(num * output)) / output
}
