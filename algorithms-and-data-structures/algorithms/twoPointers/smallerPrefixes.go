package main

func main() {
	arr := []int{1, 2, -2, 1, 3, 5}
	b := smallerPrefix(arr)
	println(b)
}

func smallerPrefix(arr []int) bool {
	sP, fP := 0, 0
	sSum, fSum := 0, 0

	for fP < len(arr) {
		sSum += arr[sP]
		fSum += arr[fP] + arr[fP+1]

		if sSum >= fSum {
			return false
		}

		sP += 1
		fP += 2
	}

	return true
}
