package math

import "log"

// mean computes the average of a slice (package-private)
func mean(xs []float64) float64 {
	if len(xs) == 0 {
		log.Fatalf("%s | mean of empty slice", moduleName)
	}
	sum := 0.0
	for _, v := range xs {
		sum += v
	}
	return sum / float64(len(xs))
}

// Mean is the exported safe version
func Mean(xs []float64) float64 {
	return mean(xs)
}
