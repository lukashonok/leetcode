package main

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, a := range asteroids {
		for len(stack) != 0 && a < 0 && stack[len(stack)-1] > 0 {
			diff := a + stack[len(stack)-1]
			if diff < 0 {
				stack = stack[:len(stack)-1]
			} else if diff > 0 {
				a = 0
			} else {
				a = 0
				stack = stack[:len(stack)-1]
			}
		}
		if a != 0 {
			stack = append(stack, a)
		}
	}
	return stack
}

// func main() {
// 	fmt.Println(asteroidCollision([]int{5, 10, -5}))
// 	fmt.Println(asteroidCollision([]int{8, -8}))
// 	fmt.Println(asteroidCollision([]int{10, 2, -5}))
// 	fmt.Println(asteroidCollision([]int{-5, -5, 5}))
// }
