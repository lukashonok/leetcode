package main

import (
	"fmt"
	"sort"
)

type Char struct {
	c byte
	i int
	j int
}

type Tr struct {
	data      map[byte]*Tr
	endOfWord bool
	word      string
}

func ConstructorTr() *Tr {
	return &Tr{data: map[byte]*Tr{}}
}

func (this *Tr) Insert(word string) {
	current := this
	for _, c := range word {
		value, exists := current.data[byte(c)]
		if !exists {
			newT := ConstructorTr()
			current.data[byte(c)], current = newT, newT
		} else {
			current = value
		}
	}
	current.endOfWord = true
	current.word = word
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	trie := ConstructorTr()
	for _, word := range words {
		trie.Insert(word)
	}

	// trie.PrintTrie("_")

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	inBounds := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(board) && j < len(board[0])
	}

	var dfs func(i, j int, node *Tr)
	dfs = func(i, j int, node *Tr) {
		char := board[i][j]
		// fmt.Println("Looking to " + string(char))
		newNode, exists := node.data[char]
		if exists && newNode.endOfWord {
			result = append(result, newNode.word)
			newNode.endOfWord = false
		} else if char == '.' || !exists {
			return
		}

		board[i][j] = '.'

		// printBoard(board)

		for _, dir := range dirs {
			_i, _j := i+dir[0], j+dir[1]
			if inBounds(_i, _j) {
				dfs(_i, _j, newNode)
			}
		}

		board[i][j] = char
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, trie)
		}
	}

	return result
}

func (t *Tr) PrintTrie(prefix string) {
	if t.endOfWord {
		fmt.Println(prefix) // print the current path as a word so far
	}
	// Sort keys so output order is consistent
	keys := make([]byte, 0, len(t.data))
	for k := range t.data {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, ch := range keys {
		child := t.data[ch]
		child.PrintTrie(prefix + string(ch))
	}
}

func printBoard(board [][]byte) {
	for i := range board {
		fmt.Println(string(board[i]))
	}
}

// func main() {
// 	fmt.Println(
// 		findWords(
// 			[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
// 			[]string{"oath", "pea", "eat", "rain", "oathi", "oathk", "oathf", "oate", "oathii", "oathfi", "oathfii"},
// 		),
// 	)

// 	// fmt.Println(
// 	// 	findWords(
// 	// 		[][]byte{{'a'}},
// 	// 		[]string{"a"},
// 	// 	),
// 	// )
// 	// fmt.Println(
// 	// 	findWords(
// 	// 		[][]byte{{'a'}},
// 	// 		[]string{"b"},
// 	// 	),
// 	// )
// 	// fmt.Println(
// 	// 	findWords(
// 	// 		[][]byte{{'a', 'b', 'c', 'e'}, {'x', 'x', 'c', 'd'}, {'x', 'x', 'b', 'a'}},
// 	// 		[]string{"abc", "abcd"},
// 	// 	),
// 	// )
// 	// fmt.Println(
// 	// 	findWords(
// 	// 		[][]byte{{'a', 'a'}},
// 	// 		[]string{"aaa"},
// 	// 	),
// 	// )
// 	// fmt.Println(
// 	// 	findWords(
// 	// 		[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
// 	// 		[]string{"oath", "pea", "eat", "rain"},
// 	// 	),
// 	// )
// }
