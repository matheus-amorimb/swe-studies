package main

// func main() {
// 	a := 'a'
// 	b := '!'
// 	c := '9'
// 	d := 'B'
// 	fmt.Printf("isAlphanumeric (%v): %v\n", string(a), isAlphaNumeric(a))
// 	fmt.Printf("isAlphanumeric (%v): %v\n", string(b), isAlphaNumeric(b))
// 	fmt.Printf("isAlphanumeric (%v): %v\n", string(c), isAlphaNumeric(c))
// 	fmt.Printf("isAlphanumeric (%v): %v\n", string(d), isAlphaNumeric(d))
// }

func isAlphaNumeric(c rune) bool {
	return isLowerCase(c) || isUpperCase(c) || isDigit(c)
}

func isLowerCase(c rune) bool {
	numRepr := int(c)
	if numRepr >= int('a') && numRepr <= int('z') {
		return true
	}
	return false
}

func isUpperCase(c rune) bool {
	numRepr := int(c)
	if numRepr >= int('A') && numRepr <= int('Z') {
		return true
	}
	return false
}

func isDigit(c rune) bool {
	numRepr := int(c)
	if numRepr >= int('0') && numRepr <= int('9') {
		return true
	}
	return false
}
