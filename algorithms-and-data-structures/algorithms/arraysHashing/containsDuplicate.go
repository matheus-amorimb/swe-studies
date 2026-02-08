package main

func main() {
	nums := []int{1, 2, 3, 4, 4}
	println(hasDuplicate(nums))
}

func hasDuplicate(nums []int) bool {
	uniqueNumsMap := map[int]bool{}
	for _, num := range nums {
		_, ok := uniqueNumsMap[num]
		println(num, ok)
		if ok {
			return true
		}

		uniqueNumsMap[num] = true
	}
	return false
}
