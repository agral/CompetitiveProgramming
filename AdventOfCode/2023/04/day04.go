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
	number   int
	winning  []int
	owned    []int
	matching int
	count    int
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
	for _, val := range w {
		num, _ := strconv.Atoi(val)
		winning = append(winning, num)
	}
	o := strings.Split(sOwned, " ")
	for _, val := range o {
		num, _ := strconv.Atoi(val)
		owned = append(owned, num)
	}
	// calculate the number of matching cards.
	// There's no intersection() / union() in golang; has to be done manually.
	matching := 0
	for _, w := range winning {
		for _, o := range owned {
			if w == o {
				matching += 1
			}
		}
	}
	return Card{cardno, winning, owned, matching, 1}
}

func (c Card) score() int {
	if c.matching == 0 {
		return 0
	}
	return 1 << (c.matching - 1)
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
	cards := []Card{} // cards[i] holds a Card with i+1 number; e.g. cards[0] is Card 1 and so on.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		c := makeCard(line)
		silver += c.score()
		cards = append(cards, c)
		// Make sure that the input consists of cards numbered from 1 to n,
		// in order. This is not guaranteed, but I want to rely on it.
		if len(cards) != c.number {

		}
	}
	for i := 0; i < len(cards); i++ {
		// the cards are so arranged that card_num + 1 + matching < len(cards)
		for j := i + 1; j < i+1+cards[i].matching; j++ {
			cards[j].count += cards[i].count
		}
		//fmt.Printf("There are %d cards #%d\n", cards[i].count, cards[i].number)
		gold += cards[i].count
	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
