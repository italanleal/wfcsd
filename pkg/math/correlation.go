package math

import "log"

// Correlation computes the Pearson correlation coefficient between x and y
func correlation(x, y []float64) float64 {
	if len(x) != len(y) {
		log.Fatalf("%s | length mismatch: %d vs %d", moduleName, len(x), len(y))
	}
	sdx := std(x)
	sdy := std(y)
	if sdx == 0 || sdy == 0 {
		log.Fatalf("%s | zero standard deviation detected", moduleName)
	}
	return covariance(x, y) / (sdx * sdy)
}

// Correlation is the exported wrapper
func Correlation(x, y []float64) float64 {
	return correlation(x, y)
}
