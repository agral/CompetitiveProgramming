package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

var nodes = make(map[string]Node)

func getLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
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

	scanner := bufio.NewScanner(file)
	directions := getLine(scanner)
	getLine(scanner)
	for scanner.Scan() {
		line := scanner.Text()
		line = line[:len(line)-1]
		split := strings.Split(line, " = (")
		leftRight := strings.Split(split[1], ", ")
		nodes[split[0]] = Node{leftRight[0], leftRight[1]}
	}
	silver, gold := 0, 1
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		if directions[silver%len(directions)] == 'L' {
			currentNode = nodes[currentNode].Left
		} else {
			currentNode = nodes[currentNode].Right
		}
		silver++
	}
	fmt.Printf("Silver: %d\n", silver)

	startNodes := []string{}
	for key := range nodes {
		if key[2] == 'A' {
			startNodes = append(startNodes, key)
		}
	}
	sz := len(startNodes)
	fmt.Println(startNodes)
	cycles := make([]int, sz)
	for i, node := range startNodes {
		step := 0
		n := node[:]
		for n[2] != 'Z' {
			if directions[step%len(directions)] == 'L' {
				n = nodes[n].Left
			} else {
				n = nodes[n].Right
			}
			step++
		}
		//fmt.Printf("Node %s: cycle found in %d steps\n", node, step)
		cycles[i] = step
		gold = lcm(gold, step)
	}

	fmt.Printf("Gold: %d\n", gold)
}
