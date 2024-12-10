package types

type Cord2D struct {
	X int
	Y int
}

var DIR_UP Cord2D = Cord2D{X: 0, Y: -1}
var DIR_DOWN Cord2D = Cord2D{X: 0, Y: 1}
var DIR_LEFT Cord2D = Cord2D{X: -1, Y: 0}
var DIR_RIGHT Cord2D = Cord2D{X: 1, Y: 0}

var DIR_4 = []Cord2D{
	DIR_UP, DIR_DOWN, DIR_LEFT, DIR_RIGHT,
}

func (c Cord2D) Add(a Cord2D) Cord2D {
	return Cord2D{
		X: c.X + a.X,
		Y: c.Y + a.Y,
	}
}

func (c Cord2D) Negative() Cord2D {
	return Cord2D{
		X: c.X * -1,
		Y: c.Y * -1,
	}
}
