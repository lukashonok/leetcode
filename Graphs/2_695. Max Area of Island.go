package main

func maxAreaOfIsland(grid [][]int) int {
	visited := map[int]map[int]bool{}
	for i := 0; i < len(grid); i++ {
		visited[i] = map[int]bool{}
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = false
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	maxArea := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 || visited[i][j] {
				continue
			}

			stack, curArea := [][]int{{i, j}}, 0
			visited[i][j] = true
			for len(stack) != 0 {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				curArea++
				maxArea = max(maxArea, curArea)

				for _, d := range dirs {
					newI, newJ := cur[0]+d[0], cur[1]+d[1]
					if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && !visited[newI][newJ] && grid[newI][newJ] == 1 {
						visited[newI][newJ] = true
						stack = append(stack, []int{newI, newJ})
					}
				}
			}
		}
	}

	return maxArea
}

// func main() {
// 	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
// 	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}))
// 	fmt.Println(maxAreaOfIsland([][]int{{1}}))
// }
