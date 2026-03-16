package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, t := range temps {
		var key int
		key = int(t/10) * 10
		groups[key] = append(groups[key], t)
	}
	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for idx, k := range keys {
		var str []string
		for _, t := range groups[k] {
			str = append(str, strconv.FormatFloat(t, 'f', 1, 64))
		}
		if idx == 0 {
			fmt.Printf("%d:{%s}", k, strings.Join(str, ", "))
		} else {
			fmt.Printf(", %d:{%s}", k, strings.Join(str, ", "))
		}
	}
	fmt.Print("\n")
}
