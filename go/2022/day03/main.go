package main

import (
	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/file"
)

// There are 2 compartments in each ruck
// We have to breakdown the input in 2
func main() {
	path, complete := aoc.Setup(2022, 3, false)
	defer complete()

	var rucks []*ruck
	for _, v := range file.TextLines(path) {
		rucks = append(rucks, newRuck(v))
	}

	// Part one: We have to find the total sum of all the items that exist in both
	// compartments. When we instantiated the `newRuck` struct, we've already
	// captured that information (see helper.go). Now, it's just a matter of
	// adding the values.
	var part1 int
	for _, r := range rucks {
		for v := range r.packedTwice {
			part1 += v.priority()
		}

	}

	// Part two: We have to find the item that exists in the batch. The batch
	// is every group of (3), and the common item in the batch is the missing
	// badge. Then, we sum the total of their priorities.
	var part2 int

	// So, we have to iterate in groups of (3) i.e. we start at the first
	// three (i := 2) in the ruck(i.e len(rucks)), grab the first three
	// (rucks[i-2] == rucks[0]) and then step every (3) lines.
	for i := 2; i < len(rucks); i += 3 {
		ruckSet := []*ruck{rucks[i-2], rucks[i-1], rucks[i]}
		count := map[item]int{}
		for i, r := range ruckSet {
			// if we seen it when rummaging trough the rucksack, then we don't
			// include it in our total count.
			weSawit := map[item]bool{}
			for _, compartment := range r.compartments {
				for _, item := range compartment {
					if weSawit[item] {
						continue
					}
					weSawit[item] = true
					if _, ok := count[item]; ok {
						count[item] += 1
					} else {
						count[item] = 1
					}
					if i == 2 && count[item] == 3 {
						part2 += item.priority()
					}
				}
			}
		}
	}
	aoc.PrintAnswer(1, part1)
	aoc.PrintAnswer(2, part2)
}
