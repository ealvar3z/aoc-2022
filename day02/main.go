package main

import (
	"bufio"
	"log"
	"os"
)

// Super hacky
func part1(f *os.File) {
	s := bufio.NewScanner(f)
	scores := map[string]struct{ part1 int }{
		"A X": {1 + 3}, "A Y": {2 + 6}, "A Z": {3 + 0},
		"B X": {1 + 0}, "B Y": {2 + 3}, "B Z": {3 + 6},
		"C X": {1 + 6}, "C Y": {2 + 0}, "C Z": {3 + 3},
	}
	score := 0
	for s.Scan() {
		score += scores[s.Text()].part1
	}
	log.Printf("Part 1 = %+v\n", score)
}

// Super hacky
func part2(f *os.File) {
	s := bufio.NewScanner(f)
	scores := map[string]struct{ part2 int }{
		"A X": {3 + 0}, "A Y": {1 + 3}, "A Z": {2 + 6},
		"B X": {1 + 0}, "B Y": {2 + 3}, "B Z": {3 + 6},
		"C X": {2 + 0}, "C Y": {3 + 3}, "C Z": {1 + 6},
	}
	score2 := 0
	for s.Scan() {
		score2 += scores[s.Text()].part2
	}
	log.Printf("Part 2 = %+v\n", score2)
}

func main() {
	// part 1
	f, _ := os.Open("input.txt")
	defer f.Close()
	part1(f)

	// part 2
	g, _ := os.Open("input.txt")
	defer g.Close()
	part2(g)
}
