package main

import "fmt"

func main() {
	arr := []rune{
		's', 'e', 'e', 'k', 'e', 'r', 'a', 'n', 'd', 'w', 'r', 'i', 't', 'e', 'r',
	}
	word := "edit"
	arr = shiftWordToBack(arr, word)
	// arr = shiftItemToEndArr(0, arr)
	fmt.Println(string(arr))
}

func shiftWordToBack(arr []rune, word string) []rune {
	s, w, i := 0, 0, 0

	for s < len(arr) {
		if i < len(word) && arr[s] == rune(word[i]) {
			i++
		} else {
			arr[w] = arr[s]
			w++
		}
		s++
	}

	for _, char := range word {
		arr[w] = char
		w++
	}

	return arr
}

func shiftItemToEndArr(idx int, arr []rune) []rune {
	itemToShift := arr[idx]
	for i := idx; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[len(arr)-1] = itemToShift

	return arr
}
