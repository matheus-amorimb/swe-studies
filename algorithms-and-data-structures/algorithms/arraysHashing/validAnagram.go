package main

func main() {
	s, t := "racecaw", "carrace"
	println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// complexity O(n)
	countLettersWord1, countLettersWord2 := map[rune]int{}, map[rune]int{}
	for idx, ch := range s {
		countLettersWord1[ch]++
		countLettersWord2[rune(t[idx])]++
	}

	// complexity O(k), with k <= n
	// therefore, overall complexity is still O(n)
	for key, value := range countLettersWord1 {
		if countLettersWord2[key] != value {
			return false
		}
	}

	return true
}
