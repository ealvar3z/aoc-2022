package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var err error
	var file *os.File
	f := "../inputs/day1.txt"
	if file, err = os.Open(f); err != nil {
		fmt.Println("Oh oh, can't open file", file)
		return
	}
	defer file.Close()

	var totalCalories []int
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totalCalories = append(totalCalories, sum)
			sum = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Oh oh, can't convert line to int", line)
			return
		}
		sum += calories
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalCalories)))
	part1 := totalCalories[0]
	part2 := totalCalories[0] + totalCalories[1] + totalCalories[3]
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
