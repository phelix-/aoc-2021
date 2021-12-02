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
	aim := 0

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
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	log.Printf("horizontal: %d depth: %d total: %d", position, depth, position*depth)
}
