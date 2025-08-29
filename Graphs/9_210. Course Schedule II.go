package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap := map[int]map[int]bool{}
	for _, prerequisite := range prerequisites {
		_, exists := preMap[prerequisite[0]]
		if !exists {
			preMap[prerequisite[0]] = map[int]bool{}
		}
		preMap[prerequisite[0]][prerequisite[1]] = true
	}

	visited := map[int]bool{}
	ordered := map[int]bool{}
	order := []int{}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if visited[course] {
			return false // we found loop, exit
		}
		if len(preMap[course]) == 0 {
			if !ordered[course] {
				order = append(order, course)
				ordered[course] = true
			}
			return true
		}
		visited[course] = true

		for preCourse := range preMap[course] {
			if !dfs(preCourse) {
				return false
			}
			delete(preMap[course], preCourse)
		}

		if !ordered[course] {
			order = append(order, course)
			ordered[course] = true
		}

		visited[course] = false

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return []int{}
		}
	}

	return order
}

// func main() {
// 	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
// }
