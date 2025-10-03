package lc0407

import "container/heap"

type Point struct {
	y int
	x int
	h int
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].h < h[j].h }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(p any) {
	*h = append(*h, p.(Point))
}

func (h *MinHeap) Pop() any {
	instance := *h
	n := len(instance)
	ans := instance[n-1]
	*h = instance[0 : n-1]
	return ans
}

// Runtime complexity: O(W * H * log(W * H))
// Auxiliary space: O(W*H)
// Subjective level: hard.
func trapRainWater(heightMap [][]int) int {
	DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	H := len(heightMap)
	W := len(heightMap[0])
	minHeap := MinHeap{}
	seen := make([][]bool, H)
	for h := range H {
		seen[h] = make([]bool, W)
	}
	ans := 0

	// Push points from the border of the map to the heap, to potentially be used as walls:
	for h := range H {
		heap.Push(&minHeap, Point{h, 0, heightMap[h][0]})
		seen[h][0] = true
		heap.Push(&minHeap, Point{h, W - 1, heightMap[h][W-1]})
		seen[h][W-1] = true
	}
	for w := 1; w < W-1; w++ {
		heap.Push(&minHeap, Point{0, w, heightMap[0][w]})
		seen[0][w] = true
		heap.Push(&minHeap, Point{H - 1, w, heightMap[H-1][w]})
		seen[H-1][w] = true
	}

	for minHeap.Len() > 0 {
		p := heap.Pop(&minHeap).(Point)
		for _, dir := range DIRS {
			y := p.y + dir[0]
			x := p.x + dir[1]
			if y < 0 || y >= H || x < 0 || x >= W || seen[y][x] {
				continue
			}
			if heightMap[y][x] < p.h {
				// this cell contains trapped water.
				ans += p.h - heightMap[y][x]
				heap.Push(&minHeap, Point{y, x, p.h})
			} else {
				// this cell does not contain trapped water,
				// but it may act as a wall for other cells.
				heap.Push(&minHeap, Point{y, x, heightMap[y][x]})
			}
			seen[y][x] = true
		}
	}

	return ans
}
