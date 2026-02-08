package main

func main() {
	input := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	// expected output:
	// [["hat"],["act", "cat"],["stop", "pots", "tops"]]
	anagramsGrouped := groupAnagrams(input)
	for _, anagram := range anagramsGrouped {
		for _, str := range anagram {
			print(str, ", ")
		}
		println()
	}
}

func groupAnagrams(strs []string) [][]string {
	anagramsMap := map[[26]int][]string{}
	for _, str := range strs {
		var arrayKey [26]int
		for _, char := range str {
			// idx := int(char) - 97
			idx := char - 'a'
			arrayKey[idx]++
		}
		anagramsMap[arrayKey] = append(anagramsMap[arrayKey], str)
	}

	anagramsArray := [][]string{}
	for _, value := range anagramsMap {
		anagramsArray = append(anagramsArray, value)
	}

	return anagramsArray
}
