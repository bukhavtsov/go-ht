package main

import "fmt"

func Filter(numbers []int, predicate func(int, int) bool) (arr []int) {
	for _, element := range numbers {
		if predicate(element, element+1) {
			arr = append(arr, element)
		}
	}
	return arr
}
func isEvenNumber(a int, b int) bool {
	return a%2 == 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Filter(numbers, isEvenNumber))
}
