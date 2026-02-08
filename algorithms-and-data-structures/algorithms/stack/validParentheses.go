package main

func main() {
	input := "("
	println(isValid(input))
}

func isValid(s string) bool {
	closeBracketMap := map[rune]bool{
		')': true,
		'}': true,
		']': true,
	}

	openBracketTocloseBracket := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	closeBracketsMustHave := []rune{}
	for _, char := range s {
		if isCloseBracket := closeBracketMap[char]; isCloseBracket {
			if len(closeBracketsMustHave) == 0 {
				return false
			}

			top := closeBracketsMustHave[len(closeBracketsMustHave)-1]
			if top != char {
				return false
			}

			closeBracketsMustHave = closeBracketsMustHave[:len(closeBracketsMustHave)-1]

		} else {
			correspondingCloseBracket := openBracketTocloseBracket[char]
			closeBracketsMustHave = append(closeBracketsMustHave, correspondingCloseBracket)
		}
	}

	return len(closeBracketsMustHave) == 0
}
