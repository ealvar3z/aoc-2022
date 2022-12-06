package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/file"
)

func main() {
	path, complete := aoc.Setup(2022, 4, false)
	defer complete()

	var pairs [][]task
	for _, ranges := range file.TextSplit(path, ",") {
		var pair []task
		for _, section := range ranges {
			values := strings.Split(section, "-")
			left, _ := strconv.Atoi(values[0])
			right, _ := strconv.Atoi(values[1])

			pair = append(pair, task{
				left:  left,
				right: right,
			})
		}
		pairs = append(pairs, pair)
	}
	// Part 1: In how many assignment pairs does one range fully contain the other?
	// now it's just a matter of calling Contains() & Overlap() over the pairs
	var part1 int
	var part2 int
	for _, pair := range pairs {
		if pair[0].Contains(pair[1]) || pair[1].Contains(pair[0]) {
			part1 += 1
		}
		if pair[0].Overlaps(pair[1]) || pair[1].Overlaps(pair[0]) {
			fmt.Println(pair[0], pair[1])
			part2 += 1
		}
	}
	aoc.PrintAnswer(1, part1)
	aoc.PrintAnswer(2, part2)
}
