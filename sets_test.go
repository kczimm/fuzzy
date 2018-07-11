package fuzzy

import (
	"fmt"
	"testing"
)

func ExampleSet_String() {
	set := Set{
		[]float64{1, 2, 3, 4},
		func(float64) float64 {
			return 1
		},
	}
	fmt.Println(set)
	// Output:
	// {1/1, 1/2, 1/3, 1/4}
}

func TestSet_IsCrisp(t *testing.T) {
	for i, tt := range []struct {
		s     Set
		crisp bool
	}{
		{
			Set{
				[]float64{0., 1.},
				func(float64) float64 {
					return 1
				},
			},
			true,
		},
	} {
		if tt.s.IsCrisp() != tt.crisp {
			t.Errorf("test: %v\tset: %v\n", i, tt.s)
		}
	}
}

func TestSet_Compliment(t *testing.T) {
	for i, tt := range []struct {
		s, want Set
	}{
		{
			Set{
				[]float64{1, 2, 3, 4},
				CrispMF,
			},
			Set{
				[]float64{1, 2, 3, 4},
				EmptyMF,
			},
		},
	} {
		got := tt.s.Compliment()
		for _, u := range got.U {
			if got.m(u) != tt.want.m(u) {
				t.Errorf("test: %v got: %v want: %v\n", i, got, tt.want)
			}
		}
	}
}
