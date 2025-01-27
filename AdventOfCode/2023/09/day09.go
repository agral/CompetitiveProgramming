package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func predictNext(sequence []int) int {
	diff := make([]int, len(sequence)-1)
	isOnlyZeroes := true
	for i := 0; i < len(sequence)-1; i++ {
		diff[i] = sequence[i+1] - sequence[i]
		if diff[i] != 0 {
			isOnlyZeroes = false
		}
	}
	if isOnlyZeroes {
		return sequence[len(sequence)-1]
	} else {
		return sequence[len(sequence)-1] + predictNext(diff)
	}
}

func predictPrev(sequence []int) int {
	diff := make([]int, len(sequence)-1)
	isOnlyZeroes := true
	for i := 0; i < len(sequence)-1; i++ {
		diff[i] = sequence[i+1] - sequence[i]
		if diff[i] != 0 {
			isOnlyZeroes = false
		}
	}
	if isOnlyZeroes {
		return sequence[0]
	} else {
		return sequence[0] - predictPrev(diff)
	}
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
	for scanner.Scan() {
		line := scanner.Text()
		n := strings.Split(line, " ")
		nums := make([]int, len(n))
		for i, strNum := range n {
			nums[i], _ = strconv.Atoi(strNum)
		}

		silver += predictNext(nums)
		gold += predictPrev(nums)
	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
