package main

import "fmt"

func main() {
	arr := []string{"join", "by", "space"}
	c := ' '
	str := join(arr, c)
	fmt.Printf(`output: "%v"`, str)
	println()
}

func join(arr []string, c rune) string {
	otp := make([]rune, 0)
	for i := range len(arr) {
		if i != 0 {
			otp = append(otp, c)
		}
		for _, c := range arr[i] {
			otp = append(otp, c)
		}
	}

	return arrayToString(otp)
}

func arrayToString(s []rune) string {
	buf := make([]byte, 0, len(s))
	for _, char := range s {
		buf = append(buf, []byte(string(char))...)
	}
	return string(buf)
}
