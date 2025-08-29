package main

func canCompleteCircuit(gas []int, cost []int) int {
	gasSum, costSum, length := 0, 0, len(gas)
	for i := 0; i < length; i++ {
		gasSum += gas[i]
		costSum += cost[i]
	}
	if gasSum < costSum {
		return -1
	}

	total, start := 0, 0
	for i := 0; i < length; i++ {
		total += gas[i] - cost[i]
		if total < 0 {
			total, start = 0, i+1
		}
	}
	return start
}

// func main() {
// 	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
// 	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4}, []int{2, 2, 4, 1}))

// 	fmt.Println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))
// }
