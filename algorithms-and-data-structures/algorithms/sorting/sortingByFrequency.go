package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	arr := "supercalifragilisticexpialidocious"
	fmt.Println(sortingByFreq(arr))
}

// As the frequency map only ever has 26 keys, sorting it takes O(26 log 26) = O(1) time. the actual bottleneck is
// building the frequency map, which takes O(n) time.
func sortingByFreq(str string) []string {
	arr := []string{}
	freqMap := map[string]int{}

	for _, char := range str {
		s := string(char)
		if freqMap[s] == 0 {
			arr = append(arr, s)
		}
		freqMap[s]++
	}

	byFreq := func(a, b string) int {
		return cmp.Or(
			cmp.Compare(freqMap[b], freqMap[a]),
			cmp.Compare(a, b),
		)
	}

	slices.SortFunc(arr, byFreq)

	return arr
}
