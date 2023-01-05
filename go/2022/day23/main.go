package main

import (
	"bufio"
	"math"
	"os"
	"sort"

	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/numbers"
)

type pos struct {
	x int
	y int
}

type elf struct {
	dirIndex int
}

var (
	N = pos{y: -1}
	S = pos{y: 1}
	E = pos{x: 1}
	W = pos{x: -1}

	DIRS = [4]pos{N, S, E, W}
	ADJS = [8]pos{N, N.add(E), E, E.add(S), S, S.add(W), W, W.add(N)}
)

func (p pos) add(other pos) pos {
	return pos{
		y: p.y + other.y,
		x: p.x + other.x,
	}
}

func main() {
	f, _ := aoc.Setup(2022, 23, false)
	elves := parse(f)
	aoc.PrintAnswer(1, partOne(elves, 10))
	elves = parse(f)
	aoc.PrintAnswer(2, partTwo(elves))
}

func parse(in string) map[pos]*elf {
	f, _ := os.Open(in)
	defer f.Close()

	s := bufio.NewScanner(f)
	elves := make(map[pos]*elf)

	var y int
	for s.Scan() {
		row := []byte(s.Text())
		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				pos := pos{y: y, x: x}
				elves[pos] = &elf{}
			}
		}
		y++
	}
	return elves
}

func partOne(elves map[pos]*elf, rounds int) int {
	for round := 0; round < rounds; round++ {
		compute(elves)
	}

	minY, maxY := math.MaxInt, math.MinInt
	minX, maxX := math.MaxInt, math.MinInt

	for e := range elves {
		minY, maxY = numbers.Min(minY, e.y), numbers.Max(maxY, e.y)
		minX, maxX = numbers.Min(minX, e.x), numbers.Max(maxX, e.x)
	}

	var res int
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if elves[pos{y: y, x: x}] == nil {
				res++
			}
		}
	}
	return res
}

func partTwo(elves map[pos]*elf) int {
	after := make([]pos, 0, len(elves))
	for position := range elves {
		after = append(after, position)
	}
	for round := 1; ; round++ {
		before := after
		compute(elves)
		after = make([]pos, 0, len(elves))
		for position := range elves {
			after = append(after, position)
		}
		if areSame(before, after) {
			return round
		}
	}
}

func compute(elves map[pos]*elf) {
	posMap := make(map[pos]pos)
	posCount := make(map[pos]int)
	for position, elf := range elves {
		solo := true
		for _, adj := range ADJS {
			adjPos := position.add(adj)
			if _, ok := elves[adjPos]; ok {
				solo = false
				break
			}
		}
		if solo {
			continue
		}

		for offset := 0; offset < len(DIRS); offset++ {
			dir := DIRS[(elf.dirIndex+offset)%len(DIRS)]
			proposedPos := position.add(dir)

			check := make([]pos, 0, 3)
			if dir.y == 0 {
				check = append(check, proposedPos, proposedPos.add(N), proposedPos.add(S))
			} else {
				check = append(check, proposedPos, proposedPos.add(W), proposedPos.add(E))
			}

			cleared := true
			for _, pos := range check {
				if _, ok := elves[pos]; ok {
					cleared = false
					break
				}
			}

			if cleared {
				posMap[position] = proposedPos
				posCount[proposedPos]++
				break
			}
		}
	}
	remove := make(map[pos]*elf)
	_add := make(map[pos]*elf)
	for _pos, elf := range elves {
		if proposedPos, ok := posMap[_pos]; ok && posCount[proposedPos] == 1 {
			// move it to proposed position
			remove[_pos] = elf
			_add[proposedPos] = elf
		}
		// cycle through the dirs
		elf.dirIndex = (elf.dirIndex + 1) % len(DIRS)
	}

	for pos := range remove {
		delete(elves, pos)
	}
	for pos, elf := range _add {
		elves[pos] = elf
	}
}

func areSame(before []pos, after []pos) bool {
	if len(before) != len(after) {
		return false
	}

	sort.Slice(before, func(i, j int) bool {
		if before[i].y == before[j].y {
			return before[i].x < before[j].x
		}
		return before[i].y < before[j].y
	})
	sort.Slice(after, func(i, j int) bool {
		if after[i].y == after[j].y {
			return after[i].x < after[j].x
		}
		return after[i].y < after[j].y
	})
	for i := 0; i < len(before); i++ {
		if before[i] != after[i] {
			return false
		}
	}
	return true
}
