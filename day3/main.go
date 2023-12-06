package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Coord struct {
	X, Y int
}

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
	return strings.Split(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n")
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
	sum := 0
	symbols := make(map[Coord]rune)
	for idx, line := range input {
		for pos, char := range line {
			if !unicode.IsDigit(char) && char != '.' {
				symbols[Coord{pos, idx}] = char
			}
		}
	}
	// fmt.Println(symbols)
	re := regexp.MustCompile(`\d+`)
	for idx, line := range input {
		numbers := re.FindAllStringIndex(line, -1)
		// fmt.Println(numbers)
	Number:
		for i := 0; i < len(numbers); i++ {
			num, _ := strconv.Atoi(line[numbers[i][0]:numbers[i][1]])
			for j := numbers[i][0]; j < numbers[i][1]; j++ {
				if isAdjacent(Coord{j, idx}, symbols) {
					// fmt.Println(num, Coord{j, idx})
					sum += num
					continue Number
				}
			}
			// fmt.Printf("%d @ %d:%d is not near a symbol\n", num, idx, numbers[i][0])
		}
	}
	fmt.Println(sum)
}

func isAdjacent(pos Coord, symbols map[Coord]rune) bool {
	_, ok := symbols[Coord{pos.X, pos.Y - 1}]
	_, ok2 := symbols[Coord{pos.X, pos.Y + 1}]
	_, ok3 := symbols[Coord{pos.X - 1, pos.Y - 1}]
	_, ok4 := symbols[Coord{pos.X - 1, pos.Y}]
	_, ok5 := symbols[Coord{pos.X - 1, pos.Y + 1}]
	_, ok6 := symbols[Coord{pos.X + 1, pos.Y - 1}]
	_, ok7 := symbols[Coord{pos.X + 1, pos.Y}]
	_, ok8 := symbols[Coord{pos.X + 1, pos.Y + 1}]
	return (ok || ok2 || ok3 || ok4 || ok5 || ok6 || ok7 || ok8)
}

func part2(input []string) {
	symbols := make(map[Coord]rune)
	adjNums := make(map[Coord][]int)
	for idx, line := range input {
		for pos, char := range line {
			if char == '*' {
				symbols[Coord{pos, idx}] = char
				adjNums[Coord{pos, idx}] = make([]int, 0)
			}
		}
	}

	// fmt.Println(symbols)
	re := regexp.MustCompile(`\d+`)
	for idx, line := range input {
		numbers := re.FindAllStringIndex(line, -1)
		// fmt.Println(numbers)
	Number:
		for i := 0; i < len(numbers); i++ {
			num, _ := strconv.Atoi(line[numbers[i][0]:numbers[i][1]])
			for j := numbers[i][0]; j < numbers[i][1]; j++ {
				if loc, ok := isAdjacent2(Coord{j, idx}, symbols); ok {
					for _, s := range loc {
						adjNums[s] = append(adjNums[s], num)
					}
					// fmt.Println(num, Coord{j, idx})
					continue Number
				}
			}
			// fmt.Printf("%d @ %d:%d is not near a symbol\n", num, idx, numbers[i][0])
		}
	}
	ratio := 0
	for _, nums := range adjNums {
		if len(nums) == 2 {
			ratio += nums[0] * nums[1]
		}
	}
	fmt.Println(ratio)
}

func isAdjacent2(pos Coord, symbols map[Coord]rune) ([]Coord, bool) {
	out := make([]Coord, 0)
	if _, ok := symbols[Coord{pos.X, pos.Y - 1}]; ok {
		out = append(out, Coord{pos.X, pos.Y - 1})
	}
	if _, ok := symbols[Coord{pos.X, pos.Y + 1}]; ok {
		out = append(out, Coord{pos.X, pos.Y + 1})
	}
	if _, ok := symbols[Coord{pos.X - 1, pos.Y - 1}]; ok {
		out = append(out, Coord{pos.X - 1, pos.Y - 1})
	}
	if _, ok := symbols[Coord{pos.X - 1, pos.Y}]; ok {
		out = append(out, Coord{pos.X - 1, pos.Y})
	}
	if _, ok := symbols[Coord{pos.X - 1, pos.Y + 1}]; ok {
		out = append(out, Coord{pos.X - 1, pos.Y + 1})
	}
	if _, ok := symbols[Coord{pos.X + 1, pos.Y - 1}]; ok {
		out = append(out, Coord{pos.X + 1, pos.Y - 1})
	}
	if _, ok := symbols[Coord{pos.X + 1, pos.Y}]; ok {
		out = append(out, Coord{pos.X + 1, pos.Y})
	}
	if _, ok := symbols[Coord{pos.X + 1, pos.Y + 1}]; ok {
		out = append(out, Coord{pos.X + 1, pos.Y + 1})
	}
	if len(out) > 0 {
		return out, true
	}
	return nil, false
}
