package fuzzy

import "math"

// Membership type
type Membership func(float64) float64

// CrispMF func
func CrispMF(x float64) float64 {
	return 1
}

// EmptyMF func
func EmptyMF(x float64) float64 {
	return 0
}

// NewGaussianMF func
func NewGaussianMF(mu, sigma float64) Membership {
	return func(x float64) float64 {
		return math.Exp(-(x - mu) * (x - mu) / (2 * sigma * sigma))
	}
}

// NewSigmoidMF func
func NewSigmoidMF(a, b float64) Membership {
	return func(x float64) float64 {
		return 1. / (1. + math.Exp(-a*(x-b)))
	}
}

// NewTrapMF func
func NewTrapMF(a, b, c, d float64) Membership {
	return func(x float64) float64 {
		switch {
		case x >= a && x < b:
			return 1 / (b - a) * (x - a)
		case x >= b && x < c:
			return 1
		case x >= c && x < d:
			return 1 / (c - d) * (x - d)
		default:
			return 0
		}
	}
}
