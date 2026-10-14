package main

import "fmt"

func main() {
	words := []string{
		"abc",
		"fg",
		"hij",
		"klm",
		"nop",
		"qrs",
		"vwx",
	}
	// words2 := []string{
	// 	"b",
	// 	"b",
	// 	"b",
	// }
	target := 1620
	// target2 := 10
	res := productOfAlphabeticalSums(words, target)
	// res2 := productOfAlphabeticalSums(words2, target2)
	fmt.Println("response:", res)
	// fmt.Println("response:", res2)
}

func productOfAlphabeticalSums(words []string, target int) bool {
	sums := map[int]struct{}{}
	for _, word := range words {
		sum := alphabeticalSum(word)
		sums[sum] = struct{}{}
	}

	//at most 26*3 (78)
	for sum1 := range sums {
		if target%sum1 != 0 {
			continue
		}
		for sum2 := range sums {
			if target%(sum1*sum2) != 0 {
				continue
			}
			rest := target / (sum1 * sum2)
			_, ok := sums[rest]
			if !ok {
				continue
			}
			return true
		}
	}

	return false
}

func alphabeticalSum(word string) int {
	sum := 0
	for _, c := range word {
		pos := c - 96
		sum += int(pos)
	}

	return sum
}
