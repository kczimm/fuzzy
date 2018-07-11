package fuzzy

import (
	"fmt"
	"strings"
)

// Set struct
type Set struct {
	U []float64
	m Membership
}

func (s Set) String() string {
	values := make([]string, len(s.U))
	for i, u := range s.U {
		values[i] = fmt.Sprintf("%v/%v", s.m(u), u)
	}
	return fmt.Sprintf("{%v}", strings.Join(values, ", "))
}

// IsCrisp func
func (s *Set) IsCrisp() bool {
	for _, u := range s.U {
		if s.m(u) != 1 {
			return false
		}
	}

	return true
}
