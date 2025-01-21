package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Golang does not allow constant slices; so don't modify it. Could be a function for enterprise code.
var CONVERSIONS = []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light",
	"light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
var NUM_CONVERSIONS = len(CONVERSIONS)

var almanac [][][]int

func getLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func processMap(scanner *bufio.Scanner) {
	_ = getLine(scanner) // skip past first line; "%s-to-%s map:" - it is redundant.
	mapping := [][]int{}
	for line := getLine(scanner); len(line) > 0; line = getLine(scanner) {
		entry := []int{}
		for _, strNum := range strings.Split(line, " ") {
			n, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatalf("Malformed map line: %s", line)
			}
			entry = append(entry, n)
		}
		mapping = append(mapping, entry)
	}
	almanac = append(almanac, mapping)
}

func translate(input int, mapping_num int) int {
	// Destination range start | Source range start | Range length
	for _, entry := range almanac[mapping_num] {
		if entry[1] <= input && entry[1]+entry[2] >= input {
			return entry[0] + (input - entry[1])
		}
	}
	// "Any source numbers that aren't mapped correspond to the same destination number."
	return input
}

func getLocation(seed int) int {
	for entry_num := range CONVERSIONS {
		seed = translate(seed, entry_num)
	}
	return seed
}

func main() {
	inputFile := "example.txt"
	if len(os.Args) < 2 {
		fmt.Printf("Input file not specified. Using %q.\n", inputFile)
	} else {
		inputFile = os.Args[1]
	}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	silver, gold := 0, 0

	scanner := bufio.NewScanner(file)
	seedsLine := getLine(scanner) // Process the seeds to an int slice:
	strSeeds := strings.Split(seedsLine, ": ")[1]
	seeds := []int{}
	for _, s := range strings.Split(strSeeds, " ") {
		n, _ := strconv.Atoi(s)
		seeds = append(seeds, n)
	}
	_ = getLine(scanner) // skip past an empty line

	for range CONVERSIONS {
		processMap(scanner)
	}

	silver = math.MaxInt
	for _, seed := range seeds {
		silver = min(silver, getLocation(seed))
	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
	//fmt.Println(almanac)
}
