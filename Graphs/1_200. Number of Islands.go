package main

func numIslands(grid [][]byte) int {
	visited := map[int]map[int]bool{}
	for i := 0; i < len(grid); i++ {
		visited[i] = map[int]bool{}
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = false
		}
	}

	result, dirs := 0, [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				visited[i][j] = true
			}
			if visited[i][j] {
				continue
			}

			stack := [][]int{{i, j}}
			for len(stack) != 0 {
				currentPoint := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				visited[currentPoint[0]][currentPoint[1]] = true

				for _, d := range dirs {
					newI, newJ := currentPoint[0]+d[0], currentPoint[1]+d[1]
					if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] != '0' && !visited[newI][newJ] {
						stack = append(stack, []int{newI, newJ})
					}
				}
			}
			result++
		}
	}

	return result
}

// func main() {
// 	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
// 	fmt.Println(numIslands([][]byte{
// 		{'1', '1', '0', '0', '0'},
// 		{'1', '1', '0', '0', '0'},
// 		{'0', '0', '1', '0', '0'},
// 		{'0', '0', '0', '1', '1'},
// 	}))
// }
