package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	position := 0
	depth := 0

	file, err := os.Open("./day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		command := split[0]
		amount, _ := strconv.Atoi(split[1])
		switch command {
		case "forward":
			position += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
			if depth < 0 {
				log.Printf("depth up to %d", depth)
			}
		}
	}
	log.Printf("horizontal: %d depth: %d total: %d", position, depth, position*depth)
}
