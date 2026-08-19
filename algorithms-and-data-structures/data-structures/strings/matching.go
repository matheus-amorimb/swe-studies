package main

import (
	"math"
)

func main() {
	s := "matheus batistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistabatistaeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	t := "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
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
		if hash == subStrHash && s[idx:idx+len(t)] == t {
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
	cPower := map[int]int{}
	cPower[0] = 1 % prm.Mod
	for idx := 1; idx < prm.WindowSize; idx++ {
		cPower[idx] = (prm.Base * cPower[idx-1]) % prm.Mod
	}

	currentHash := 0
	for idx := 0; idx < prm.WindowSize; idx++ {
		currentHash = (currentHash + (int(prm.S[idx]) * cPower[prm.WindowSize-(idx+1)])) % prm.Mod
	}
	rollHash = append(rollHash, currentHash)

	for idx := 1; idx <= (len(prm.S) - prm.WindowSize); idx++ {
		// remove the contribution of the first char in the window
		currentHash = (currentHash - int(prm.S[idx-1])*cPower[prm.WindowSize-1]) % prm.Mod

		// shift the windod by on char and add the new char to the hash
		currentHash = (currentHash*prm.Base + int(prm.S[idx+prm.WindowSize-1])) % prm.Mod

		// prevent negative result from mod
		currentHash = ((currentHash % prm.Mod) + prm.Mod) % prm.Mod
		rollHash = append(rollHash, currentHash)
	}

	return rollHash
}
