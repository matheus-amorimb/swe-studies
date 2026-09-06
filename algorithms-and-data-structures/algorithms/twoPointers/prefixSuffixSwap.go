package main

func main() {
	arr := []rune{
		'b', 'a', 'd', 'r', 'e', 'v', 'i', 'e', 'w',
	}
	arr = prefixSuffixSwap(arr)
	println(string(arr))
}

// r	e	v	i	e	w	b	a	d
// 							p1
// 									p2

func prefixSuffixSwap(arr []rune) []rune {
	p1, p2 := 0, len(arr)/3

	for p2 < len(arr) {
		arr[p1], arr[p2] = arr[p2], arr[p1]
		p1++
		p2++
	}

	return arr
}
