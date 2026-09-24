package main

import "fmt"

func main() {
	heights := []int{7, 1, 7, 2, 2, 4}
	larea := largestRectangleArea(heights)
	fmt.Println(larea)
}

// monotonic stack: elements are strictly maintained in a specific sorted order.

type bar struct {
	height   int
	startIdx int
}

func largestRectangleArea(heights []int) int {
	//7, 1, 7, 2, 2, 4

	bars := []bar{}
	maxArea := 0
	for i, el := range heights {
		currentBar := bar{
			height:   el,
			startIdx: i,
		}
		if len(bars) == 0 {
			bars = append(bars, currentBar)
			continue
		}

		lastBarIdx := len(bars) - 1
		lastBar := bars[lastBarIdx]
		if lastBar.height > currentBar.height {
			bars[lastBarIdx] = currentBar
			newArea := lastBar.height * (currentBar.startIdx - lastBar.startIdx)

			if newArea > maxArea {
				maxArea = newArea
			}
			currentBar.startIdx = lastBar.startIdx
			continue
		}

		if lastBar.height < currentBar.height {
			bars = append(bars, currentBar)
		}
	}

	return maxArea

	//i=0 elem=7 stack={7}
	//i=1 elem=1 stack={1} area = stack.pop() * (current_index - stack.pop().index) => 7 * (1-0) => 7
	//i=2 elem=7 stack={1, 7}
	//i=3 elem=2 stack={1, 2} area = 7 * (3-2) = 7
	//i=4 elem=2 stack={1, 2, 2}
	//i=5 elem=4 stack={1, 2, 2, 4}
}
