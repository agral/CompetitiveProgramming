package lc0417

type Point struct {
	y int
	x int
}

type PointQueue []Point

func (q *PointQueue) Enqueue(val Point) {
	*q = append(*q, val)
}

// Dequeue() without bounds checking; only run it on queues of nonzero lengths.
func (q *PointQueue) Dequeue() Point {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *PointQueue) Empty() bool {
	return len(*q) == 0
}

// Runtime complexity: O(H*W)
// Auxiliary space: O(H*W)
// Subjective level: hard.
// Solved on: 2026-01-02
func pacificAtlantic(heights [][]int) [][]int {
	DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	H := len(heights)
	W := len(heights[0])

	// we'll be going backwards, from tiles bordering with the ocean UP higher where available.
	bfs := func(q *PointQueue, seen *[][]bool) {
		for !q.Empty() {
			p := q.Dequeue()
			h := heights[p.y][p.x]
			for _, dir := range DIRS {
				y := p.y + dir[0]
				x := p.x + dir[1]
				if y >= 0 && y < H && x >= 0 && x < W && !(*seen)[y][x] && heights[y][x] >= h {
					q.Enqueue(Point{y, x})
					(*seen)[y][x] = true
				}
			}
		}
	}

	flowsToA := make([][]bool, H)
	flowsToP := make([][]bool, H)
	for h := range H {
		flowsToA[h] = make([]bool, W)
		flowsToP[h] = make([]bool, W)
	}
	qA := PointQueue{} // queue for Atlantic-adjacent tiles
	qP := PointQueue{} // queue for Pacific-adjacent tiles

	for y := range H {
		qP.Enqueue(Point{y, 0})
		flowsToP[y][0] = true
		qA.Enqueue(Point{y, W - 1})
		flowsToA[y][W-1] = true
	}
	for x := range W {
		qP.Enqueue(Point{0, x})
		flowsToP[0][x] = true
		qA.Enqueue(Point{H - 1, x})
		flowsToA[H-1][x] = true
	}

	ans := [][]int{}
	bfs(&qA, &flowsToA)
	bfs(&qP, &flowsToP)

	for y := range H {
		for x := range W {
			if flowsToA[y][x] && flowsToP[y][x] {
				ans = append(ans, []int{y, x})
			}
		}
	}

	return ans
}
