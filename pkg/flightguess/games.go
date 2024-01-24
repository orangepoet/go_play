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
