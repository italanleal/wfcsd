package math

import "log"

// cov computes the raw covariance, assumes x and y are valid
func covariance(x, y []float64) float64 {
	mx := mean(x)
	my := mean(y)
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += (x[i] - mx) * (y[i] - my)
	}
	return sum / float64(len(x)-1)
}

// Covariance is the safe, exported function
func Covariance(x, y []float64) float64 {
	if len(x) != len(y) {
		log.Fatalf("%s | length mismatch: %d vs %d", moduleName, len(x), len(y))
	}
	if len(x) < 2 {
		log.Fatalf("%s | covariance requires at least 2 values, got %d", moduleName, len(x))
	}
	return covariance(x, y)
}
