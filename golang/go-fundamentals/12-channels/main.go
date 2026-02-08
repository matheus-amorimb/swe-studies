package main

func main() {
	concurrency("testing...")
	deadlock()
}

func rangeCh(n int) []int {
	fibList := []int{}
	intsCh := make(chan int)
	go fibonacci(n, intsCh)
	for value := range intsCh {
		fibList = append(fibList, value)
	}

	return fibList
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
