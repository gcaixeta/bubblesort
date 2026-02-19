package main

import "fmt"

func Bubblesort(array []int) []int {
	a := make([]int, len(array))
	copy(a, array)

	lastSwappedIndex := len(array) - 1

	for _ = range len(array) {
		swapped := false

		for j := range lastSwappedIndex {
			if array[j] > array[j+1] {
				aux := array[j+1]
				array[j+1] = array[j]
				array[j] = aux
				swapped = true
				lastSwappedIndex = j + 1
			}
		}

		if !swapped {
			break
		}
	}

	fmt.Printf("Array original: %v\n", a)
	fmt.Printf("Array ordenado: %v\n", array)
	return array
}

func main() {
	testArray := []int{1, 2, 5, 2, 4, 7}

	Bubblesort(testArray)
}
