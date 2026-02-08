package main

import (
	"strconv"
)

func main() {
	strs := []string{"5#abc", "   ", "!@3#$%^&*()_+", "LongStringWithNoSpaces", "Another, String With, Commas"}
	encodedString := Encode(strs)
	output := Decode(encodedString)
	for _, str := range output {
		print(str + ", ")
	}
}

type Solution struct{}

func Encode(strs []string) string {
	encodedStr := ""
	//O(m)
	for _, str := range strs {
		encodedStr += strconv.Itoa(len(str)) + "#" + str
	}
	return encodedStr
}

func Decode(str string) []string {
	decodedStrs := []string{}
	i := 0
	for i < len(str) {
		j := i
		for string(str[j]) != "#" {
			j++
		}
		lenStr, _ := strconv.Atoi(str[i:j])
		//j is the index of the char #, therefore it must be added only by one
		j += 1
		decodedStr := str[j : j+lenStr]
		decodedStrs = append(decodedStrs, decodedStr)
		i = j + lenStr
	}

	return decodedStrs
}
