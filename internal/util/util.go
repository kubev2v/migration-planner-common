package util

import "math"

// Round rounds a float64 to 2 decimal places
func Round(f float64) float64 {
	return math.Round(f*100) / 100
}
