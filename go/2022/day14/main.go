package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]byte, 1000)

	// simulate the falling sand. How many units of sand come to rest before
	// sand starts flowing into the abyss below?
	simulate := func() (n int) {
		n = 0
		for {
			// starting point
			x, y := 500, 0
			for y < 999 {
				drop := func(dx int) bool {
					// so i can see it in the debugger
					down := y + 1
					diag := x + dx
					if grid[down][diag] == '.' {
						y += 1
						x += dx
						return true
					}
					return false
				}
				// we've come to a rest
				if !drop(0) && !drop(-1) && !drop(1) {
					grid[y][x] = 'o'
					break
				}
			}
			if y == 999 || x == 500 && y == 0 {
				break
			}
			n += 1
		}
		return n
	}

	for i := 0; i < 1000; i++ {
		grid[i] = make([]byte, 1000)
		for j := 0; j < 1000; j++ {
			grid[i][j] = '.'
		}
	}
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " -> ")
		var x, y int
		fmt.Sscanf(pairs[0], "%d,%d", &x, &y)
		for i := 1; i < len(pairs); i++ {
			var toX, toY int
			fmt.Sscanf(pairs[i], "%d,%d", &toX, &toY)
			dx, dy := direction(x, toX), direction(y, toY)
			for ; x != toX || y != toY; x, y = x+dx, y+dy {
				grid[y][x] = '#'
			}
			grid[y][x] = '#'
		}
	}
	part_one := simulate()
	fmt.Println(part_one)
	// part two
	// Using your scan, simulate the falling sand until the source of the sand
	// becomes blocked. How many units of sand come to rest?
	floor := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[y][x] != '.' && y > floor {
				floor = y
			}
		}
	}
	floor += 2
	for x := 0; x < 1000; x++ {
		// we got it!
		grid[floor][x] = '#'
	}
	fmt.Println(part_one + simulate() + 1)
}

func direction(from, to int) int {
	if from == to {
		return 0
	}
	if from < to {
		return 1
	}
	return -1
}
