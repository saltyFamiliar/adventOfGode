package ergo

type MatT interface{}

type Pt struct {
	y int
	x int
}

func (p Pt) Unpack() (int, int) {
	return p.y, p.x
}

func (p Pt) IsInBounds(matrix [][]MatT) bool {
	if len(matrix) == 0 {
		return false
	}

	y, x := p.Unpack()
	return y >= 0 && y < len(matrix) && x >= 0 && x < len(matrix[0])
}

func (p Pt) SearchAroundPoint(matrix [][]MatT, searchDist int, searchFunc func(t MatT) bool) (points []Pt) {
	y, x := p.Unpack()
	for i := y - searchDist; i <= y+searchDist; i++ {
		for j := x - searchDist; j <= x+searchDist; j++ {
			if i == y && j == x {
				continue
			}

			searchPt := Pt{i, j}
			if searchPt.IsInBounds(matrix) && searchFunc(matrix[i][j]) {
				points = append(points, searchPt)
			}
		}
	}
	return points
}
