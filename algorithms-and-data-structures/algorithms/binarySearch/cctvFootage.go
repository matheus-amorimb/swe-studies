package main

import (
	"fmt"
	"math/rand"
	"time"
)

var parkedAt time.Time = time.Date(2026, 9, 15, 7, 0, 0, 0, time.UTC)
var foundMissingAt time.Time = time.Date(2026, 9, 15, 8, 0, 0, 0, time.UTC)
var stolenAt int64 = rand.Int63n(foundMissingAt.Unix()-parkedAt.Unix()+1) + parkedAt.Unix() + 1

func main() {
	t1 := parkedAt.Add(time.Second).Unix()
	t2 := foundMissingAt.Add(-time.Second).Unix()
	delt := int64(1)

	firstMissing := int64(0)

	//invariant t1 < t2; isStolen(t1) == false, isStolen(t2) == true
	for t2-t1 > int64(delt) {
		mid := (t1 + t2) / 2
		if !isStolen(mid) {
			t1 = mid
		} else {
			t2 = mid
		}
	}

	fmt.Println("stolenAt:", stolenAt)
	fmt.Println("firstMissing:", firstMissing)
}

func isStolen(t int64) bool {
	return t >= stolenAt
}
