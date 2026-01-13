package lc3453

import "sort"

// Runtime complexity: O(sort) == O(n*logn)
// Auxiliary space: O(2*3n) == O(n)
// Subjective level: medium+
// Solved on: 2026-01-13
func separateSquares(squares [][]int) float64 {
	halfArea := 0.0
	for _, square := range squares {
		halfArea += float64(0.5) * float64(square[2]) * float64(square[2])
	}

	type Geometry struct {
		yPos    int
		isStart bool // either true (start of square) or false (end of square).
		length  int
	}

	geometry := make([]Geometry, 0, 2*len(squares))

	for _, square := range squares {
		yPos := square[1]
		length := square[2]
		geometry = append(geometry, Geometry{yPos, true, length})
		geometry = append(geometry, Geometry{yPos + length, false, length})
	}

	sort.Slice(geometry, func(i, j int) bool {
		return geometry[i].yPos < geometry[j].yPos
	})

	area := float64(0.0)
	width := 0
	lastY := 0
	for _, geom := range geometry {
		deltaArea := float64(width * (geom.yPos - lastY))
		if area+deltaArea >= halfArea {
			return float64(lastY) + float64(halfArea-area)/float64(width)
		}

		// otherwise keep closing in on the actual half area:
		area += deltaArea
		if geom.isStart {
			width += geom.length
		} else {
			width -= geom.length
		}
		lastY = geom.yPos
	}

	// Should never be reached, but return a nonsense answer instead of throwing or doing other stuff.
	return -42.0
}
