package fuzzy

import (
	"fmt"
	"testing"
)

func ExampleSet_String() {
	a := NewCrispSet([]float64{1, 2, 3, 4})
	fmt.Println(a)
	b := NewFuzzySetFromMF([]float64{1, 2, 3, 4}, EmptyMF)
	fmt.Println(b)
	c := NewFuzzySetFromMF([]float64{1, 2, 3, 4}, func(x float64) float64 {
		return 0.5
	})
	fmt.Println(c)
	// Output:
	// {1/1, 1/2, 1/3, 1/4}
	// {0/1, 0/2, 0/3, 0/4}
	// {0.5/1, 0.5/2, 0.5/3, 0.5/4}
}

func TestSet_AddElement(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("bad grade did not panic")
		}
	}()

	s := NewCrispSet([]float64{})
	s.AddElement(0, -1)
}

func TestSet_IsCrisp(t *testing.T) {
	for i, tt := range []struct {
		s     Set
		crisp bool
	}{
		{
			NewCrispSet([]float64{1, 2}),
			true,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, EmptyMF),
			true,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, NewGaussianMF(0, 1)),
			false,
		},
	} {
		if tt.s.IsCrisp() != tt.crisp {
			t.Errorf("test: %v\tset: %v\n", i, tt.s)
		}
	}
}

func TestSet_IsEqual(t *testing.T) {
	for i, tt := range []struct {
		a, b Set
		want bool
	}{
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{1, 2}),
			true,
		},
		{
			NewCrispSet([]float64{1, 2}),
			NewFuzzySetFromMF([]float64{1, 2}, CrispMF),
			true,
		},
		{
			NewCrispSet([]float64{1, 2}),
			NewFuzzySetFromMF([]float64{1, 2}, EmptyMF),
			false,
		},
		{
			NewCrispSet([]float64{1, 2, 3}),
			NewCrispSet([]float64{1, 2}),
			false,
		},
	} {
		got := tt.a.IsEqual(tt.b)
		if tt.want != got {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestNewFuzzySet_BadLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("element and grade length mismatch did not panic")
		}
	}()

	NewFuzzySet([]float64{1}, []float64{1, 2})
}

func TestNewFuzzySet_BadGrade(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("grade not within interval did not panic")
		}
	}()

	NewFuzzySet([]float64{1}, []float64{-1})
}

func TestSet_Core(t *testing.T) {
	for i, tt := range []struct {
		s, want Set
	}{
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewCrispSet([]float64{2, 3}),
		},
	} {
		got := tt.s.Core()
		if !tt.want.IsEqual(got) {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_Grade(t *testing.T) {
	for i, tt := range []struct {
		s       Set
		u, want float64
	}{
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			2,
			1,
		},
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			5,
			0,
		},
	} {
		got := tt.s.Grade(tt.u)
		if tt.want != got {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_Support(t *testing.T) {
	for i, tt := range []struct {
		s, want Set
	}{
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewCrispSet([]float64{2, 3}),
		},
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.5, 1, 1, 0.5}),
			NewCrispSet([]float64{1, 2, 3, 4}),
		},
	} {
		got := tt.s.Support()
		if !tt.want.IsEqual(got) {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	for i, tt := range []struct {
		a, b, want Set
	}{
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewCrispSet([]float64{2, 3}),
		},
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{1, 0, 0, 1}),
			NewCrispSet([]float64{}),
		},
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{3, 4}),
			NewCrispSet([]float64{}),
		},
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.1, 0.2, 0.3, 0.4}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.4, 0.3, 0.2, 0.1}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.1, 0.2, 0.2, 0.1}),
		},
	} {
		got := tt.a.Intersection(tt.b)
		if !tt.want.IsEqual(got) {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_Union(t *testing.T) {
	for i, tt := range []struct {
		a, b, want Set
	}{
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewCrispSet([]float64{2, 3}),
		},
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0, 1, 1, 0}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{1, 0, 0, 1}),
			NewCrispSet([]float64{1, 2, 3, 4}),
		},
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{3, 4}),
			NewCrispSet([]float64{1, 2, 3, 4}),
		},
		{
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.1, 0.2, 0.3, 0.4}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.4, 0.3, 0.2, 0.1}),
			NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.4, 0.3, 0.3, 0.4}),
		},
	} {
		got := tt.a.Union(tt.b)
		if !tt.want.IsEqual(got) {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_AlphaCut(t *testing.T) {
	for i, tt := range []struct {
		s, want Set
		alpha   float64
	}{
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{1, 2}),
			0.5,
		},
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{1, 2}),
			1.0,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, func(x float64) float64 {
				return 0.5
			}),
			NewCrispSet([]float64{1, 2}),
			0.5,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, func(x float64) float64 {
				return 0.5
			}),
			NewFuzzySetFromMF([]float64{}, EmptyMF),
			0.6,
		},
	} {
		got := tt.s.AlphaCut(tt.alpha)
		if !tt.want.IsEqual(got) {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_StrongAlphaCut(t *testing.T) {
	for i, tt := range []struct {
		s, want Set
		alpha   float64
	}{
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{1, 2}),
			0.5,
		},
		{
			NewCrispSet([]float64{1, 2}),
			NewCrispSet([]float64{}),
			1.0,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, func(x float64) float64 {
				return 0.5
			}),
			NewCrispSet([]float64{}),
			0.5,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, func(x float64) float64 {
				return 0.5
			}),
			NewFuzzySetFromMF([]float64{}, EmptyMF),
			0.6,
		},
	} {
		got := tt.s.StrongAlphaCut(tt.alpha)
		if !tt.want.IsEqual(got) {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
		}
	}
}

func TestSet_IsEmpty(t *testing.T) {
	for i, tt := range []struct {
		s     Set
		crisp bool
	}{
		{
			NewCrispSet([]float64{1, 2}),
			false,
		},
		{
			NewFuzzySetFromMF([]float64{1, 2}, EmptyMF),
			true,
		},
	} {
		if tt.s.IsEmpty() != tt.crisp {
			t.Errorf("test: %v\tset: %v\n", i, tt.s)
		}
	}
}

func TestSet_Compliment(t *testing.T) {
	for i, tt := range []struct {
		s, want Set
	}{
		{
			NewCrispSet([]float64{1, 2, 3, 4}),
			NewFuzzySetFromMF([]float64{1, 2, 3, 4}, EmptyMF),
		},
	} {
		got := tt.s.Compliment().Grades()
		want := tt.want.Grades()
		for j, g := range got {
			if want[j] != g {
				t.Errorf("test: %v got: %v want: %v\n", i, got, tt.want)
			}
		}
	}
}
