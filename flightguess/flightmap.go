package flightguess

import (
	"fmt"
	"time"
)

type FlightMap struct {
	n  int
	ps []Position
}

// enum all flight groups by groupSize
func (fm *FlightMap) listFlightGroup(groupSize int) [][]Flight {
	start := time.Now()

	var ret [][]Flight
	allFlights := fm.listAllFlight()
	listFlightGroup0(groupSize, make([]Flight, 0), allFlights, &ret)

	end := time.Since(start)
	println(fmt.Sprintf("listFlightGroup elapsed: %d", end))
	return ret
}

// recursive get
func listFlightGroup0(groupSize int, flightGroup []Flight, allFlights []Flight, result *[][]Flight) {
	if len(flightGroup) == groupSize {
		*result = append(*result, flightGroup)
		return
	}
	pMap := make(map[Position]bool)
	for _, f := range flightGroup {
		for _, p := range f.positions {
			pMap[p] = true
		}
	}

	for _, flight := range allFlights {
		if !isOverLap(flight.positions, pMap) {
			fs2 := append(flightGroup, flight)
			listFlightGroup0(groupSize, fs2, allFlights, result)
		}
	}
}

// test if ps's positions is overlap for pMap's positions ps: test position array, pMap: exists position map;
func isOverLap(ps []Position, pMap map[Position]bool) bool {
	for _, p := range ps {
		if _, ok := pMap[p]; ok {
			return true
		}
	}
	return false
}

func makeFlightMap(n int) FlightMap {
	var ps []Position
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ps = append(ps, Position{i, j})
		}
	}
	return FlightMap{n, ps}
}

func (fm *FlightMap) listAllFlight() []Flight {
	fs := []Flight{}
	for i := 1; i <= fm.n; i++ {
		for j := 1; j <= fm.n; j++ {
			fs = append(fs, makeFlightsByHead(i, j, fm.n)...)
		}
	}
	return fs
}
