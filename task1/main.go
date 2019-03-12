package main

import "fmt"

func Filter(numbers []int, predicate func(int, int) bool) (arr []int) {
	for index, element := range numbers {
		if predicate(element, index) {
			arr = append(arr, element)
		}
	}
	return arr
}
func equals(item int, index int) bool {
	return item == index
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Filter(numbers, equals))
}
