package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

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
	}
	fmt.Printf("Silver: %d\n", silver)

}
