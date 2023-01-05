package main

import (
	"bufio"
	"container/list"
	"os"
	"strings"

	"github.com/ealvar3z/aoc2022/go/lib/aoc"
	"github.com/ealvar3z/aoc2022/go/lib/numbers"
)

type pos struct {
	x int
	y int
}

type state struct {
	pos  pos
	time int
}

var (
	U = pos{y: -1}
	D = pos{y: 1}
	L = pos{x: -1}
	R = pos{x: 1}

	DirMap = map[byte]pos{
		'^': U,
		'v': D,
		'<': L,
		'>': R,
	}
)

func (p pos) add(other pos) pos {
	return pos{
		y: p.y + other.y,
		x: p.x + other.x,
	}
}

func (p pos) mult(val int) pos {
	return pos{
		y: p.y * val,
		x: p.x * val,
	}
}

func (p pos) mod(y, x int) pos {
	return pos{
		y: ((p.y % y) + y) % y,
		x: ((p.x % x) + x) % x,
	}
}

func main() {
	f, _ := aoc.Setup(2022, 24, false)
	board, begin, end := parse(f)
	// fmt.Println(printBoard(board))
	// fmt.Println(begin, end)
	blizzards := blizzGenerator(board)
	aoc.PrintAnswer(1, partOne(board, begin, end, blizzards))
	aoc.PrintAnswer(2, partTwo(board, begin, end, blizzards))
}

func parse(in string) ([][]byte, pos, pos) {
	f, _ := os.Open(in)
	defer f.Close()

	s := bufio.NewScanner(f)

	var board [][]byte
	for s.Scan() {
		row := []byte(s.Text())
		board = append(board, row)
	}

	var begin, end pos
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == '.' {
			begin = pos{y: 0, x: i}
		}
		if board[len(board)-1][i] == '.' {
			end = pos{y: len(board) - 1, x: i}
		}
	}
	return board, begin, end
}

func blizzGenerator(board [][]byte) []map[pos]bool {
	dy := len(board) - 2
	dx := len(board[0]) - 2
	maxTime := numbers.Lcm(dy, dx)

	var blizzards []map[pos]bool
	for t := 0; t < maxTime; t++ {
		_blizzards := make(map[pos]bool)
		for y := 1; y < len(board)-1; y++ {
			for x := 1; x < len(board[0])-1; x++ {
				if board[y][x] != '.' {
					inititally := pos{y: y - 1, x: x - 1}
					dir := DirMap[board[y][x]]
					pos := inititally.add(dir.mult(t)).mod(dy, dx).add(pos{y: 1, x: 1})
					_blizzards[pos] = true
				}
			}
		}
		blizzards = append(blizzards, _blizzards)
	}
	return blizzards
}

func partOne(board [][]byte, begin pos, end pos, blizzards []map[pos]bool) int {
	return bfs(board, begin, end, blizzards, 0)
}

func partTwo(board [][]byte, begin pos, end pos, blizzards []map[pos]bool) int {
	x := bfs(board, begin, end, blizzards, 0)
	// where we came from
	y := bfs(board, end, begin, blizzards, x)
	z := bfs(board, begin, end, blizzards, y)

	return z
}

func bfs(board [][]byte, begin pos, end pos, blizzards []map[pos]bool, startTime int) int {
	q := list.New()
	q.PushBack(state{pos: begin, time: startTime})
	seen := make(map[state]bool)
	for q.Len() != 0 {
		cur := q.Remove(q.Front()).(state)
		relativeTime := cur.time % len(blizzards)

		if cur.pos.y < 0 || cur.pos.y == len(board) || cur.pos.x < 0 || cur.pos.x == len(board[0]) {
			continue
		}
		if board[cur.pos.y][cur.pos.x] == '#' {
			continue
		}
		if blizzards[relativeTime][cur.pos] {
			continue
		}
		if seen[state{
			pos:  cur.pos,
			time: relativeTime,
		}] {
			continue
		}
		seen[state{
			pos:  cur.pos,
			time: relativeTime,
		}] = true

		if cur.pos == end {
			return cur.time
		}
		for _, dir := range DirMap {
			next := cur.pos.add(dir)
			q.PushBack(state{
				pos:  next,
				time: cur.time + 1,
			})
		}
		q.PushBack(state{
			pos:  cur.pos,
			time: cur.time + 1,
		})
	}
	return 0
}

func printBoard(board [][]byte) string {
	var b strings.Builder
	for _, row := range board {
		for _, cell := range row {
			b.WriteByte(cell)
		}
		b.WriteByte('\n')
	}
	return b.String()
}
