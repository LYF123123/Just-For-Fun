package checkneighbor

import "testing"

func TestIsNeighbor(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{-1, -1}
	p3 := Point{1, 1}
	p4 := Point{2, 2}
	res := IsNeighbor(p1, p2)
	t.Log(res)
	res = IsNeighbor(p1, p3)
	t.Log(res)
	res = IsNeighbor(p1, p4)
	t.Log(res)
}
