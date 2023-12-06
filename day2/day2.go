package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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
	return strings.Split(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, "\n")
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

func part1(lines []string) {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	possible := 0
	re := regexp.MustCompile("(\\d+) (\\w+)")
Game:
	for idx, line := range lines {
		rounds := strings.Split(line, ";")
		for _, round := range rounds {
			marbles := re.FindAllStringSubmatch(round, -1)
			if marbles == nil {
				log.Fatalf("No match for game %d", idx+1)
			}
			// fmt.Println(marbles)
			for _, pair := range marbles {
				count, _ := strconv.Atoi(pair[1])
				color := pair[2]
				if max[color] < count {
					// fmt.Println(idx+1, color, count)
					continue Game
				}
			}
		}
		possible += idx + 1
	}
	fmt.Println(possible)
}

func part2(lines []string) {
	psum := 0
	re := regexp.MustCompile("(\\d+) (\\w+)")

	for idx, line := range lines {
		req := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		rounds := strings.Split(line, ";")
		for _, round := range rounds {
			marbles := re.FindAllStringSubmatch(round, -1)
			if marbles == nil {
				log.Fatalf("No match for game %d", idx+1)
			}
			// fmt.Println(marbles)
			for _, pair := range marbles {
				count, _ := strconv.Atoi(pair[1])
				color := pair[2]
				req[color] = max(req[color], count)
			}
		}
		power := max(1, req["red"]) * max(1, req["blue"]) * max(1, req["green"])
		fmt.Println(idx+1, power)
		psum += power
	}
	fmt.Println(psum)
}
