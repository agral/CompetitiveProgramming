package lc1970

import "fmt"

type Point struct {
	y int
	x int
}

type PointQueue []Point

func (q *PointQueue) Enqueue(val Point) {
	*q = append(*q, val)
}

func (q *PointQueue) Dequeue() (Point, error) {
	if len(*q) == 0 {
		return Point{0, 0}, fmt.Errorf("Dequeue() from an empty PointQueue.")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *PointQueue) Empty() bool {
	return len(*q) == 0
}

// Runtime complexity: O(log2(row*col) * row * col).
//
//	this comes from having to process (at worst case) the entire grid,
//	and repeat this until binary search for low > high terminates
//	(which happens at log2(high-low)).
//
// Auxiliary space: O(row*col). Required for queuing up to the entire grid.
// Subjective level: hard, or maybe medium++.
// Solved on: 2025-12-31
func latestDayToCross(row int, col int, cells [][]int) int {
	ans := 0
	low, high := 1, len(cells)-1

	isWalkable := func(day int) bool {
		// 1. Recreate the state of the grid on day `day`:
		grid := make([][]int, row)
		for r := range row {
			grid[r] = make([]int, col)
		}
		for d := range day {
			y, x := cells[d][0]-1, cells[d][1]-1
			grid[y][x] = 1
		}

		// 2. Put all the available starting squares into the queue:
		q := PointQueue{}
		for x := range col {
			if grid[0][x] == 0 {
				q.Enqueue(Point{0, x})
				// mark this to-be-processed tile as water, so it won't be revisited:
				grid[0][x] = 1
			}
		}

		// 3. Keep visiting the tiles and reaching for neighboring available tiles:
		DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
		for !q.Empty() {
			p, _ := q.Dequeue()
			for _, dir := range DIRS {
				y := p.y + dir[0]
				x := p.x + dir[1]
				if y >= 0 && y < row && x >= 0 && x < col && grid[y][x] == 0 {
					// If reached the bottom row, the grid is still walkable - return true immediately.
					if y == row-1 {
						return true
					}
					q.Enqueue(Point{y, x})
					grid[y][x] = 1
				}
			}
		}
		// All available tiles visited, bottom row not reached -> the grid is unwalkable.
		return false
	}

	for low <= high {
		mid := (low + high) / 2
		if isWalkable(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
