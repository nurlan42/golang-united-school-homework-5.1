package square

<<<<<<< HEAD
=======
type Square struct {
	start Point
	a     uint
	end   Point
}

>>>>>>> master
type Point struct {
	x, y int
}

<<<<<<< HEAD
type Square struct {
	start Point
	a     uint
}

func (receiver) End() Point {
	// implement me
}

func (receiver) Area() uint {
	// implement me
}

func (receiver) Perimeter() uint {
	// implement me
=======
func (s *Square) End() Point {
	s.end = Point{x: s.start.x + int(s.a), y: s.start.y + int(s.a)}
	return s.end
}
func (s *Square) Perimeter() uint {
	return s.a * 4
}

func (s *Square) Area() uint {
	return s.a * s.a
>>>>>>> master
}
