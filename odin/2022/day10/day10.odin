package main

import "core:fmt"
import "core:strings"
import "core:strconv"


main :: proc() {
	input := #load("input.txt")
	solve(input)
}


solve :: proc(input: []u8) {
	// single CPU register w/ default value of 1
	X := 1

	// clock cycle's constant rate is a tick
	cycles: [dynamic]int
	defer delete(cycles)

	it := string(input)
	for line in strings.split_lines_iterator(&it) {
		parts := strings.split(line, " ")
		instructions : = parts[0]
		if strings.compare(instructions, "noop") == 0 {
			append(&cycles, X)
		} else {
			// addx instruction takes (2) cycles
			for i in 0..<2 {
				append(&cycles, X)
			}
			add := strconv.atoi(parts[1])
			X += add
		}
	}

	signal_cycles := [6]int{
			20,
			60,
			100,
			140,
			180,
			220,
	}

	sum := 0
	for v in signal_cycles {
		strength := signal_strength(&cycles, v)
		sum += strength
	}
	fmt.println(sum)
	
	// part two
	for i in 0..<len(cycles) {
		// the CRT renders pixes (40) wide
		pixel := i % 40
		if pixel == 0 {
			// grab more input
			fmt.println()
		}
		fmt.print(get_pixels(pixel, cycles[i]))
	}
}

// signal strength (the cycle number multiplied by the value of the X register)
signal_strength :: proc(cycles: ^[dynamic]int, value: int) -> int {
	return value * cycles[value - 1]
}

get_pixels :: proc(cycle, pixel: int) -> rune {
	// left pixel (pixel-1) cycle right pixel (pixel+1)
	if pixel-1 <= cycle && pixel+1 >= cycle {
		return '#'
	}
	return '.'
}
