package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func readFile() []string {
	fh, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fh), "\n")
	lines = lines[:len(lines)-1] // remove the last line because its empty
	return lines
}

func useTestData() []string {
	return strings.Split(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`, "\n")
}

func main() {
	var lines []string
	if len(os.Args) > 2 && os.Args[2] == "d" {
		fmt.Println("==! Using test data !==")
		lines = useTestData()
	} else {
		lines = readFile()
	}
	if len(os.Args) > 1 && os.Args[1] == "2" {
		fmt.Println("Part 2")
		part2(lines)
	} else {
		fmt.Println("Part 1")
		part1(lines)
	}
}

func part1(input []string) {
	parsed := make([][]int, len(input))
	re := regexp.MustCompile(`\d+`)
	total := 0
	for idx, line := range input {
		m := re.FindAllString(line, -1)
		// fmt.Printf("m: %v\n", m)
		parsed[idx] = make([]int, len(m)-1) // 35
		for i, num := range m[1:] {
			parsed[idx][i], _ = strconv.Atoi(num)
		}
		slices.Sort(parsed[idx])
		matches := 0
		for i := 0; i < len(parsed[idx])-1; i++ {
			if parsed[idx][i] == parsed[idx][i+1] {
				matches++
			}
		}
		if matches > 0 {
			total += 1 << (matches - 1)
		}
	}
	fmt.Printf("total: %v\n", total)
}

func part2(input []string) {

}
