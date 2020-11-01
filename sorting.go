package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

func main() {
	a := []int{5, 8, 2, 5, 4, 9, 7, 8, 5, 3, 1, 6}
	mergeSort(a, 0, 11)
	fmt.Println(a)
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

func insertSort(a []int) []int {
	arrSize := len(a)
	for i := 1; i < arrSize; i++ {
		j := i - 1
		update := false
		for a[i] < a[j] {
			update = true
			j--
			if j < 0 {
				j = -1
				break
			}
		}
		if update {
			// insert
			a = append(a[:j+1], append([]int{a[i]}, a[j+1:]...)...)
			// delete
			a = append(a[:i+1], a[i+2:]...)
		}
	}
	return a
}

func bubbleSort(a []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				swapped = true
			}
		}
	}
	return a
}

func partition(a []int, lo int, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func quickSort(a []int, lo int, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
}

func merge(a []int, l int, m int, r int) {
	n1, n2 := m-l+1, r-m
	var left = make([]int, n1)
	var right = make([]int, n2)
	copy(left, a[l:m+1])
	copy(right, a[m+1:r+1])
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
		k++
	}
	for i < n1 {
		a[k] = left[i]
		i++
		k++
	}
	for j < n2 {
		a[k] = right[j]
		j++
		k++
	}
}

func mergeSort(a []int, l int, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(a, l, m)
		mergeSort(a, m+1, r)
		merge(a, l, m, r)
	}
}
