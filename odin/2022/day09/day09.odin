package main

import "core:fmt"
import "core:strings"
import "core:strconv"
import "core:math"

Position :: [2]int


main :: proc() {
	data := #load("input.txt")
	parse_input(data)
}

parse_input :: proc(data: []u8) {
		directions := map[rune]Position{
				'U' = Position{  0,  1 },
				'D' = Position{  0, -1 },
				'R' = Position{  1,  0 },
				'L' = Position{ -1,  0 },
		}

		knots := [10]Position{
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
				Position{ 0, 0 },
		}

		tails : [dynamic]Position
		defer delete(tails)

		knot9s : [dynamic]Position
		defer delete(knot9s)

		append(&tails, knots[1]) // part1
		append(&knot9s, knots[9]) // part2

		it := string(data)
		for line in strings.split_lines_iterator(&it) {
				parts := strings.split(line, " ")
				direction := rune(parts[0][0])
				steps := strconv.atoi(parts[1])
				for i in 0..<steps {
					knots[0] += directions[direction]
					for i in 1..<len(knots) {
						if !head_is_near_tails(&knots[i-1], &knots[i]) {
							move_tail(&knots[i-1], &knots[i])
						} else {
							// stop moving
							break
						}
					}
					add_location(&tails, &knots[1])
					add_location(&knot9s, &knots[9])
				}
		}
		fmt.printf("%d\n", len(tails))
		fmt.printf("%d\n", len(knot9s))
}

head_is_near_tails :: proc(head, tail: ^Position) -> bool {
		if math.abs(head.x - tail.x) >= 2 ||
				math.abs(head.y - tail.y) >= 2 {
						return false
				}
		return true
}

move_tail :: proc(head, tail: ^Position) {
		diff := head^ - tail^
		delta := Position{0, 0}
		if diff.x > 0 {
				delta.x += 1
		} else if diff.x < 0 {
				delta.x -= 1
		}
		if diff.y > 0 {
				delta.y += 1
		} else if diff.y < 0 {
				delta.y -= 1
		}

		tail^ += delta
}


add_location :: proc(locs: ^[dynamic]Position, tail: ^Position) {
		for loc in locs {
				if loc.x == tail.x && loc.y == tail.y {
						// boom! it's already in there!
						return
				}
		}
		append(locs, tail^)
}
