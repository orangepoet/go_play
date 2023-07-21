package flightguess

import (
	"fmt"
	"time"
)

// start guess
func Start() {
	config := Config{mapSize: 10, flightGroupSize: 3}

	start := time.Now()
	flightGroups := listFlightGroup(config)

	expected := []Flight{makeUpFlight(Position{3, 3}), makeDownFlight(Position{6, 6}), makeDownFlight(Position{8, 10})}

	left := len(expected)

	history := map[Position]bool{}
	times := 0
	for {
		next := guessNext(flightGroups, history)
		println(fmt.Sprintf("guess pos: (%d, %d)", next.x, next.y))
		history[next] = true
		guessResult := judgeResult(next, expected)
		println(fmt.Sprintf("guessResult: %d", guessResult))
		if guessResult == 2 {
			left--
			if left == 0 {
				println(fmt.Sprintf("attack all, WIN, times: %d", times))
				break
			}
		}
		times++
		if times > 20 {
			println("failed, guess exceed max times")
			break
		}

		flightGroups = refreshGroups(flightGroups, next, guessResult)
		println(fmt.Sprintf("left case size: %d", len(flightGroups)))
	}

	elapse := time.Since(start)
	println(fmt.Sprintf("all elaspsed: %d (ms)", elapse.Milliseconds()))
}

// branch reduction
func refreshGroups(groups [][]Flight, p Position, result int) [][]Flight {
	removes := map[int]bool{}
	for idx, fs := range groups {
		if !filterGroup(fs, p, result) {
			removes[idx] = true
		}
	}
	x := [][]Flight{}
	for idx, flights := range groups {
		if _, ok := removes[idx]; !ok {
			x = append(x, flights)
		}
	}
	return x
}

// false: need to remove, true: keep
func filterGroup(group []Flight, p Position, r int) bool {
	switch r {
	case 0:
		for _, flight := range group {
			if p.isPart(flight) {
				return false
			}
		}
		return true
	case 1:
		for _, flight := range group {
			if p.isHead(flight) {
				return false
			}
			if p.isBody(flight) {
				return true
			}
		}
		return false
	case 2:
		for _, flight := range group {
			if p.isBody(flight) {
				return false
			}
			if p.isHead(flight) {
				return true
			}
		}
		return false
	}
	return false
}

// guess next
func guessNext(flightUnits [][]Flight, history map[Position]bool) Position {
	headMap := map[Position]int{}
	var max Position
	for _, unit := range flightUnits {
		for _, flight := range unit {
			_, ok := history[flight.head]
			if ok {
				continue
			}
			headMap[flight.head] = headMap[flight.head] + 1
			if max.x == 0 {
				max = flight.head
			} else if headMap[flight.head] > headMap[max] {
				max = flight.head
			}
		}
	}
	return max
}

// judge guess result
func judgeResult(p Position, expected []Flight) int {
	r := 0
	for _, flight := range expected {
		if p.isHead(flight) {
			r = 2
			break
		} else {
			for _, position := range flight.body {
				if p == position {
					r = 1
				}
			}
		}
	}
	return r
}
