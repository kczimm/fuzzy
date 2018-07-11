package fuzzy

import "math"

// CrispMF func
func CrispMF(x float64) float64 {
	return 1
}

// EmptyMF func
func EmptyMF(x float64) float64 {
	return 0
}

// NewGaussianComboMF func
func NewGaussianComboMF(mu1, sigma1, mu2, sigma2 float64) MembershipFunc {
	if mu1 > mu2 {
		// swap
		mu1, mu2 = mu2, mu1
		sigma1, sigma2 = sigma2, sigma1
	}
	f := NewGaussianMF(mu1, sigma1)
	g := NewGaussianMF(mu2, sigma2)
	return func(x float64) float64 {
		switch {
		case x < mu1:
			return f(x)
		case x > mu2:
			return g(x)
		default:
			return 1
		}
	}
}

// NewGaussianMF func
func NewGaussianMF(mu, sigma float64) MembershipFunc {
	return func(x float64) float64 {
		return math.Exp(-(x - mu) * (x - mu) / (2 * sigma * sigma))
	}
}

// NewSigmoidMF func
func NewSigmoidMF(a, b float64) MembershipFunc {
	return func(x float64) float64 {
		return 1. / (1. + math.Exp(-a*(x-b)))
	}
}

// NewDiffSigmoidMF func
func NewDiffSigmoidMF(a, b, c, d float64) MembershipFunc {
	f := NewSigmoidMF(a, b)
	g := NewSigmoidMF(c, d)
	return func(x float64) float64 {
		return f(x) - g(x)
	}
}

// NewTrapMF func
func NewTrapMF(a, b, c, d float64) MembershipFunc {
	return func(x float64) float64 {
		switch {
		case x >= a && x < b:
			return (x - a) / (b - a)
		case x >= b && x < c:
			return 1
		case x >= c && x < d:
			return (x - d) / (c - d)
		default:
			return 0
		}
	}
}

// NewTriangleMF func
func NewTriangleMF(a, b, c float64) MembershipFunc {
	return func(x float64) float64 {
		switch {
		case x == b:
			return 1
		case x >= a && x < b:
			return (x - a) / (b - a)
		case x > b && x < c:
			return (x - b) / (c - b)
		default:
			return 0
		}
	}
}

// NewBellMF func
func NewBellMF(a, b, c float64) MembershipFunc {
	return func(x float64) float64 {
		return 1. / (1. + math.Pow(math.Abs((x-c)/a), 2*b))
	}
}
