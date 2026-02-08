package main

import "sort"

func main() {
	nums := []int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}
	k := 3
	output := topKFrequent(nums, k)
	for _, elem := range output {
		println(elem, ", ")
	}
}

func topKFrequentNotOptimal(nums []int, k int) []int {
	frequencyMap := map[int]int{}
	//O(n)
	for _, num := range nums {
		frequencyMap[num]++
	}

	var topKFrequent []int
	for num, _ := range frequencyMap {
		topKFrequent = append(topKFrequent, num)
	}

	sort.Slice(topKFrequent, func(i int, j int) bool {
		return frequencyMap[topKFrequent[i]] > frequencyMap[topKFrequent[j]]
	})

	return topKFrequent[:k]
}

func topKFrequent(nums []int, k int) []int {
	frequencyMap := map[int]int{}
	for _, num := range nums {
		frequencyMap[num]++
	}

	bucketMap := map[int][]int{}
	for num, numFreq := range frequencyMap {
		bucketMap[numFreq] = append(bucketMap[numFreq], num)
	}

	topKFrequent := []int{}
	for i := len(nums); i >= 0; i-- {
		for _, num := range bucketMap[i] {
			topKFrequent = append(topKFrequent, num)
			if len(topKFrequent) == k {
				return topKFrequent
			}
		}
	}

	return topKFrequent
}
