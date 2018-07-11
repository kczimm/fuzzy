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
	return &Set{U, Crisp}
}

// StrongAlphaCut func
func (s Set) StrongAlphaCut(alpha float64) *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) > alpha {
			U = append(U, u)
		}
	}
	return &Set{U, Crisp}
}

// Support func
func (s Set) Support() *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) > 0 {
			U = append(U, u)
		}
	}
	return &Set{U, Crisp}
}

// Core func
func (s Set) Core() *Set {
	U := make([]float64, 0)
	for _, u := range s.U {
		if s.m(u) == 1 {
			U = append(U, u)
		}
	}
	return &Set{U, Crisp}
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
