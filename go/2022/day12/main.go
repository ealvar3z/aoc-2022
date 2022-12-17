package main

import (
	"fmt"

	"github.com/ealvar3z/aoc2022/go/lib/algos"
	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/file"
)

func main() {
	path, complete := aoc.Setup(2022, 12, false)
	defer complete()

	grid := file.TextLines(path)
	// dijkstra
	g := algos.Graph{}
	var start, end string
	var begins []string
	for y, row := range grid {
		for x, col := range row {
			switch col {
			case 'S':
				start = position(x, y)
				grid[y] = replace(grid[y], x, 'a')
			case 'E':
				end = position(x, y)
				grid[y] = replace(grid[y], x, 'z')
			case 'a':
				begins = append(begins, position(x, y))
			}
		}
	}
	for y, row := range grid {
		for x, col := range row {
			if _, found_it := g[position(x, y)]; !found_it {
				node := map[string]int{}
				edge := func(x, y int) {
					if x >= 0 && x < len(row) &&
						y >= 0 && y < len(grid) &&
						int(grid[y][x])-int(col) < 2 {
						node[position(x, y)] = 1
					}
				}
				edge(x+1, y)
				edge(x-1, y)
				edge(x, y+1)
				edge(x, y-1)
				g[position(x, y)] = node
			}
		}
	}
	_, node1, _ := g.Path(start, end)
	// set an arbitraty large amount of nodes
	node2 := 1000
	for _, s := range begins {
		_, cost, _ := g.Path(s, end)
		if cost > 0 && cost < node2 {
			node2 = cost
		}
	}
	fmt.Println(node1, node2)
}

func position(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func replace(s string, i int, r rune) string {
	runes := []rune(s)
	runes[i] = r
	return string(runes)
}
