package lib

//TODO we need to return different types for X,y
//for example, raylib uses float32 for some cases (not sure why)
func NewPoint(x, y int32) Point {
	return &point{x, y}
}

type Point interface {
	GetX() int32
	GetY() int32
	Add(Point)
}

type point struct {
	x, y int32
}

func (p *point) GetX() int32 {
	return p.x
}

func (p *point) GetY() int32 {
	return p.y
}

func (p *point) Add(addPoint Point) {
	p.x += addPoint.GetX()
	p.y += addPoint.GetY()
}
