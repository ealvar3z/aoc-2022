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
	f := "input.txt"
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
	top3 := totalCalories[0] + totalCalories[1] + totalCalories[2]
	fmt.Println(top3)
}
