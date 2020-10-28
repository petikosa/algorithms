package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

func main() {
	a := []int{5, 8, 2, 5, 4, 9, 7, 8, 5, 3, 1, 6}
	fmt.Println(selectSort(a))
}

func findMin(a []int) (int, int) {
	min := maxInt
	minIndex := 0
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
			minIndex = i
		}
	}
	return minIndex, min
}

func selectSort(a []int) []int {
	var b []int
	arrSize := len(a)
	for i := 0; i < arrSize; i++ {
		minIndex, min := findMin(a)
		b = append(b, min)
		// delete min element
		a = append(a[:minIndex], a[minIndex+1:]...)
	}
	return b
}
