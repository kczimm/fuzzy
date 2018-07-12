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
}

// NewEmptySet func
func NewEmptySet() Set {
	return Set{
		make(map[float64]float64),
	}
}

// NewFuzzySet func
func NewFuzzySet(u, g []float64) Set {
	if len(u) != len(g) {
		panic("elements and grades have different lengths")
	}
	s := NewEmptySet()
	for i, v := range u {
		s.AddElement(v, g[i])
	}
	return s
}

// NewFuzzySetFromMF func
func NewFuzzySetFromMF(u []float64, m MembershipFunc) Set {
	s := Set{
		make(map[float64]float64),
	}
	for _, v := range u {
		s.AddElement(v, m(v))
	}
	return s
}

// NewCrispSet func
func NewCrispSet(u []float64) Set {
	return NewFuzzySetFromMF(u, CrispMF)
}

// AddElement method
func (s *Set) AddElement(u, g float64) {
	if g < 0 || g > 1 {
		panic("grade outside the interval [0, 1]")
	}
	s.data[u] = g
}

// Compliment method
func (s Set) Compliment() Set {
	c := NewCrispSet(s.Elements())
	for k, v := range c.data {
		c.data[k] = 1 - v
	}
	return c
}

// Intersection method
func (s Set) Intersection(other Set) Set {
	n := NewEmptySet()
	for k, v := range s.data {
		g, exists := other.data[k]
		if exists {
			if g < v {
				n.AddElement(k, g)
			} else {
				n.AddElement(k, v)
			}
		}
	}
	return n
}

// Union method
func (s Set) Union(other Set) Set {
	n := NewEmptySet()
	for k, v := range s.data {
		g, exists := other.data[k]
		if exists && g > v {
			n.AddElement(k, g)
		} else {
			n.AddElement(k, v)
		}
	}
	for k, v := range other.data {
		g, exists := s.data[k]
		if exists && g > v {
			n.AddElement(k, g)
		} else {
			n.AddElement(k, v)
		}
	}
	return n
}

// Grades func
func (s Set) Grades() []float64 {
	grades := make([]float64, len(s.data))
	for i, e := range s.Elements() {
		grades[i] = s.data[e]
	}
	return grades
}

// Grade method
func (s Set) Grade(u float64) float64 {
	g, exists := s.data[u]
	if !exists {
		return 0
	}
	return g
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

// IsEqual func
func (s Set) IsEqual(other Set) bool {
	for e := range s.data {
		g, exists := other.data[e]
		if !exists || g != s.data[e] {
			return false
		}
	}
	return true
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
	for e, g := range s.data {
		if g >= alpha {
			U = append(U, e)
		}
	}
	return NewCrispSet(U)
}

// StrongAlphaCut func
func (s Set) StrongAlphaCut(alpha float64) Set {
	U := make([]float64, 0)
	for e, g := range s.data {
		if g > alpha {
			U = append(U, e)
		}
	}
	return NewCrispSet(U)
}

// Support func
func (s Set) Support() Set {
	U := make([]float64, 0)
	for e, g := range s.data {
		if g > 0 {
			U = append(U, e)
		}
	}
	return NewCrispSet(U)
}

// Core func
func (s Set) Core() Set {
	U := make([]float64, 0)
	for e, g := range s.data {
		if g == 1 {
			U = append(U, e)
		}
	}
	return NewCrispSet(U)
}

// IsCrisp func
func (s Set) IsCrisp() bool {
	for _, g := range s.data {
		if g != 1 && g != 0 {
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
