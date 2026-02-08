package main

func main() {
	input := []int{0, 3, 2, 5, 4, 6, 1, 1}
	println("length longest consecutive sequence: ", longestConsecutive(input))
}

func longestConsecutive(nums []int) int {
	numToConsecutive := map[int]int{}

	//O(n)
	for _, num := range nums {
		numToConsecutive[num] = num + 1
	}

	//O(n)
	lenLongestSequence := 0
	for key, value := range numToConsecutive {
		if _, hasPrevious := numToConsecutive[key-1]; !hasPrevious {
			lenCurrentSequence := 1
			for {
				_, ok := numToConsecutive[value]
				if !ok {
					break
				}
				lenCurrentSequence++
				value++
			}

			if lenCurrentSequence > lenLongestSequence {
				lenLongestSequence = lenCurrentSequence
			}
		}
	}

	return lenLongestSequence
}
