package main

import (
	"fmt"
	"math"
)

type p [2]int

type DetectSquares struct {
	m map[p]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		m: map[p]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	this.m[p{point[0], point[1]}]++
}

func (this *DetectSquares) Count(query []int) int {
	total := 0
	px, py := query[0], query[1]
	for point := range this.m {
		x, y := point[0], point[1]
		if math.Abs(float64(py-y)) != math.Abs(float64(px-x)) || x == px || y == py {
			continue
		}
		total += this.m[p{x, py}] * this.m[p{px, y}] * this.m[p{x, y}]
	}
	return total
}

func main() {
	obj := Constructor()
	obj.Add([]int{3, 10})
	obj.Add([]int{11, 2})
	obj.Add([]int{3, 2})
	fmt.Println(obj.Count([]int{11, 10}))
	fmt.Println(obj.Count([]int{14, 8}))
	obj.Add([]int{11, 2})
	fmt.Println(obj.Count([]int{11, 10}))
}
