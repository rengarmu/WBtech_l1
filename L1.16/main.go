package main

import "fmt"

func main() {

	s := []int{1, 7, 59, 346, 76, 2, 29, 1, 8, 478, 59, 1, 72, 6, 20, 57}
	sortSlice := quickSort(s)
	fmt.Println(sortSlice)
}

func quickSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	pivot := s[0]
	less := make([]int, 0)
	middle := make([]int, 0)
	greater := make([]int, 0)
	for _, v := range s {
		if v < pivot {
			less = append(less, v)
		} else if v == pivot {
			middle = append(middle, v)
		} else {
			greater = append(greater, v)
		}

	}
	less = quickSort(less)
	greater = quickSort(greater)
	return append(append(less, middle...), greater...)
}
