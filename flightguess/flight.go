package flightguess

import (
	"fmt"
	"time"
)

type Position struct {
	x int
	y int
}

// Flight Model
type Flight struct {
	head      Position
	body      []Position
	positions []Position
}

type Config struct {
	mapSize         int
	flightGroupSize int
}

// enum all flight groups by groupSize
func listFlightGroup(c Config) [][]Flight {
	start := time.Now()

	var ret [][]Flight
	allFlights := listAllFlight(c.mapSize)
	listFlightGroup0(c.flightGroupSize, make([]Flight, 0), allFlights, &ret)

	println(fmt.Sprintf("listFlightGroup elapsed: %d (ms)", time.Since(start).Milliseconds()))
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

func listAllFlight(mapSize int) []Flight {
	fs := []Flight{}
	for i := 1; i <= mapSize; i++ {
		for j := 1; j <= mapSize; j++ {
			fs = append(fs, makeFlightsByHead(i, j, mapSize)...)
		}
	}
	return fs
}

func (p *Position) isPart(f Flight) bool {
	return p.isHead(f) || p.isBody(f)
}

func (p *Position) isHead(f Flight) bool {
	return *p == f.head
}

func (p *Position) isBody(f Flight) bool {
	for _, v := range f.body {
		if *p == v {
			return true
		}
	}
	return false
}

func makeUpFlight(head Position) Flight {
	x := head.x
	y := head.y
	var body []Position
	for i := x - 2; i <= x+2; i++ {
		body = append(body, Position{i, y + 1})
	}
	for i := x - 1; i <= x+1; i++ {
		body = append(body, Position{i, y + 3})
	}
	body = append(body, Position{x, y + 2})
	return Flight{head, body, combine(head, body)}
}

func makeDownFlight(head Position) Flight {
	x := head.x
	y := head.y
	var body []Position

	for i := x - 2; i <= x+2; i++ {
		body = append(body, Position{i, y - 1})

	}
	for i := x - 1; i <= x+1; i++ {
		body = append(body, Position{i, y - 3})
	}
	body = append(body, Position{x, y - 2})
	return Flight{head, body, combine(head, body)}
}

func makeLeftFlight(head Position) Flight {
	x := head.x
	y := head.y
	var body []Position
	for j := y - 2; j <= y+2; j++ {
		body = append(body, Position{x + 1, j})
	}
	for j := y - 1; j <= y+1; j++ {
		body = append(body, Position{x + 3, j})
	}
	body = append(body, Position{x + 2, y})
	return Flight{head, body, combine(head, body)}
}

func makeRightFlight(head Position) Flight {
	x := head.x
	y := head.y
	var body []Position
	for j := y - 2; j <= y+2; j++ {
		body = append(body, Position{x - 1, j})
	}
	for j := y - 1; j <= y+1; j++ {
		body = append(body, Position{x - 3, j})
	}
	body = append(body, Position{x - 2, y})
	return Flight{head, body, combine(head, body)}
}

// list all flights with head equals (x,y)
func makeFlightsByHead(x, y, n int) []Flight {
	head := Position{x, y}
	flights := []Flight{makeUpFlight(head), makeDownFlight(head), makeLeftFlight(head), makeRightFlight(head)}
	return filter(flights, func(flight Flight) bool {
		for _, p := range flight.positions {
			if p.x <= 0 || p.x > n || p.y <= 0 || p.y > n {
				return false
			}
		}
		return true
	})
}

func filter(flights []Flight, predicate func(flight Flight) bool) []Flight {
	var filtered []Flight
	for _, x := range flights {
		if predicate(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func combine(head Position, body []Position) []Position {
	positions := body[:]
	positions = append(positions, head)
	return positions
}
