package main

func main() {
	s := "mam"
	b := isPalindrome(s)
	println(b)
}

func isPalindrome(s string) bool {
	leftP := 0
	rightP := len(s) - 1
	for leftP < rightP {
		if s[rightP] != s[leftP] {
			return false
		}
		leftP++
		rightP--
	}

	return true
}
