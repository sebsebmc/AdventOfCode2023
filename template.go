package main

import (
	"fmt"
	"log"
	"os"
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
	return strings.Split(``, "\n")
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

}

func part2(input []string) {

}
