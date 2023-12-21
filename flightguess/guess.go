package flightguess

import (
	"fmt"
	"time"
)

// start guess
func Start() {

	start := time.Now()
	flightGroups := listFlightGroup()

	expected := []*Flight{
		makeFlight(Position{3, 3}, UP), makeFlight(Position{6, 6}, DOWN), makeFlight(Position{8, 10}, DOWN),
	}

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
func refreshGroups(groups []FlightGroup, p Position, result int) []FlightGroup {
	removes := map[int]bool{}
	for idx, g := range groups {
		if !filterGroup(g, p, result) {
			removes[idx] = true
		}
	}
	gs := make([]FlightGroup, 0)
	for idx, flights := range groups {
		if _, ok := removes[idx]; !ok {
			gs = append(gs, flights)
		}
	}
	return gs
}

// false: need to remove, true: keep
func filterGroup(group FlightGroup, p Position, r int) bool {
	switch r {
	case 0: // none
		for _, flight := range group {
			if p.isPartOf(flight) {
				return false
			}
		}
		return true
	case 1: // body
		for _, flight := range group {
			if p.isHeadOf(flight) {
				return false
			}
			if p.isBodyOf(flight) {
				return true
			}
		}
		return false
	case 2: // head
		for _, flight := range group {
			if p.isBodyOf(flight) {
				return false
			}
			if p.isHeadOf(flight) {
				return true
			}
		}
		return false
	}
	return false
}

// guess next
func guessNext(flightUnits []FlightGroup, history map[Position]bool) Position {
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
func judgeResult(p Position, expected FlightGroup) int {
	r := 0
	for _, flight := range expected {
		if p.isHeadOf(flight) {
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
