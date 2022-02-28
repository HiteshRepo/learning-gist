package utils

import (
	"math"
)

func CheckFloatEquals(x, y float64) bool {
	tolerance := 0.0000001
	diff := math.Abs(x - y)
	mean := math.Abs(x + y) / 2.0
	factor := diff / mean
	return factor < tolerance
}
