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

// Compliment method
func (s Set) Compliment() Set {
	u := make([]float64, len(s.U))
	copy(u, s.U)
	return Set{
		u,
		func(x float64) float64 {
			return 1 - s.m(x)
		},
	}
}

// Grades func
func (s Set) Grades() []float64 {
	g := make([]float64, len(s.U))
	for i, u := range s.U {
		g[i] = s.m(u)
	}
	return g
}

func (s Set) String() string {
	values := make([]string, len(s.U))
	for i, u := range s.U {
		values[i] = fmt.Sprintf("%v/%v", s.m(u), u)
	}
	return fmt.Sprintf("{%v}", strings.Join(values, ", "))
}

// AlphaCut func
func (s Set) AlphaCut(alpha float64) *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) >= alpha {
			U = append(U, u)
		}
	}
	return &Set{U, CrispMF}
}

// StrongAlphaCut func
func (s Set) StrongAlphaCut(alpha float64) *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) > alpha {
			U = append(U, u)
		}
	}
	return &Set{U, CrispMF}
}

// Support func
func (s Set) Support() *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) > 0 {
			U = append(U, u)
		}
	}
	return &Set{U, CrispMF}
}

// Core func
func (s Set) Core() *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) == 1 {
			U = append(U, u)
		}
	}
	return &Set{U, CrispMF}
}

// IsCrisp func
func (s Set) IsCrisp() bool {
	for _, u := range s.U {
		if s.m(u) != 1 {
			return false
		}
	}

	return true
}

// IsEmpty func
func (s Set) IsEmpty() bool {
	for _, u := range s.U {
		if s.m(u) != 0 {
			return false
		}
	}

	return true
}
