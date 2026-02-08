package main

func main() {
	input := []int{-1, 0, 1, 2, 3}
	// output := productExceptSelf(input)
	output := productExceptSelfOptimalSolution(input)
	for _, value := range output {
		print(value, ", ")
	}
}

func productExceptSelf(nums []int) []int {
	//O(n)
	totalProduct := 1
	zeroCount := 0
	for _, num := range nums {
		if num != 0 {
			totalProduct *= num
		} else {
			zeroCount++
		}
	}

	output := make([]int, len(nums))
	if zeroCount > 1 {
		return output
	}

	for idx, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				output[idx] = totalProduct
			} else {
				output[idx] = 0
			}
		} else {
			output[idx] = totalProduct / num
		}
	}

	return output
}

// prefix/sufix technique
func productExceptSelfOptimalSolution(nums []int) []int {
	lenNums := len(nums)

	prefix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < lenNums; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	sufix := make([]int, lenNums)
	sufix[lenNums-1] = 1
	for i := lenNums - 2; i >= 0; i-- {
		sufix[i] = sufix[i+1] * nums[i+1]
	}

	output := make([]int, lenNums)
	for i := 0; i < lenNums; i++ {
		output[i] = prefix[i] * sufix[i]
	}

	return output
}
