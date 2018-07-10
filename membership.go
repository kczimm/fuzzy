package fuzzy

type Membership func(float64) float64

func Crisp(x float64) float64 {
	return 1
}
