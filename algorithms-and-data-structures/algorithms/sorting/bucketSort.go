package main

func main() {
	input := []float32{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68, 0.1, 1.1, 80}
	output := bucketSort(input)
	for _, item := range output {
		print(item, " ")
	}
}

func bucketSort[T Number](input []T) []T {
	output := []T{}
	inputLen := len(input)
	buckets := make([][]T, inputLen)
	for _, item := range input {
		idx := int(float64(item) * float64(inputLen))
		if idx >= inputLen {
			idx = inputLen - 1
		}
		buckets[idx] = append(buckets[idx], item)
	}

	for idx, bucket := range buckets {
		buckets[idx] = insertionSort(bucket)
	}

	for _, bucket := range buckets {
		output = append(output, bucket...)
	}

	return output
}

type Number interface {
	~int | ~float32 | ~float64
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
