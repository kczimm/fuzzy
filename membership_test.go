package fuzzy

import (
	"math"
	"sort"
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
		s := NewFuzzySet(
			tt.u,
			NewGaussianMF(tt.mu, tt.sigma),
		)
		got := s.Grades()
		sort.Float64s(got)
		sort.Float64s(tt.want)
		for j := 0; j < len(got); j++ {
			if math.Abs(got[j]-tt.want[j]) > eps {
				t.Errorf("test: %v got: %v want: %v\n", i, got[j], tt.want[j])
				break
			}
		}
	}
}

func TestNewSigmoidMF(t *testing.T) {
	eps := 0.0001
	for i, tt := range []struct {
		a, b    float64
		u, want []float64
	}{
		{
			2,
			4,
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0.0003, 0.0025, 0.018, 0.1192, 0.5, 0.8808, 0.9820, 0.9975, 0.9997, 1, 1},
		},
	} {
		s := NewFuzzySet(
			tt.u,
			NewSigmoidMF(tt.a, tt.b),
		)
		got := s.Grades()
		sort.Float64s(got)
		sort.Float64s(tt.want)
		for j := 0; j < len(got); j++ {
			if math.Abs(got[j]-tt.want[j]) > eps {
				t.Errorf("test: %v got: %v want: %v\n", i, got[j], tt.want[j])
				break
			}
		}
	}
}

func TestNewTrapMF(t *testing.T) {
	eps := 0.0001
	for i, tt := range []struct {
		a, b, c, d float64
		u, want    []float64
	}{
		{
			1, 5, 7, 10,
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 0, 0.25, 0.5, 0.75, 1, 1, 1, 0.6666, 0.3333, 0},
		},
	} {
		s := NewFuzzySet(
			tt.u,
			NewTrapMF(tt.a, tt.b, tt.c, tt.d),
		)
		got := s.Grades()
		sort.Float64s(got)
		sort.Float64s(tt.want)
		for j := 0; j < len(got); j++ {
			if math.Abs(got[j]-tt.want[j]) > eps {
				t.Errorf("test: %v got: %v want: %v\n", i, got[j], tt.want[j])
				break
			}
		}
	}
}

func TestNewTriangleMF(t *testing.T) {
	eps := 0.0001
	for i, tt := range []struct {
		a, b, c float64
		u, want []float64
	}{
		{
			3, 6, 8,
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 0, 0, 0, 0.3333, 0.6666, 1, 0.5, 0, 0, 0},
		},
	} {
		s := NewFuzzySet(
			tt.u,
			NewTriangleMF(tt.a, tt.b, tt.c),
		)
		got := s.Grades()
		sort.Float64s(got)
		sort.Float64s(tt.want)
		for j := 0; j < len(got); j++ {
			if math.Abs(got[j]-tt.want[j]) > eps {
				t.Errorf("test: %v got: %v want: %v\n", i, got[j], tt.want[j])
				break
			}
		}
	}
}

func TestNewBellMF(t *testing.T) {
	eps := 0.0001
	for i, tt := range []struct {
		a, b, c float64
		u, want []float64
	}{
		{
			2, 4, 6,
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0.0002, 0.0007, 0.0039, 0.0375, 0.5, 0.9961, 1, 0.9961, 0.5, 0.0375, 0.0039},
		},
	} {
		s := NewFuzzySet(
			tt.u,
			NewBellMF(tt.a, tt.b, tt.c),
		)
		got := s.Grades()
		sort.Float64s(got)
		sort.Float64s(tt.want)
		for j := 0; j < len(got); j++ {
			if math.Abs(got[j]-tt.want[j]) > eps {
				t.Errorf("test: %v got: %v want: %v\n", i, got[j], tt.want[j])
				break
			}
		}
	}
}
