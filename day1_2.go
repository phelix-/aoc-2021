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
	buffer := NewBuffer(windowSize)
	transitions := 0

	for i := 0; scanner.Scan(); i++ {
		value, _ := strconv.Atoi(scanner.Text())

		previousSum := buffer.Sum()
		buffer.Push(value)

		if i >= windowSize && buffer.Sum() > previousSum {
			transitions++
		}
	}
	log.Printf("transitions: %d", transitions)
}


type Buffer struct {
	buffer []int
	size int
	capacity int
}

func NewBuffer(capacity int) *Buffer {
	return &Buffer{
		buffer: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (b *Buffer) Push(i int) {
	b.buffer[b.size % b.capacity] = i
	b.size++
}

func (b Buffer) Sum() int {
	sum := 0
	for i := 0; i < b.size && i < b.capacity; i++ {
		sum += b.buffer[i]
	}
	return sum
}