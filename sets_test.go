package fuzzy

import (
	"fmt"
	"testing"
)

func ExampleFuzzySet_String() {
	set := FuzzySet{
		[]float64{1, 2, 3, 4},
		func(float64) float64 {
			return 1
		},
	}
	fmt.Println(set)
	// Output:
	// {1/1, 1/2, 1/3, 1/4}
}

func TestIsCrisp(t *testing.T) {
	for ii, s := range []struct {
		FuzzySet
		bool
	}{
		{
			FuzzySet{
				[]float64{0., 1.},
				func(float64) float64 {
					return 1
				},
			},
			true,
		},
	} {
		got := s.FuzzySet.IsCrisp()
		want := s.bool
		if got != want {
			t.Errorf("test: %v\tset: %v\n", ii, s)
		}
	}
}
