package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const EMPTY byte = '.'

// schematic is a safe copy of input data with empty border of size 1.
var schematic = []string{}

func isSymbol(b byte) bool {
	return b != EMPTY && !(b >= '0' && b <= '9')
}

func getBorderingSymbol(y int, x int, length int) byte {
	boundaryCoordinates := [][2]int{{y - 1, x - 1}, {y, x - 1}, {y + 1, x - 1},
		{y - 1, x + length}, {y, x + length}, {y + 1, x + length}}
	for _, b := range boundaryCoordinates {
		if isSymbol(schematic[b[0]][b[1]]) {
			return schematic[b[0]][b[1]]
		}
	}

	for xx := x; xx < x+length; xx++ {
		if isSymbol(schematic[y-1][xx]) {
			return schematic[y-1][xx]
		}
		if isSymbol(schematic[y+1][xx]) {
			return schematic[y+1][xx]
		}
	}
	return EMPTY
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, "."+line+".")
	}
	H := len(schematic)
	W := len(schematic[0]) - 2
	schematic = append([]string{strings.Repeat(".", W+2)}, schematic...) // prepend
	schematic = append(schematic, strings.Repeat(".", W+2))              // append

	silver, gold := 0, 0
	for h := 1; h <= H; h++ {
		for w := 1; w <= W; w++ {
			if schematic[h][w] >= '0' && schematic[h][w] <= '9' {
				n := 0
				length := 0
				startH, startW := h, w
				for schematic[h][w] >= '0' && schematic[h][w] <= '9' {
					n *= 10
					n += int(schematic[h][w] - '0')
					length += 1
					w++
				}

				// Does this number have a neighboring symbol, that is not a digit nor "."?
				// If so, it's a part number, and should be added to silver.
				s := getBorderingSymbol(startH, startW, length)
				if s != EMPTY {
					silver += n
				} else {
				}
			}
		}
	}

	for h := 1; h <= H; h++ {
		for w := 0; w <= W; w++ {
			if schematic[h][w] != '.' && (schematic[h][w] < '0' || schematic[h][w] > '9') {
				if h > 0 && w > 0 {

				}
			}
		}
	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
