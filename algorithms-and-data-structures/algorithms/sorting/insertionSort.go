package main

func main() {
	input := []int{10, 8, 9, 7, 6, 5, 4, 3, 2, 1}
	output := insertionSort(input)
	for _, item := range output {
		print(item, " ")
	}
}

func insertionSort[T Number](input []T) []T {
	for i := 1; i < len(input); i++ {
		sortElem := input[i]
		j := i - 1

		for j >= 0 && input[j] > sortElem {
			input[j+1] = input[j]
			j--
		}

		input[j+1] = sortElem
	}

	return input
}

type Number interface {
	~int | ~float32 | ~float64
}
