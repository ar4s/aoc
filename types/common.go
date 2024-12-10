package types

type Cord2D struct {
	X int
	Y int
}

type Cords2D []Cord2D

func (c Cords2D) Len() int {
	return len(c)
}
func (c Cords2D) Less(i, j int) bool {
	return c[i].X > c[j].X && c[i].Y > c[j].X
}
func (c Cords2D) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Direction2D Cord2D

var DIR_UP Direction2D = Direction2D{X: 0, Y: -1}
var DIR_DOWN Direction2D = Direction2D{X: 0, Y: 1}
var DIR_LEFT Direction2D = Direction2D{X: -1, Y: 0}
var DIR_RIGHT Direction2D = Direction2D{X: 1, Y: 0}

var DIR_4 = []Direction2D{
	DIR_UP, DIR_DOWN, DIR_LEFT, DIR_RIGHT,
}

func (c Cord2D) Add(a Cord2D) Cord2D {
	return Cord2D{
		X: c.X + a.X,
		Y: c.Y + a.Y,
	}
}
func (c Cord2D) IsOutOfBound(maxSize int) bool {
	return c.X < 0 || c.Y < 0 || c.X >= maxSize || c.Y >= maxSize
}

func (c Cord2D) ApplyDirection(a Direction2D) Cord2D {
	return Cord2D{
		X: c.X + a.X,
		Y: c.Y + a.Y,
	}
}

func (c Cord2D) Less(x, y int) bool {
	return false
}

func (c Direction2D) Negative() Direction2D {
	return Direction2D{
		X: c.X * -1,
		Y: c.Y * -1,
	}
}

func (c Direction2D) RotateCW() Direction2D {
	switch c {
	case DIR_UP:
		return DIR_RIGHT
	case DIR_RIGHT:
		return DIR_DOWN
	case DIR_DOWN:
		return DIR_LEFT
	case DIR_LEFT:
		return DIR_UP
	}
	return c
}
