package main

import (
	"fmt"
)

func main() {
	c := NewChecker("on")
	fmt.Println(c.ExpandsInto("nnt"))
	fmt.Println(c.ExpandsInto("not"))
	fmt.Println(c.ExpandsInto("now"))
	fmt.Println(c.ExpandsInto("on"))
}

type Checker struct {
	len  int
	freq map[rune]int
}

func NewChecker(w string) Checker {
	freq := map[rune]int{}
	for _, r := range w {
		if _, ok := freq[r]; !ok {
			freq[r] = 0
		}
		freq[r] += 1
	}

	return Checker{
		freq: freq,
		len:  len(w),
	}
}

func (c *Checker) ExpandsInto(w string) bool {
	if c.len+1 != len(w) {
		return false
	}

	used := map[rune]int{}
	extra := 0
	for _, r := range w {
		if c.freq[r] > used[r] {
			used[r]++
		} else {
			extra++
			if extra > 1 {
				return false
			}
		}
	}

	return true
}
