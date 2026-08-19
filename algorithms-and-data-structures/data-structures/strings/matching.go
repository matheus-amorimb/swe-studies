package main

import (
	"math"
)

func main() {
	s := "matheus batista"
	t := "eus"
	idx := indexOf(s, t)
	println("index: ", idx)
}

func indexOf(s, t string) int {
	base := 31
	mod := int(math.Pow(10, 7) + 7)

	prm := RollingHashParameters{
		S:          s,
		WindowSize: len(t),
		Base:       base,
		Mod:        mod,
	}
	rollHashs := rollingHash(prm)
	subStrHash := rollingHash(RollingHashParameters{
		S:          t,
		WindowSize: len(t),
		Base:       base,
		Mod:        mod,
	})[0]

	for idx, hash := range rollHashs {
		if hash == subStrHash {
			return idx
		}
	}

	return -1
}

type RollingHashParameters struct {
	S          string
	WindowSize int
	Base       int
	Mod        int
}

func rollingHash(prm RollingHashParameters) []int {
	//hash(substring) = (c1 * a^(k-1) + c2 * a^(k-2) + ... + ck * a^0) mod m
	// where c1, c2, ..., ck are the ASCII values of the characters in the substring,
	// a is a prime number (e.g., 31),
	// k is the length of the substring,
	// and m is a large prime number (e.g., 10^9 + 7).

	rollHash := make([]int, 0)

	for idx := 0; idx <= (len(prm.S) - prm.WindowSize); idx++ {
		subStr := prm.S[idx : idx+prm.WindowSize]

		var subStrAcc float64
		for idx := 0; idx < len(subStr); idx++ {
			subStrAcc += float64(subStr[idx]) * math.Pow(float64(prm.Base), float64(prm.WindowSize-(idx+1)))
		}

		subStrHash := int64(subStrAcc) % int64(prm.Mod)
		rollHash = append(rollHash, int(subStrHash))
	}

	return rollHash
}
