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

func getLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getNumWaysToBeat(time int, dist int) int {
	ans := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		raceTime := time - holdTime
		if raceTime*holdTime > dist {
			ans++
		}
	}
	return ans
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

	oneOrMoreSpaces := regexp.MustCompile(" +")

	silver, gold := 1, 1
	scanner := bufio.NewScanner(file)
	strTimeFull := oneOrMoreSpaces.ReplaceAllString(getLine(scanner), " ")
	strDistFull := oneOrMoreSpaces.ReplaceAllString(getLine(scanner), " ")
	strTime := strings.Split(strTimeFull, " ")[1:]
	strDist := strings.Split(strDistFull, " ")[1:]

	for i := 0; i < len(strTime); i++ {
		time, _ := strconv.Atoi(strTime[i])
		dist, _ := strconv.Atoi(strDist[i])
		silver *= getNumWaysToBeat(time, dist)
	}

	fmt.Printf("Silver: %d\n", silver)
	fmt.Printf("Gold: %d\n", gold)
}
