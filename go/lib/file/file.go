package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readText(s *bufio.Scanner, f func(input string)) {
	line := s.Text()
	f(line)
}

func readFile(fpath string, f func(input string)) {
	file, err := os.Open(fpath)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	s := bufio.NewScanner(file)

	for s.Scan() {
		readText(s, f)
	}
}

func ToText(fpath string) string {
	var response []string
	readFile(fpath, func(input string) {
		response = append(response, input)
	})

	return strings.Join(response, "\n")
}

func TextLines(fpath string) []string {
	var response []string
	readFile(fpath, func(input string) {
		response = append(response, input)
	})

	return response
}

func ToNum(fpath string) []int {
	var response []int
	readFile(fpath, func(input string) {
		num, _ := strconv.Atoi(input)
		response = append(response, num)
	})

	return response
}

func TextSplit(fpath, split string) [][]string {
	var response [][]string
	readFile(fpath, func(input string) {
		var line []string

		for _, val := range strings.Split(input, split) {
			line = append(line, val)
		}
		response = append(response, line)
	})

	return response
}

func NumSplit(fpath, split string) [][]int {
	var response [][]int

	readFile(fpath, func(input string) {
		line := []int{}

		for _, val := range strings.Split(input, split) {
			num, _ := strconv.Atoi(val)
			line = append(line, num)
		}

		response = append(response, line)
	})

	return response
}
