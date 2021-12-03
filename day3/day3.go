package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1() int64 {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	gamma := CalcGamma(input)
	epsilon := CalcEpsilon(input)

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaInt * epsilonInt
}

func Part2() int64 {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}


	inputCopy := input
	o2Rating := FindOxygenGeneratorRating(input)
	co2Rating := FindCo2ScrubberRating(inputCopy)
	o2RatingInt, _ := strconv.ParseInt(o2Rating, 2, 64)
	co2RatingInt, _ := strconv.ParseInt(co2Rating, 2, 64)

	return o2RatingInt * co2RatingInt
}


func CalcGamma(input []string) string {
	onesCount := make([]int, len(input[0]))
	for _, currentInput := range input {
		for i := 0; i < len(currentInput); i++ {
			if currentInput[i] == '1' {
				onesCount[i] = onesCount[i] + 1
			}
		}
	}
	gamma := ""
	for i := 0; i < len(onesCount); i++ {
		if onesCount[i] >= len(input) / 2 {
			gamma = gamma + "1"
		} else {
			gamma = gamma + "0"
		}
	}
	return gamma
}

func CalcEpsilon(input []string) string {
	gamma := CalcGamma(input)
	epsilon := ""
	for i := 0; i < len(gamma); i++ {
		if gamma[i] == '1' {
			epsilon = epsilon + "0"
		} else {
			epsilon = epsilon + "1"
		}
	}
	return epsilon
}

func FindOxygenGeneratorRating(input []string) string {
	for i := 0; i < len(input[0]); i++ {
		gamma := CalcGamma(input)
		candidates := make([]string, 0)
		for _, possibleValue := range input {
			if possibleValue[i] == gamma[i] {
				candidates = append(candidates, possibleValue)
			}
		}
		if len(candidates) == 1 {
			return candidates[0]
		} else {
			input = candidates
		}
	}
	log.Fatal("Ended with too many numbers", input)
	return ""
}

func FindCo2ScrubberRating(input []string) string {
	for i := 0; i < len(input[0]); i++ {
		epsilon := CalcEpsilon(input)
		candidates := make([]string, 0)
		for _, possibleValue := range input {
			if possibleValue[i] == epsilon[i] {
				candidates = append(candidates, possibleValue)
			}
		}
		if len(candidates) == 1 {
			return candidates[0]
		} else {
			input = candidates
		}
	}
	log.Fatal("Ended with too many numbers", input)
	return ""
}