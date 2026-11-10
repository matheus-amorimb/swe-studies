package main

import (
	"fmt"
)

func main() {
	actions := []action{
		{act: _go, url: "matheus.com"},
		{act: _go, url: "google.com"},
		{act: _go, url: "wikipedia.com"},
		{act: _go, url: "amazon.com"},
		{act: back, times: 4},
		{act: _go, url: "youtube.com"},
		{act: _go, url: "netflix.com"},
		{act: back, times: 5},
	}

	fmt.Println(currentCurl(actions))
}

type act string

const (
	_go  act = "go"
	back act = "back"
)

type action struct {
	act   act
	url   string
	times int
}

func currentCurl(actions []action) string {
	lifo := []string{}

	for _, action := range actions {
		if action.act == _go {
			lifo = append(lifo, action.url)
		}

		if action.act == back {
			//REUSABLE IDEIA
			//USE MAX INSTEAD OF (IF REMAINING > THRESHOLD, REMAINING = THRESHOLD)
			remaining := max(len(lifo)-action.times, 1)
			if remaining <= len(lifo) {
				lifo = lifo[:remaining]
			}
		}
	}

	return lifo[len(lifo)-1]
}
