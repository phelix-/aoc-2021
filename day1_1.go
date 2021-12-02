package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	transitions := 0
	scanner.Scan()
	previousValue, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		if value > previousValue {
			transitions++
		}
		previousValue = value
	}
	log.Printf("transitions: %d", transitions)
}
