package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Setup gets us up and running for AoC each night. Its return value (a func)
// ought to be referred. Thus, resulting in a log of the task and the execution
// time.
func Setup(yr, day int, ex bool) (string, func()) {
	start := time.Now()

	dayName := fmt.Sprintf("day%.2d", day)
	path, _ := os.Getwd()

	_ = os.Chdir(filepath.Join(path, fmt.Sprintf("%d", yr), dayName))

	fn := fmt.Sprintf("%s.txt", dayName)

	if ex {
		printExFlag()
		fn = fmt.Sprintf("%s_example.txt", dayName)
	}

	aocRepo := os.Getenv("AOCDIR")
	inputPath := filepath.Join(fmt.Sprintf("%s", aocRepo), "input", strconv.Itoa(yr), fn)
	if aocRepo == "" {
		fmt.Fprintf(os.Stderr, "ENV var for AOCDIR is not set")
	}
	fmt.Println()

	return inputPath, func() {
		fmt.Printf("(%d:%d): %+v\n", yr, day, time.Since(start))

		if ex {
			printExFlag()
		}
	}
}

// PrintAnswer() logs the answer
func PrintAnswer(part int, answer interface{}) {
	fmt.Printf("part = %d, answer = %+v\n", part, answer)
}

func printExFlag() {
	fmt.Print("\n########### ###########\n# EXAMPLE # # EXAMPLE #\n########### ###########\n")
}
