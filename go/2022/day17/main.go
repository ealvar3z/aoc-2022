package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ealvar3z/aoc2022/go/lib/aoc"
)

type pos struct {
	x, y int
}

func (p pos) add(next pos) pos {
	return pos{
		x: p.x + next.x,
		y: p.y + next.y,
	}
}

var (
	rockTemplates = [][]pos{
		{
			{x: 2, y: 4},
			{x: 3, y: 4},
			{x: 4, y: 4},
			{x: 5, y: 4},
		},
		{
			{x: 3, y: 4},
			{x: 2, y: 5},
			{x: 3, y: 5},
			{x: 4, y: 5},
			{x: 3, y: 6},
		},
		{
			{x: 2, y: 4},
			{x: 3, y: 4},
			{x: 4, y: 4},
			{x: 4, y: 5},
			{x: 4, y: 6},
		},
		{
			{x: 2, y: 4},
			{x: 2, y: 5},
			{x: 2, y: 6},
			{x: 2, y: 7},
		},
		{
			{x: 2, y: 4},
			{x: 3, y: 4},
			{x: 2, y: 5},
			{x: 3, y: 5},
		},
	}

	directions = map[string]pos{
		"<": {x: -1},
		">": {x: 1},
	}
)

func main() {
	file, complete := aoc.Setup(2022, 17, false)
	defer complete()

	input := parse(file)
	aoc.PrintAnswer(1, partOne(input, 2022))
	aoc.PrintAnswer(2, partTwo(input, 1000000000000))
}

func parse(fpath string) []pos {
	f, _ := os.Open(fpath)
	defer f.Close()

	s := bufio.NewScanner(f)

	var positions []pos
	for s.Scan() {
		fields := strings.Split(s.Text(), "")
		for _, field := range fields {
			positions = append(positions, directions[field])
		}
	}
	return positions
}

func partOne(jetDirs []pos, rockCount int) int {
	var board [][7]byte
	jetPos := 0
	for r := 0; r < rockCount; r++ {
		rock := placeIt(&board, rockTemplates[r%len(rockTemplates)])

		for i := 0; ; i++ {
			if i%2 == 0 {
				moveIt(&board, rock, jetDirs[jetPos])
				jetPos = (jetPos + 1) % len(jetDirs) // update the position
			} else {
				if !moveIt(&board, rock, pos{y: -1}) {
					// we've reached the bottom
					settleIt(board, rock)
					break
				}
			}
		}
	}
	return len(board)
}

func partTwo(jetDirs []pos, rockCount uint64) uint64 {
	// let's find the cycle len & height
	var offset int
	var offsetHeight int
	var cycleLen int
	var cycleHeight int
	var cycleJetPos int
	var board [][7]byte
	jetPos := 0
	jetPosAppearances := make(map[int][2]int)
	for r := 0; uint64(r) < rockCount; r++ {
		prevHeight := len(board)
		prevJetPos := jetPos

		rock := placeIt(&board, rockTemplates[r%len(rockTemplates)])

		for i := 0; ; i++ {
			if i%2 == 0 {
				moveIt(&board, rock, jetDirs[jetPos])
				jetPos = (jetPos + 1) % len(jetDirs)
			} else {
				if !moveIt(&board, rock, pos{y: -1}) {
					// we're at the bottom
					settleIt(board, rock)
					break
				}
			}
		}
		if r != 0 && r%len(rockTemplates) == 0 && len(board) == prevHeight+1 {
			if prev, ok := jetPosAppearances[prevJetPos]; ok {
				offset = prev[0]
				offsetHeight = prev[1]
				cycleLen = r - offset
				cycleHeight = prevHeight - prev[1]
				cycleJetPos = prevJetPos
				break
			}
			jetPosAppearances[prevJetPos] = [2]int{r, prevHeight}
		}
	}
	// simulate the remainder of the rocks (i.e. rockCount % cycleLen)
	cycleCount := (rockCount - uint64(offset)) / uint64(cycleLen)
	rockCount = (rockCount - uint64(offset)) % uint64(cycleLen)
	board = nil
	jetPos = cycleJetPos
	for r := 0; uint64(r) < rockCount; r++ {
		rock := placeIt(&board, rockTemplates[r%len(rockTemplates)])

		for i := 0; ; i++ {
			if i%2 == 0 {
				moveIt(&board, rock, jetDirs[jetPos])
				jetPos = (jetPos + 1) % len(jetDirs)
			} else {
				if !moveIt(&board, rock, pos{y: -1}) {
					settleIt(board, rock)
					break
				}
			}
		}
	}
	return uint64(offsetHeight) + cycleCount*uint64(cycleHeight) + uint64(len(board))
}

func placeIt(board *[][7]byte, template []pos) []pos {
	height := len(*board)

	rock := make([]pos, len(template))
	for i, shim := range template {
		rock[i] = shim.add(pos{y: height - 1})
		// extend vertically
		for rock[i].y >= len(*board) {
			*board = append(*board, [7]byte{})
		}
		(*board)[rock[i].y][rock[i].x] = '@'
	}
	return rock
}

func moveIt(board *[][7]byte, rock []pos, dir pos) bool {
	canMove := true
	for _, shim := range rock {
		moveShim := shim.add(dir)
		if moveShim.y < 0 || moveShim.y == len(*board) || moveShim.x < 0 || moveShim.x == 7 {
			// out of bounds error
			canMove = false
			break
		}
		if (*board)[moveShim.y][moveShim.x] == '#' {
			// we've hit a rock that's already set
			canMove = false
			break
		}
	}
	if !canMove {
		return false
	}

	for _, shim := range rock {
		(*board)[shim.y][shim.x] = 0
	}

	for i := range rock {
		rock[i] = rock[i].add(dir)
		(*board)[rock[i].y][rock[i].x] = '@'
	}
	// shrink the board vertically
	for isOcuppied := false; !isOcuppied; {
		for i := 0; i < 7; i++ {
			if (*board)[len(*board)-1][i] != 0 {
				isOcuppied = true
				break
			}
		}
		if !isOcuppied {
			*board = (*board)[:len(*board)-1]
		}
	}
	return true
}

func settleIt(board [][7]byte, rock []pos) {
	for _, shim := range rock {
		board[shim.y][shim.x] = '#'
	}
}
