package flightguess

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
		body = append(body, Position{x: i, y: y + 1})
	}
	for i := x - 1; i <= x+1; i++ {
		body = append(body, Position{x: i, y: y + 3})
	}
	body = append(body, Position{x, y + 2})
	return Flight{head, body, combine(head, body)}
}

func makeDownFlight(head Position) Flight {
	x := head.x
	y := head.y
	var body []Position

	for i := x - 2; i <= x+2; i++ {
		body = append(body, Position{x: i, y: y - 1})
	}
	for i := x - 1; i <= x+1; i++ {
		body = append(body, Position{x: i, y: y - 3})
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
	head := Position{x: x, y: y}
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
