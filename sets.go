package fuzzy

import (
	"fmt"
	"strings"
)

type FuzzySet struct {
	U []float64
	m Membership
}

func (f FuzzySet) String() string {
	values := make([]string, len(f.U))
	for i, u := range f.U {
		values[i] = fmt.Sprintf("%v/%v", f.m(u), u)
	}
	return fmt.Sprintf("{%v}", strings.Join(values, ", "))
}

func (f *FuzzySet) IsCrisp() bool {
	for _, u := range f.U {
		if f.m(u) != 1 {
			return false
		}
	}

	return true
}
