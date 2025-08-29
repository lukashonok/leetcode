package main

type WordDictionary struct {
	data      map[rune]*WordDictionary
	endOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{data: map[rune]*WordDictionary{}}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for _, c := range word {
		value, exists := current.data[c]
		if !exists {
			newWordDictionary := Constructor()
			current.data[c], current = &newWordDictionary, &newWordDictionary
		} else {
			current = value
		}
	}
	current.endOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	current := this

	for i, c := range word {
		value, exists := current.data[c]

		if c == '.' {
			atleastOne := false
			for _, possibleDictionary := range current.data {
				atleastOne = possibleDictionary.Search(word[i+1:])
				if atleastOne {
					return true
				}
			}
		}

		if !exists {
			return false
		}

		current = value
	}

	return current.endOfWord
}

// func main() {
// 	wordDictionary := Constructor()
// 	wordDictionary.AddWord("bad")
// 	wordDictionary.AddWord("dad")
// 	wordDictionary.AddWord("mad")
// 	fmt.Println(wordDictionary.Search("pad")) // return False
// 	fmt.Println(wordDictionary.Search("bad")) // return True
// 	fmt.Println(wordDictionary.Search(".ad")) // return True
// 	fmt.Println(wordDictionary.Search("b..")) // return True
// }
