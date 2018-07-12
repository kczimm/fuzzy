package fuzzy

import (
	"fmt"
	"math"
	"testing"
)

func ExampleNewEmptySet() {
	s := NewEmptySet()

	fmt.Println(s)
	// Output:
	// {}
}

func ExampleNewCrispSet() {
	s := NewCrispSet([]float64{1, 2, 3, 4})

	fmt.Println(s)
	// Output:
	// {1/1, 1/2, 1/3, 1/4}
}

func ExampleNewFuzzySet() {
	s := NewFuzzySet([]float64{1, 2, 3, 4}, []float64{0.1, 0.2, 0.3, 0.4})

	fmt.Println(s)
	// Output:
	// {0.1/1, 0.2/2, 0.3/3, 0.4/4}
}

func ExampleNewFuzzySetFromMF() {
	a := NewFuzzySetFromMF([]float64{1, 2, 3, 4}, EmptyMF)
	b := NewFuzzySetFromMF([]float64{1, 2, 3, 4}, CrispMF)
	c := NewFuzzySetFromMF(
		[]float64{1, 2, 3, 4},
		func(x float64) float64 {
			return 0.5
		},
	)
	d := NewFuzzySetFromMF(
		[]float64{0, 1, 2, 3, 4, 5},
		NewTrapMF(0, 2, 3, 5),
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// Output:
	// {0/1, 0/2, 0/3, 0/4}
	// {1/1, 1/2, 1/3, 1/4}
	// {0.5/1, 0.5/2, 0.5/3, 0.5/4}
	// {0/0, 0.5/1, 1/2, 1/3, 0.5/4, 0/5}
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

func TestSet_Centroid(t *testing.T) {
	eps := 1e-4
	for i, tt := range []struct {
		s    Set
		want float64
	}{
		{
			NewFuzzySetFromMF(
				[]float64{10.0, -9.9, -9.8, -9.7, -9.6, -9.5, -9.4, -9.3, -9.2, -9.1, -9.0, -8.9, -8.8, -8.7, -8.6, -8.5, -8.4, -8.3, -8.2, -8.1, -8.0, -7.9, -7.8, -7.7, -7.6, -7.5, -7.4, -7.3, -7.2, -7.1, -7.0, -6.9, -6.8, -6.7, -6.6, -6.5, -6.4, -6.3, -6.2, -6.1, -6.0, -5.9, -5.8, -5.7, -5.6, -5.5, -5.4, -5.3, -5.2, -5.1, -5.0, -4.9, -4.8, -4.7, -4.6, -4.5, -4.4, -4.3, -4.2, -4.1, -4.0, -3.9, -3.8, -3.7, -3.6, -3.5, -3.4, -3.3, -3.2, -3.1, -3.0, -2.9, -2.8, -2.7, -2.6, -2.5, -2.4, -2.3, -2.2, -2.1, -2.0, -1.9, -1.8, -1.7, -1.6, -1.5, -1.4, -1.3, -1.2, -1.1, -1.0, -0.9, -0.8, -0.7, -0.6, -0.5, -0.4, -0.3, -0.2, -0.1, 0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 4.0, 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5.0, 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0, 6.1, 6.2, 6.3, 6.4, 6.5, 6.6, 6.7, 6.8, 6.9, 7.0, 7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8, 7.9, 8.0, 8.1, 8.2, 8.3, 8.4, 8.5, 8.6, 8.7, 8.8, 8.9, 9.0, 9.1, 9.2, 9.3, 9.4, 9.5, 9.6, 9.7, 9.8, 9.9, 10.0},
				NewTrapMF(-10, -8, -4, 7),
			),
			-3.2857,
		},
	} {
		got := tt.s.Centroid()
		if math.Abs(got-tt.want) > eps {
			t.Errorf("test: %v got: %v want: %v", i, got, tt.want)
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
