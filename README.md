# fuzzy [![Build Status](https://travis-ci.org/kczimm/fuzzy.svg?branch=master)](https://travis-ci.org/kczimm/fuzzy) [![Coverage Status](https://coveralls.io/repos/github/kczimm/fuzzy/badge.svg)](https://coveralls.io/github/kczimm/fuzzy?branch=master) [![GoDoc](https://godoc.org/github.com/kczimm/fuzzy?status.svg)](https://godoc.org/github.com/kczimm/fuzzy)

## Example
```
package main

import (
	"fmt"
	"github.com/kczimm/fuzzy"
)

func main() {
	t := fuzzy.NewVariable("temp") // degrees F
	t.AddClass("cold", fuzzy.NewTrapMF(0, 0, 30, 40))
	t.AddClass("hot", fuzzy.NewTrapMF(60, 70, 100, 100))

	f := fuzzy.NewVariable("fanspeed") // RPM
	f.AddClass("slow", fuzzy.NewTrapMF(0, 0, 10, 20))
	f.AddClass("fast", fuzzy.NewTrapMF(40, 50, 100, 100))

	sys := fuzzy.FIS{}
	sys.AddInput(t)
	sys.AddOutput(f)

	sys.AddRule([]string{"temp", "hot", "then", "fanspeed", "fast"})
	sys.AddRule([]string{"temp", "cold", "then", "fanspeed", "slow"})

	temp := 80
	fanspeed := sys.Evaluate(temp)
	fmt.Println(fanspeed)
}
```
