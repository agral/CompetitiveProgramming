package lc1926

import (
	"fmt"
	"log"
)

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

// Runtime complexity: O(WH)
// Auxiliary space: O(WH)
func nearestExit(maze [][]byte, entrance []int) int {
	DIRS := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	H := len(maze)
	W := len(maze[0])
	q := Queue{}
	q.Enqueue(entrance)
	visited := make([][]bool, H)
	for h := range H {
		visited[h] = make([]bool, W)
	}
	visited[entrance[0]][entrance[1]] = true

	for stepNum := 1; !q.Empty(); stepNum++ {
		for sz := q.Len(); sz > 0; sz-- {
			place, err := q.Dequeue()
			if err != nil {
				log.Fatal("Queue emptied mid-size. Should never happen!")
			}
			py, px := place[0], place[1]
			for _, d := range DIRS {
				y := py + d[0]
				x := px + d[1]
				if y < 0 || y >= H || x < 0 || x >= W {
					continue
				}
				if visited[y][x] || maze[y][x] == '+' {
					continue
				}
				if y == 0 || y == H-1 || x == 0 || x == W-1 {
					return stepNum
				}
				q.Enqueue([]int{y, x})
				visited[y][x] = true
			}
		}
	}
	return -1
}
