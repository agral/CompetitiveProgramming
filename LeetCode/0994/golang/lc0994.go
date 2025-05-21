package lc0994

import "fmt"

type Queue [][]int

func (q *Queue) Enqueue(val []int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() ([]int, error) {
	if len(*q) == 0 {
		return []int{0, 0}, fmt.Errorf("Dequeue() from an empty Queue.")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *Queue) Len() int {
	return len(*q)
}
func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// Runtime complexity: O(HW)
// Auxiliary space: O(HW)
// where H, W are height and width of the `grid`.
func orangesRotting(grid [][]int) int {
	DIRS := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	H := len(grid)
	W := len(grid[0])
	q := Queue{}
	numFreshOranges := 0

	for y := range H {
		for x := range W {
			if grid[y][x] == 1 {
				numFreshOranges += 1
			} else if grid[y][x] == 2 {
				q.Enqueue([]int{y, x})
			}
		}
	}
	if numFreshOranges == 0 {
		return 0
	}

	stepNum := 0
	for !q.Empty() {
		stepNum += 1
		for range q.Len() {
			startPos, _ := q.Dequeue()
			sy, sx := startPos[0], startPos[1]
			for _, dir := range DIRS {
				y := sy + dir[0]
				x := sx + dir[1]
				if y < 0 || y >= H || x < 0 || x >= W {
					continue
				}
				if grid[y][x] == 1 {
					grid[y][x] = 2
					q.Enqueue([]int{y, x})
					numFreshOranges--
				}
			}
		}
	}
	if numFreshOranges == 0 {
		return stepNum - 1
	}

	return -1
}
