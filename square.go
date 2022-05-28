package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (Super Square) End() Point {
	// implement me
	End := Point{Super.start.x + int(Super.a), Super.start.y + int(Super.a)}
	return End
}

func (Super Square) Area() uint {
	Area := Super.a * Super.a
	return Area
}

func (Super Square) Perimeter() uint {
	// implement me
	Perimeter := Super.a * 4
	return Perimeter
}
