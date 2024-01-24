package flightguess

// Direction 方向
type Direction int

const (
	UNKNOWN Direction = iota
	// DOWN 向下
	DOWN
	// UP 向上
	UP
	// LEFT 向左
	LEFT
	// RIGHT 向右
	RIGHT
)

const (
	// MapSize 地图size
	MapSize = 10
	// GroupSize 飞机组大小
	GroupSize = 3
)
