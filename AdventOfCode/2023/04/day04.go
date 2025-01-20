package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	number  int
	winning []int
	owned   []int
}

func makeCard(input string) Card {
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	space := regexp.MustCompile(" +")
	s := space.ReplaceAllString(input, " ")
	cardno_and_rest := strings.Split(s, ": ")
	cardno, _ := strconv.Atoi(strings.Split(cardno_and_rest[0], "Card ")[1])
	parts := strings.Split(cardno_and_rest[1], " | ")
	winning := []int{}
	owned := []int{}
	sWinning, sOwned := parts[0], parts[1]
	w := strings.Split(sWinning, " ")
	fmt.Printf("%q\n", sWinning)
	for _, val := range w {
		num, _ := strconv.Atoi(val)
		winning = append(winning, num)
	}
	o := strings.Split(sOwned, " ")
	for _, val := range o {
		num, _ := strconv.Atoi(val)
		owned = append(owned, num)
	}
	return Card{cardno, winning, owned}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing the input file. Aborting.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	silver, gold := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Println(makeCard(line))
	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
