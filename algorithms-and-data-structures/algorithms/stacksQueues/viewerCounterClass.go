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
	counter.join(5, "follower")
	counter.join(6, "follower")
	counter.join(10, "follower")
	fmt.Println("subscriber [0,10]:", counter.getViewers(10, "subscriber"))
	fmt.Println("guest [0,10]:", counter.getViewers(10, "subscriber"))
	fmt.Println("follower [0,10]:", counter.getViewers(10, "follower"))
	fmt.Println("follower [3,13]:", counter.getViewers(13, "follower"))
	fmt.Println("follower [4,14]:", counter.getViewers(14, "follower"))
}

type ViewerCounter struct {
	window  int
	viewers map[string][]int
}

func NewViewerCounter(w int) *ViewerCounter {
	return &ViewerCounter{
		window:  w,
		viewers: map[string][]int{},
	}
}

func (vc *ViewerCounter) join(t int, v string) {
	if _, ok := vc.viewers[v]; !ok {
		vc.viewers[v] = []int{}
	}
	vc.removeOldViewers(t, v)
	vc.viewers[v] = append(vc.viewers[v], t)
}

func (vc *ViewerCounter) getViewers(t int, v string) int {
	vc.removeOldViewers(t, v)
	return len(vc.viewers[v])
}

func (vc *ViewerCounter) removeOldViewers(t int, v string) {
	viewers := vc.viewers[v]
	for len(viewers) > 0 && viewers[0] < t-vc.window {
		viewers = viewers[1:]
	}
	vc.viewers[v] = viewers
}
