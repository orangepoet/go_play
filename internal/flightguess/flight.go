package flightguess

// Position 位置
type Position struct {
	x int
	y int
}

// Body 机身
type Body []Position

// Flight Model
type Flight struct {
	head Position
	body Body
}

func (f *Flight) positions() []Position {
	ps := f.body
	ps = append(ps, f.head)
	return ps
}

// enum all flight groups by GroupSize

func (p *Position) isPartOf(f *Flight) bool {
	return p.isHeadOf(f) || p.isBodyOf(f)
}

func (p *Position) isHeadOf(f *Flight) bool {
	return *p == f.head
}

func (p *Position) isBodyOf(f *Flight) bool {
	for _, v := range f.body {
		if *p == v {
			return true
		}
	}
	return false
}

func (fg FlightGroup) filter(predicate func(flight *Flight) bool) FlightGroup {
	var filtered []*Flight
	for _, x := range fg {
		if predicate(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func downBody(head Position) *Body {
	x := head.x
	y := head.y
	var body Body

	for i := x - 2; i <= x+2; i++ {
		body = append(body, Position{i, y - 1})

	}
	for i := x - 1; i <= x+1; i++ {
		body = append(body, Position{i, y - 3})
	}
	body = append(body, Position{x, y - 2})
	return &body
}

func upBody(head Position) *Body {
	x := head.x
	y := head.y
	var body Body
	for i := x - 2; i <= x+2; i++ {
		body = append(body, Position{i, y + 1})
	}
	for i := x - 1; i <= x+1; i++ {
		body = append(body, Position{i, y + 3})
	}
	body = append(body, Position{x, y + 2})
	return &body
}

func leftBody(head Position) *Body {
	x := head.x
	y := head.y
	var body Body
	for j := y - 2; j <= y+2; j++ {
		body = append(body, Position{x + 1, j})
	}
	for j := y - 1; j <= y+1; j++ {
		body = append(body, Position{x + 3, j})
	}
	body = append(body, Position{x + 2, y})
	return &body
}

func rightBody(head Position) *Body {
	x := head.x
	y := head.y
	var body Body
	for j := y - 2; j <= y+2; j++ {
		body = append(body, Position{x - 1, j})
	}
	for j := y - 1; j <= y+1; j++ {
		body = append(body, Position{x - 3, j})
	}
	body = append(body, Position{x - 2, y})
	return &body
}

func makeFlight(p Position, d Direction) *Flight {
	var body *Body
	switch d {
	case DOWN:
		body = downBody(p)
	case UP:
		body = upBody(p)
	case LEFT:
		body = leftBody(p)
	case RIGHT:
		body = rightBody(p)
	default:
		panic("unhandled default case")
	}
	return &Flight{head: p, body: *body}
}
