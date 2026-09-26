package main

import (
	"fmt"
	"slices"
)

func main() {
	pageCounts := []int{20, 15, 17, 10}
	days := 14
	pgs := minPagesPerDay(pageCounts, days)
	fmt.Println("pages:", pgs)
}

// two different chapters can't be studied on the same day
// len(pageCounts) <= days
func minPagesPerDay(arr []int, days int) int {
	min, max := 1, slices.Max(arr)

	l, r := min, max
	for r-l > 1 {
		mid := (l + r) / 2
		fmt.Println("l:", l, "|", "r:", r, "|", "mid:", mid, "|", "daysToFinish:", daysToFinishReading(arr, mid), "|", "isCompletedBeforeDaysTarget:", isCompletedAfterDaysTarget(arr, mid, days))
		if isCompletedAfterDaysTarget(arr, mid, days) {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}

func isCompletedAfterDaysTarget(pageCounts []int, pagesPerDay int, days int) bool {
	daysToFinish := daysToFinishReading(pageCounts, pagesPerDay)
	// daysToFinish < days, we're reading too much, we should read less
	// daysToFinish > days, we're reading too little, we should read more
	return daysToFinish > days
}

func daysToFinishReading(pageCounts []int, pagesPerDay int) int {
	days := 0
	for _, chapterPages := range pageCounts {
		days += chapterPages / pagesPerDay
		if chapterPages%pagesPerDay != 0 {
			days++
		}
	}
	// fmt.Println("pagesPerDay:", pagesPerDay, "|", "daysToFinishReading:", days)
	return days
}
