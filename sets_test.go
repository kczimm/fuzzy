package fuzzy

import (
	"fmt"
	"testing"
)

func ExampleSet_String() {
	set := NewCrispSet([]float64{1, 2, 3, 4})
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
			NewCrispSet([]float64{1, 2}),
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
