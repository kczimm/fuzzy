package fuzzy

import "math"

// Membership type
type Membership func(float64) float64

// Crisp func
func Crisp(x float64) float64 {
	return 1
}

// NewGaussianMF func
func NewGaussianMF(mu, sigma float64) Membership {
	return func(x float64) float64 {
		return math.Exp(-(x - mu) * (x - mu) / (2 * sigma * sigma))
	}
}
