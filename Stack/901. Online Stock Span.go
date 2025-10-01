package main

type StockSpanner struct {
	stack [][2]int // 0 - value, 1 - counter
}

func Constructor() StockSpanner {
	return StockSpanner{stack: [][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for len(this.stack) != 0 && price >= this.stack[len(this.stack)-1][0] {
		count += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, count})
	return this.stack[len(this.stack)-1][1]
}

// func main() {
// 	stock := Constructor()
// 	fmt.Println(stock.Next(100))
// 	fmt.Println(stock.Next(80))
// 	fmt.Println(stock.Next(60))
// 	fmt.Println(stock.Next(70))
// 	fmt.Println(stock.Next(60))
// 	fmt.Println(stock.Next(75))
// 	fmt.Println(stock.Next(85))
// }
