package main

import "fmt"

func main() {
	a := []int{0, 2, 7, 12, 24, 27, 29, 31, 36, 39, 40, 42, 46, 49, 50, 52, 57, 72, 84}
	index := binarySearch(a, 36)
	fmt.Println(index)
}

func binarySearch(a []int, x int) int {
	low, high := 0, len(a)-1
	mid := (low + high) / 2

	for low <= high {
		mid = (low + high) / 2
		if x == a[mid] {
			return mid
		} else if x < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
