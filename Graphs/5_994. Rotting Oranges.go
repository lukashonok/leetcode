package main

func orangesRotting(grid [][]int) int {
	visited, fresh, rottenQueue := map[int]map[int]bool{}, 0, [][]int{}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < len(grid); i++ {
		visited[i] = map[int]bool{}
		for j := 0; j < len(grid[0]); j++ {
			visited[i][j] = false
			switch grid[i][j] {
			case 1:
				fresh++
			case 2:
				visited[i][j] = true
				rottenQueue = append(rottenQueue, []int{i, j})
			}
		}
	}

	minutes := 0
	for len(rottenQueue) != 0 {
		somethingWasRotten := false

		for range rottenQueue {
			cur := rottenQueue[0]
			rottenQueue = rottenQueue[1:]

			for _, d := range directions {
				newI, newJ := cur[0]+d[0], cur[1]+d[1]
				if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] == 1 && !visited[newI][newJ] {
					grid[newI][newJ] = 2
					fresh--
					somethingWasRotten = true
					visited[newI][newJ] = true
					rottenQueue = append(rottenQueue, []int{newI, newJ})
				}
			}
		}

		if somethingWasRotten {
			minutes++
		}

	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

// func main() {
// 	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
// 	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
// 	fmt.Println(orangesRotting([][]int{{0, 2}}))
// }
