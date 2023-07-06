package lib

func NewPoint(x, y int32) Point {
	return &PointData{x, y}
}

type Point interface {
	GetX() int32
	GetY() int32
	Add(Point)
}

type PointData struct {
	X int32 `yaml:"X"`
	Y int32 `yaml:"Y"`
}

func (p *PointData) GetX() int32 {
	return p.X
}

func (p *PointData) GetY() int32 {
	return p.Y
}

func (p *PointData) Add(addPoint Point) {
	p.X += addPoint.GetX()
	p.Y += addPoint.GetY()
}
