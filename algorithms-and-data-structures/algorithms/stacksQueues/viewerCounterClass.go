package main

import "fmt"

func main() {
	counter := NewViewerCounter(10)
	counter.join(1, "subscriber")
	counter.join(1, "guest")
	counter.join(2, "follower")
	counter.join(2, "follower")
	counter.join(2, "follower")
	counter.join(3, "follower")
	fmt.Println("subscriber [0,10]:", counter.getViewers(10, "subscriber"))
	fmt.Println("guest [0,10]:", counter.getViewers(10, "subscriber"))
	fmt.Println("follower [0,10]:", counter.getViewers(10, "follower"))
	fmt.Println("follower [3,13]:", counter.getViewers(13, "follower"))
}

type ViewerCounter struct {
	window  int
	viewers map[int]map[string]int
}

func NewViewerCounter(w int) *ViewerCounter {
	return &ViewerCounter{
		window:  w,
		viewers: map[int]map[string]int{},
	}
}

func (vc *ViewerCounter) join(t int, v string) {
	if _, ok := vc.viewers[t]; !ok {
		vc.viewers[t] = map[string]int{}
	}
	vc.viewers[t][v]++
}

func (vc *ViewerCounter) getViewers(t int, v string) int {
	init := t - vc.window
	end := t
	counter := 0
	for i := init; i <= end; i++ {
		counter += vc.viewers[i][v]
	}
	return counter
}

// I need store the data partitioned to make it easier to find them when getViewers is called.
// The window isn't fixed. So I can't partition by window.
// If I partition by time, I can iterate over window and sum the values up.
