package flightguess

import (
	"fmt"
	"time"
)

// FlightGroup 飞机组
type FlightGroup []*Flight

func listFlightGroup() []FlightGroup {
	start := time.Now()

	var allFlightGroups []FlightGroup
	allFlights := listAllFlight()
	listFlightGroup0(FlightGroup{}, allFlights, &allFlightGroups)

	fmt.Printf("listFlightGroup elapsed: %d (ms)\n", time.Since(start).Milliseconds())
	return allFlightGroups
}

// recursive get
func listFlightGroup0(flightGroup FlightGroup, allFlights FlightGroup, result *[]FlightGroup) {
	if len(flightGroup) == GroupSize {
		*result = append(*result, flightGroup)
		return
	}
	pMap := make(map[Position]bool)
	for _, f := range flightGroup {
		for _, p := range f.positions() {
			pMap[p] = true
		}
	}

	for _, flight := range allFlights {
		if !isOverLap(flight.positions(), pMap) {
			fs2 := append(flightGroup, flight)
			listFlightGroup0(fs2, allFlights, result)
		}
	}
}

// test if ps's positions is overlap for pMap's positions ps: test positions array, pMap: exists positions map;
func isOverLap(ps []Position, pMap map[Position]bool) bool {
	if pMap == nil || len(pMap) == 0 || len(ps) == 0 {
		return false
	}
	for _, p := range ps {
		if _, ok := pMap[p]; ok {
			return true
		}
	}
	return false
}

func listAllFlight() FlightGroup {
	fs := FlightGroup{}
	for i := 1; i <= MapSize; i++ {
		for j := 1; j <= MapSize; j++ {
			fs = append(fs, listFlightsByHead(i, j)...)
		}
	}
	return fs
}

// list all flights with head equals (x,y)
func listFlightsByHead(x, y int) FlightGroup {
	head := Position{x, y}
	flights := FlightGroup{
		makeFlight(head, DOWN),
		makeFlight(head, UP),
		makeFlight(head, LEFT),
		makeFlight(head, RIGHT),
	}
	return flights.filter(func(f *Flight) bool {
		for _, p := range f.positions() {
			if p.x <= 0 || p.x > MapSize || p.y <= 0 || p.y > MapSize {
				return false
			}
		}
		return true
	})
}

func Guess() {
	start := time.Now()
	flightGroups := listFlightGroup()

	expected := []*Flight{
		makeFlight(Position{3, 3}, UP),
		makeFlight(Position{6, 6}, DOWN),
		makeFlight(Position{8, 10}, DOWN),
	}

	left := len(expected)

	history := make(map[Position]interface{})
	times := 0
	for {
		next := guessNext(flightGroups, history)
		fmt.Printf("guess pos: (%d, %d)", next.x, next.y)
		history[next] = 0
		guessResult := judgeResult(next, expected)
		fmt.Printf("guessResult: %d\n", guessResult)
		if guessResult == 2 {
			left--
			if left == 0 {
				fmt.Printf("attack all, WIN, times: %d", times)
				break
			}
		}
		times++
		if times > 20 {
			fmt.Println("failed, guess exceed max times")
			break
		}

		flightGroups = refreshGroups(flightGroups, next, guessResult)
		fmt.Printf("left case size: %d", len(flightGroups))
	}

	elapse := time.Since(start)
	fmt.Printf("all elaspsed: %d (ms)\n", elapse.Milliseconds())
}

// branch reduction
func refreshGroups(groups []FlightGroup, p Position, result int) []FlightGroup {
	removes := make(map[int]interface{})
	for idx, g := range groups {
		if !filterGroup(g, p, result) {
			removes[idx] = 0
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
func guessNext(flightUnits []FlightGroup, history map[Position]interface{}) Position {
	headMap := map[Position]int{}
	var max0 *Position
	for _, unit := range flightUnits {
		for _, flight := range unit {
			_, ok := history[flight.head]
			if ok {
				continue
			}
			headMap[flight.head] = headMap[flight.head] + 1
			if max0 == nil {
				max0 = &flight.head
			} else if headMap[flight.head] > headMap[*max0] {
				max0 = &flight.head
			}
		}
	}
	return *max0
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
