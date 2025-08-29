package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := map[int]map[int]bool{}
	for _, prerequisite := range prerequisites {
		_, exists := preMap[prerequisite[0]]
		if !exists {
			preMap[prerequisite[0]] = map[int]bool{}
		}
		preMap[prerequisite[0]][prerequisite[1]] = true
	}

	visited := map[int]bool{}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if visited[course] {
			return false // we found loop
		}
		if len(preMap[course]) == 0 {
			return true
		}
		visited[course] = true

		for preCourse := range preMap[course] {
			if !dfs(preCourse) {
				return false
			}
			delete(preMap[course], preCourse)
		}

		visited[course] = false

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
