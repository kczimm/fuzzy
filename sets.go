package fuzzy

import (
	"fmt"
	"sort"
	"strings"
)

// MembershipFunc type
type MembershipFunc func(float64) float64

// Set struct
type Set struct {
	data map[float64]float64
	m    MembershipFunc
}

// NewFuzzySet func
func NewFuzzySet(u []float64, m MembershipFunc) Set {
	s := Set{
		make(map[float64]float64),
		m,
	}
	for _, v := range u {
		s.AddElement(v)
	}
	return s
}

// NewCrispSet func
func NewCrispSet(u []float64) Set {
	return NewFuzzySet(u, CrispMF)
}

// AddElement method
func (s *Set) AddElement(x float64) {
	g := s.m(x)
	if g < 0 || g > 1 {
		panic("MF function returned a grade outside the interval [0, 1]")
	}
	s.data[x] = s.m(x)
}

// Compliment method
func (s Set) Compliment() Set {
	U := s.Elements()
	return NewFuzzySet(
		U,
		func(x float64) float64 {
			return 1 - s.m(x)
		},
	)
}

// Grades func
func (s Set) Grades() []float64 {
	grades := make([]float64, len(s.data))
	for i, e := range s.Elements() {
		grades[i] = s.data[e]
	}
	return grades
}

// Elements func
func (s Set) Elements() []float64 {
	e := make([]float64, len(s.data))
	i := 0
	for u := range s.data {
		e[i] = u
		i++
	}
	sort.Float64s(e)
	return e
}

func (s Set) String() string {
	values := make([]string, len(s.data))
	elements := s.Elements()
	for i, e := range elements {
		values[i] = fmt.Sprintf("%v/%v", s.data[e], e)
	}
	return fmt.Sprintf("{%v}", strings.Join(values, ", "))
}

// AlphaCut func
func (s Set) AlphaCut(alpha float64) Set {
	U := make([]float64, 0)
	for _, g := range s.data {
		if g >= alpha {
			U = append(U, g)
		}
	}
	return NewCrispSet(U)
}

// StrongAlphaCut func
func (s Set) StrongAlphaCut(alpha float64) Set {
	U := make([]float64, 0)
	for _, g := range s.data {
		if g > alpha {
			U = append(U, g)
		}
	}
	return NewCrispSet(U)
}

// Support func
func (s Set) Support() Set {
	U := make([]float64, 0)
	for _, g := range s.data {
		if g > 0 {
			U = append(U, g)
		}
	}
	return NewCrispSet(U)
}

// Core func
func (s Set) Core() Set {
	U := make([]float64, 0)
	for _, g := range s.data {
		if g == 1 {
			U = append(U, g)
		}
	}
	return NewCrispSet(U)
}

// IsCrisp func
func (s Set) IsCrisp() bool {
	for _, g := range s.data {
		if g != 1 {
			return false
		}
	}

	return true
}

// IsEmpty func
func (s Set) IsEmpty() bool {
	for _, g := range s.data {
		if g != 0 {
			return false
		}
	}

	return true
}
