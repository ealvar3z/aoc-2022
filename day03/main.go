package main

import (
	"fmt"

	"github.com/ealvar3z/aoc2022/lib/aoc"
	"github.com/ealvar3z/aoc2022/lib/file"
)

func main() {
	path, complete := aoc.Setup(2022, 3, true)
	defer complete()

	lines := file.TextSplit(path, " ")

	fmt.Println(lines)
}
