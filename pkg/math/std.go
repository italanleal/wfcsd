package math

import (
	"log"
	"math"
)

// std computes the standard deviation of a slice
func std(xs []float64) float64 {
	if len(xs) < 2 {
		log.Fatalf("%s | standard deviation requires at least 2 values", moduleName)
	}
	m := mean(xs)
	sum := 0.0
	for _, v := range xs {
		sum += (v - m) * (v - m)
	}
	return math.Sqrt(sum / float64(len(xs)-1))
}
