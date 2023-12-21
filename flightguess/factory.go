package flightguess

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
