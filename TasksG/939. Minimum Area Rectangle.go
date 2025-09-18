package main

import (
	"math"
)

type p [2]int

func minAreaRect(points [][]int) int {
	visited := map[p]bool{}
	yPoints := map[int][]p{}
	xPoints := map[int][]p{}
	for _, point := range points {
		tmp := p{point[0], point[1]}
		visited[tmp] = false
		xPoints[tmp[0]] = append(xPoints[tmp[0]], tmp)
		yPoints[tmp[1]] = append(yPoints[tmp[1]], tmp)
	}

	minArea := math.MaxInt
	for point := range visited {
		px, py := point[0], point[1]
		for _, adjPointX := range xPoints[px] {
			if adjPointX == point || visited[adjPointX] {
				continue
			}

			for _, adjPointY := range yPoints[py] {
				if adjPointY == point || visited[adjPointY] || adjPointX == adjPointY {
					continue
				}

				x, y := adjPointY[0], adjPointX[1]
				if _, ok := visited[p{x, y}]; ok {
					minArea = min(minArea, int(math.Abs(float64((px-x)*(py-y)))))
				}
			}
		}

		visited[point] = true
	}

	if minArea == math.MaxInt {
		return 0
	}

	return minArea
}

// func main() {
// 	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}))
// 	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}))
// 	fmt.Println(minAreaRect([][]int{{3, 3}, {4, 1}, {4, 3}}))
// 	fmt.Println(minAreaRect([][]int{{3, 2}, {1, 3}, {3, 0}, {3, 4}, {2, 1}, {0, 4}, {0, 3}, {4, 1}, {2, 4}}))
// }
