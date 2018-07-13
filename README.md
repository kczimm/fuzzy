# fuzzy [![Build Status](https://travis-ci.org/kczimm/fuzzy.svg?branch=master)](https://travis-ci.org/kczimm/fuzzy) [![Coverage Status](https://coveralls.io/repos/github/kczimm/fuzzy/badge.svg)](https://coveralls.io/github/kczimm/fuzzy?branch=master) [![GoDoc](https://godoc.org/github.com/kczimm/fuzzy?status.svg)](https://godoc.org/github.com/kczimm/fuzzy)

## Example
```
package main

import (
	"fmt"
	"github.com/kczimm/fuzzy"
)

func main() {
	scale := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := fuzzy.NewVariable("service")
	s.AddClass(
		"poor",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewGaussianMF(0, 3),
		),
	)
	s.AddClass(
		"good",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewGaussianMF(5, 3),
		),
	)
	s.AddClass(
		"excellent",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewGaussianMF(10, 3),
		),
	)

	f := fuzzy.NewVariable("food")
	f.AddClass(
		"rancid",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewTrapMF(0, 0, 1, 3),
		),
	)
	f.AddClass(
		"delicious",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewTrapMF(0, 6, 9, 10),
		),
	)

	scale = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	t := fuzzy.NewVariable("tip")
	t.AddClass(
		"cheap",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewTriangleMF(0, 5, 10),
		),
	)
	t.AddClass(
		"average",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewTriangleMF(10, 15, 20),
		),
	)
	t.AddClass(
		"generous",
		fuzzy.NewFuzzySetFromMF(
			scale,
			fuzzy.NewTriangleMF(15, 25, 30),
		),
	)

	sys := fuzzy.FIS{}
	sys.AddInput(f)
	sys.AddInput(s)
	sys.AddOutput(t)

	sys.AddRule("service", "poor", "or", "food", "rancid", "then", "tip", "cheap")
	sys.AddRule("serice", "good", "then", "tip", "average")
	sys.AddRule("serice", "excellent", "or", "food", "delicious", "then", "tip", "generous")

	service, food := 3, 8
	tip := sys.Evaluate(service, food)
	fmt.Println(tip)
}
```
