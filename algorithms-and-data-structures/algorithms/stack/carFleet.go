// A car fleet is a non-empty set of cars driving at the same position and same speed. A single car is also considered a car fleet.
package main

import (
	"slices"
)

func main() {
	target := 10
	position := []int{4, 1, 0, 7}
	speed := []int{2, 2, 1, 1}
}

type car struct {
	position int
	speed    int
	arriveAt float64
}

type fleet struct {
	arriveAt float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := []car{}
	nCars := len(position)
	for i := 0; i < nCars; i++ {
		position := position[i]
		speed := speed[i]
		arriveAt := float64(target-position) / float64(speed)
		cars = append(cars, car{
			arriveAt: arriveAt,
			position: position,
			speed:    speed,
		})
	}

	fleets := []fleet{}
	//O(n log n)
	slices.SortFunc(cars, func(a, b car) int {
		return b.position - a.position
	})

	for i := 0; i < nCars; i++ {
		arriveAt := cars[i].arriveAt
		if len(fleets) == 0 {
			fleets = append(fleets, fleet{
				arriveAt: arriveAt,
			})
		} else {
			lastFleet := fleets[len(fleets)-1]
			if lastFleet.arriveAt < arriveAt {
				fleets = append(fleets, fleet{
					arriveAt: arriveAt,
				})
			}
		}
	}

	return len(fleets)
}
