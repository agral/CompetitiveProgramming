package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing argument - input file. Aborting.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	silver := 0
	gold := 0
	digits := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	for scanner.Scan() {
		line := scanner.Text()

		first := 0
		last := len(line) - 1
		for {
			if line[first] >= '0' && line[first] <= '9' {
				break
			}
			first++
		}
		for {
			if line[last] >= '0' && line[last] <= '9' {
				break
			}
			last--
		}

		tens := int(line[first] - '0')
		ones := int(line[last] - '0')
		silver += (10*tens + ones)

		for k, v := range digits {
			firstIdx := strings.Index(line, k)
			lastIdx := strings.LastIndex(line, k)
			if firstIdx >= 0 && firstIdx < first {
				first = firstIdx
				tens = v
			}
			if lastIdx > last {
				last = lastIdx
				ones = v
			}
		}
		gold += (10*tens + ones)

	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
