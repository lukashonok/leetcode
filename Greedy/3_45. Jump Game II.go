package main

func jump(nums []int) int {
	res, l, r := 0, 0, 0
	for r < len(nums)-1 {
		farthest := 0
		for i := l; i <= r; i++ {
			farthest = max(farthest, i+nums[i])
		}
		l, r = r+1, farthest
		res++
	}
	return res
}

// func main() {
// 	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
// }
