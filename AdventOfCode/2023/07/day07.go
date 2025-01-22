package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var CARD_SENIORITY = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type HandType int

const (
	NOT_PROCESSED   HandType = iota
	HIGH_CARD                // e.g. 23456
	ONE_PAIR                 // e.g. A23A4
	TWO_PAIR                 // e.g. 23432
	THREE_OF_A_KIND          // e.g. TTT98
	FULL_HOUSE               // e.g. 23332
	FOUR_OF_A_KIND           // e.g. AA8AA
	FIVE_OF_A_KIND           // e.g. AAAAA
)

type Deal struct {
	Hand string
	Bid  int
	Type HandType
}

func MakeDeal(hand string, bid int) Deal {
	// Determine the instance's HandType:
	handType := HIGH_CARD
	counts := make(map[byte]int)
	for i := 0; i < len(hand); i++ {
		counts[hand[i]]++
	}
	combos := make([]int, 6)
	for _, v := range counts {
		combos[v]++
	}
	if combos[5] == 1 {
		handType = FIVE_OF_A_KIND
	} else if combos[4] == 1 {
		handType = FOUR_OF_A_KIND
	} else if combos[3] == 1 {
		if combos[2] == 1 {
			handType = FULL_HOUSE
		} else {
			handType = THREE_OF_A_KIND
		}
	} else if combos[2] == 2 {
		handType = TWO_PAIR
	} else if combos[2] == 1 {
		handType = ONE_PAIR
	}

	return Deal{
		Hand: hand,
		Bid:  bid,
		Type: handType,
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
	deals := []Deal{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])
		deals = append(deals, MakeDeal(hand, bid))
	}
	sort.Slice(deals, func(i, j int) bool {
		a, b := deals[i], deals[j]
		if a.Type == b.Type {
			for i := 0; i < len(a.Hand); i++ {
				if a.Hand[i] != b.Hand[i] {
					return CARD_SENIORITY[a.Hand[i]] < CARD_SENIORITY[b.Hand[i]]
				}
			}
		}
		return a.Type < b.Type
	})
	for i, deal := range deals {
		silver += (i + 1) * deal.Bid
	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
