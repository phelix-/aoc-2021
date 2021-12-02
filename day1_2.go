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

	windowSize := 3
	window := make([]int, windowSize)
	previousSum := 0
	transitions := 0

	for i := 0; scanner.Scan(); i++ {
		value, _ := strconv.Atoi(scanner.Text())
		window[i%windowSize] = value

		currentSum := 0
		if (i+1) >= windowSize {
			for j := 0; j < windowSize; j++ {
				currentSum += window[(i-j) % windowSize]
			}
		}

		if i >= windowSize && currentSum > previousSum {
			transitions++
		}
		previousSum = currentSum
	}
	log.Printf("transitions: %d", transitions)
}
