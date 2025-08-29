package main

import (
	"fmt"
	"slices"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	adj, wordLength := map[string]map[string]bool{}, len(beginWord)
	if slices.Index(wordList, endWord) == -1 {
		return 0
	}
	wordList = append(wordList, beginWord)

	for _, word := range wordList {
		for i := 0; i < wordLength; i++ {
			pattern := word[:i] + "*" + word[i+1:]
			_, exist := adj[pattern]
			if !exist {
				adj[pattern] = map[string]bool{}
			}
			adj[pattern][word] = true
		}
	}
	for pattern := range adj {
		if len(adj[pattern]) == 1 {
			delete(adj, pattern)
		}
	}

	queue, visited, res := []string{beginWord}, map[string]bool{beginWord: true}, 1
	for len(queue) != 0 {
		for _, word := range queue {
			queue = queue[1:]
			if word == endWord {
				return res
			}
			for i := 0; i < wordLength; i++ {
				pattern := word[:i] + "*" + word[i+1:]
				_, exists := adj[pattern]
				if !exists {
					continue
				}

				for innerWord := range adj[pattern] {
					if !visited[innerWord] {
						visited[innerWord] = true
						queue = append(queue, innerWord)
					}
				}
			}
		}
		res++
	}

	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Println(ladderLength("cat", "sag", []string{"bat", "bag", "sag", "dag", "dot"}))
}
