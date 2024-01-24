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
