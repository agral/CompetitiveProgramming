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
var CARD_SENIORITY_GOLD = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
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
	Hand     string
	Bid      int
	Type     HandType
	GoldType HandType
}

func MakeDeal(hand string, bid int) Deal {
	// Determine the instance's HandType (no jokers introduced):
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

	// Determine the instance's "gold" HandType (new joker rule):
	goldType := HIGH_CARD
	goldCounts := make(map[byte]int)
	for i := 0; i < len(hand); i++ {
		goldCounts[hand[i]]++
	}
	goldCombos := make([]int, 6)
	for k, v := range goldCounts {
		if k != 'J' {
			goldCombos[v]++
		}
	}

	if goldCombos[5] == 1 {
		goldType = FIVE_OF_A_KIND
	} else if goldCombos[4] == 1 {
		goldType = FOUR_OF_A_KIND
	} else if goldCombos[3] == 1 {
		if goldCombos[2] == 1 {
			goldType = FULL_HOUSE
		} else {
			goldType = THREE_OF_A_KIND
		}
	} else if goldCombos[2] == 2 {
		goldType = TWO_PAIR
	} else if goldCombos[2] == 1 {
		goldType = ONE_PAIR
	}

	for i := 0; i < goldCounts['J']; i++ {
		// Joker ascension rules:
		switch goldType {
		case HIGH_CARD:
			goldType = ONE_PAIR // from e.g. 789TJ to e.g. 789TT, best result is one pair.
		case ONE_PAIR:
			// from e.g. TTJQK can go either two pairs (TTQQK) or three-of-a-kind (TTTQK),
			goldType = THREE_OF_A_KIND // and three-of-a-kind is best.
		case TWO_PAIR:
			goldType = FULL_HOUSE // from e.g. TTJQQ best result is TTQQQ
		case THREE_OF_A_KIND:
			goldType = FOUR_OF_A_KIND // from e.g. TTTJQ best result is TTTTQ
		case FULL_HOUSE: // a full house e.g. TTTJJ or TTJJJ will not be reported as a full house,
			// because J cards are not counted for goldCounts. So, seeing any ascension from a FULL_HOUSE
			// would be a logic error.
			fmt.Printf("Hand: %s\n", hand)
			panic("Ascension from full house should never happen!")
		case FOUR_OF_A_KIND: // from e.g. TTJTT to TTTTT
			goldType = FIVE_OF_A_KIND
		case FIVE_OF_A_KIND: // the only ascension that can happen from FIVE_OF_A_KIND
			// Due to the way converting hands to hand types is realized,
			// a hand of exclusively "JJJJJ" is initially sorted as HIGH_CARD, and then
			// ascended five times: to ONE_PAIR (1), THREE_OF_A_KIND (2), FOUR_OF_A_KIND (3)
			// and then twice to FIVE_OF_A_KIND (4, 5). This is valid only if starting from five jacks;
			// any other five-of-a-kind can not ascend.
			if hand == "JJJJJ" {
				goldType = FIVE_OF_A_KIND // won't get better
			}
			if hand != "JJJJJ" {
				fmt.Printf("Hand: %s\n", hand)
				panic("Ascension from five-of-a-kind that is not \"JJJJJ\" should never happen")
			}
		}
	}

	return Deal{
		Hand:     hand,
		Bid:      bid,
		Type:     handType,
		GoldType: goldType,
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

	// Gold part is same as silver, but sort by joker rule (GoldType instead of Type)
	sort.Slice(deals, func(i, j int) bool {
		a, b := deals[i], deals[j]
		if a.GoldType == b.GoldType {
			for i := 0; i < len(a.Hand); i++ {
				if a.Hand[i] != b.Hand[i] {
					return CARD_SENIORITY_GOLD[a.Hand[i]] < CARD_SENIORITY_GOLD[b.Hand[i]]
				}
			}
		}
		return a.GoldType < b.GoldType
	})
	for i, deal := range deals {
		gold += (i + 1) * deal.Bid
	}
	fmt.Printf("Gold: %d\n", gold)
}
