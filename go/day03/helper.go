package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/ealvar3z/aoc2022/go/lib/cache"
)

type item rune

func (i item) String() string {
	return fmt.Sprintf("%v: %s", string(i), strconv.Itoa(i.priority()))
}

// priority simple converts every item in the ruck to the specifies priority in
// the problem's description.
//
// lowercase items a ... z = 1 ... 26
// uppercase items A ... Z = 27 ... 52
func (i item) priority() int {
	if unicode.IsUpper(rune(i)) {
		return int(i) - 64 + 26
	}
	return int(i) - 96
}

type ruck struct {
	compartments [][]item // a slice for each compartment
	packedTwice  map[item]int
}

func newRuck(items string) *ruck {
	var compartment []item
	var compartments [][]item
	packedTwice := map[item]int{}

	seen := cache.New[rune, bool]()

	for i, r := range []rune(items) {
		compartment = append(compartment, item(r))

		if i+1 <= len(items)/2 {
			seen.Set(r, true)
		} else if seen.Has(r) {
			// if the elf packed the same item twice, then we've seen it on
			// both of the compartments. Thus, we keep track of it.
			if _, ok := packedTwice[item(r)]; ok {
				packedTwice[item(r)] += 1
			} else {
				packedTwice[item(r)] = 1
			}
		}
		if i+1 == len(items)/2 || i+1 == len(items) {
			compartments = append(compartments, compartment)
			compartment = []item{}
		}
	}
	return &ruck{
		compartments: compartments,
		packedTwice:  packedTwice,
	}
}
