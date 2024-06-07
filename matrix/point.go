package matrix

type Pt struct {
	X int
	Y int
}

type MatT interface{}

func (p Pt) Unpack() (int, int) {
	return p.Y, p.X
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
func PtsBetween(from Pt, to Pt) (between []Pt) {
	throughPoint := from
	between = append(between, throughPoint)
	for throughPoint != to {
		if throughPoint.X < to.X {
			throughPoint.X += 1
		} else if throughPoint.X > to.X {
			throughPoint.X -= 1
		}

		if throughPoint.Y < to.Y {
			throughPoint.Y += 1
		} else if throughPoint.Y > to.Y {
			throughPoint.Y -= 1
		}

		between = append(between, Pt{X: throughPoint.X, Y: throughPoint.Y})
	}

	return between
}

func FindDuplicatePts(ptList []Pt) map[Pt]int {
	ptMap := make(map[Pt]int)
	for _, pt := range ptList {
		ptMap[pt] += 1
	}

	dupeMap := make(map[Pt]int)
	for key, val := range ptMap {
		if val > 1 {
			dupeMap[key] = val
		}
	}

	return dupeMap
}
