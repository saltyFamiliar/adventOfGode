package toolbelt

type Pt struct {
	X int
	Y int
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
