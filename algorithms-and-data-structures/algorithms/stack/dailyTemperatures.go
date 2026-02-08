package main

func main() {
	input := []int{22, 21, 20}
	output := dailyTemperatures(input)
	for _, item := range output {
		print(item, ", ")
	}
}

type temperatureStruct struct {
	tmp int
	idx int
}

func dailyTemperatures(temperatures []int) []int {
	stack := []temperatureStruct{}
	result := make([]int, len(temperatures))

	for idx, tmp := range temperatures {
		for len(stack) != 0 && tmp > stack[len(stack)-1].tmp {
			result[stack[len(stack)-1].idx] = idx - stack[len(stack)-1].idx
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, temperatureStruct{
			tmp,
			idx,
		})
	}

	return result
}
