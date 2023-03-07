package checkneighbor

type Point struct {
	X int
	Y int
}

func IsNeighbor(p1, p2 Point) bool {

	neighbors := []Point{
		{p1.X - 1, p1.Y - 1}, {p1.X, p1.Y - 1}, {p1.X + 1, p1.Y - 1},
		{p1.X - 1, p1.Y}, {p1.X + 1, p1.Y},
		{p1.X - 1, p1.Y + 1}, {p1.X, p1.Y}, {p1.X + 1, p1.Y + 1},
	}
	for _, n := range neighbors {
		if p2 == n {
			return true
		}
	}
	return false

}
