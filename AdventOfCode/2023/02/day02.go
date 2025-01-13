package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing the input file. Aborting.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	silver := 0
	gold := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		sGameNo, sGamesAll := split[0], split[1]
		numGame, err := strconv.Atoi(strings.Split(sGameNo, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		// silver: Can this game be realized with a bag of 12 red, 13 green, and 14 blue cubes?
		isValid := true

		sGames := strings.Split(sGamesAll, "; ")
		for _, sGame := range sGames {
			sCubes := strings.Split(sGame, ", ")
			for _, sCube := range sCubes {
				split := strings.Split(sCube, " ")
				count, err := strconv.Atoi(split[0])
				if err != nil {
					log.Fatal(err)
				}
				switch split[1] {
				case "red":
					if count > 12 {
						isValid = false
					}
				case "green":
					if count > 13 {
						isValid = false
					}
				case "blue":
					if count > 14 {
						isValid = false
					}
				default: // color other than red, green and blue
					isValid = false
				}
			}
		}
		if isValid {
			silver += numGame
		}

	}
	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
