package main

type Domino struct {
	index int
	value rune
}

func printFromTo(str []rune, from int, to int, value rune) []rune {
	for i := from; i <= to; i++ {
		str[i] = value
	}
	return str
}

func pushDominoes(dominoes string) string {
	dominoesAfter := []rune(dominoes)

	symbols := []Domino{Domino{index: 0, value: 'L'}}
	for i, v := range dominoes {
		if v != '.' {
			symbols = append(symbols, Domino{index: i, value: v})
		}
	}
	symbols = append(symbols, Domino{index: len(dominoes) - 1, value: 'R'})

	for i := 0; i < len(symbols)-1; i++ {
		cur, next := symbols[i], symbols[i+1]
		if cur.value == next.value {
			dominoesAfter = printFromTo(dominoesAfter, cur.index, next.index, next.value)
		} else if cur.value == 'R' && next.value == 'L' {
			dominoesAfter = printFromTo(dominoesAfter, cur.index, (cur.index+next.index)/2, 'R')
			dominoesAfter = printFromTo(dominoesAfter, ((cur.index+next.index)/2)+1, next.index, 'L')
			if (next.index-cur.index)%2 == 0 {
				dominoesAfter[(cur.index+next.index)/2] = '.'
			}
		}
	}

	return string(dominoesAfter)
}

// func main() {
// 	fmt.Println(pushDominoes("R"))
// 	fmt.Println(pushDominoes(".L.R."))
// 	fmt.Println(pushDominoes("RR.L"))
// 	fmt.Println(pushDominoes(".L.R...LR..L.."))
// }
