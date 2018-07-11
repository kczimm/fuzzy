package fuzzy

import (
	"math"
	"testing"
)

func TestNewGaussianMF(t *testing.T) {
	eps := 0.0001
	for i, tt := range []struct {
		mu, sigma float64
		u, want   []float64
	}{
		{
			5.,
			2.,
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0.0439, 0.1353, 0.3247, 0.6065, 0.8825, 1, 0.8825, 0.6065, 0.3247, 0.1353, 0.0439},
		},
	} {
		s := Set{
			tt.u,
			NewGaussianMF(tt.mu, tt.sigma),
		}
		got := s.Grades()
		for j := 0; j < len(got); j++ {
			if math.Abs(got[j]-tt.want[j]) > eps {
				t.Errorf("test: %v got: %v want: %v\n", i, got, tt.want)
				break
			}
		}
	}
}
