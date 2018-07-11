package fuzzy

import (
	"fmt"
	"testing"
)

func ExampleSet_String() {
	a := NewCrispSet([]float64{1, 2, 3, 4})
	fmt.Println(a)
	b := NewFuzzySet([]float64{1, 2, 3, 4}, EmptyMF)
	fmt.Println(b)
	c := NewFuzzySet([]float64{1, 2, 3, 4}, func(x float64) float64 {
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
			t.Errorf("bad membership function did not panic")
		}
	}()

	s := NewFuzzySet([]float64{}, func(x float64) float64 {
		return 2
	})
	s.AddElement(0)
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
			NewFuzzySet([]float64{1, 2}, EmptyMF),
			true,
		},
		{
			NewFuzzySet([]float64{1, 2}, NewGaussianMF(0, 1)),
			false,
		},
	} {
		if tt.s.IsCrisp() != tt.crisp {
			t.Errorf("test: %v\tset: %v\n", i, tt.s)
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
			NewFuzzySet([]float64{1, 2}, EmptyMF),
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
			NewFuzzySet([]float64{1, 2, 3, 4}, EmptyMF),
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
