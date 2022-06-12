package square

type Square struct {
	start Point
	a     uint
	end   Point
}

type Point struct {
	x, y int
}

func (s *Square) End() Point {
	s.end = Point{x: s.start.x + int(s.a), y: s.start.y + int(s.a)}
	return s.end
}
func (s *Square) Perimeter() uint {
	return s.a * 4
}

func (s *Square) Area() uint {
	return s.a * s.a
}
