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
	value, ok := wordsMap[elementStr]
	if ok {
		return value
	}
	return "unknown"
}

var wordsMap = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func Convert(numbers []int) []string {
	return MapTo(numbers, makeWord)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Convert(numbers))
}
