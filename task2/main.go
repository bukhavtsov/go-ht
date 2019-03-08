package main

import (
	"fmt"
	"strconv"
)

func MapTo(numbers []int, makeWord func(elem int, _ int) string) (stringElements []string) {
	for index, number := range numbers {
		stringElements = append(stringElements, makeWord(number, index))
	}
	return
}

func makeWord(element, _ int) string {
	elementStr := strconv.Itoa(element)
	switch elementStr {
	case "1":
		return "one"
	case "2":
		return "two"
	case "3":
		return "three"
	case "4":
		return "four"
	case "5":
		return "five"
	case "6":
		return "six"
	case "7":
		return "seven"
	case "8":
		return "eight"
	case "9":
		return "nine"
	default:
		return "unknown"
	}
}

func Convert(numbers []int) []string {
	return MapTo(numbers, makeWord)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T", MapTo(numbers, makeWord))
}
