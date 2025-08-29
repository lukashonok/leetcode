package main

type Trie struct {
	data      map[rune]*Trie
	endOfWord bool
}

func Constructor1() Trie {
	return Trie{data: map[rune]*Trie{}}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, c := range word {
		value, exists := current.data[c]
		if !exists {
			newTrie := Constructor1()
			current.data[c], current = &newTrie, &newTrie
		} else {
			current = value
		}
	}
	current.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	current := this

	for _, c := range word {
		value, exists := current.data[c]

		if !exists {
			return false
		}

		current = value
	}

	return current.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this

	for _, c := range prefix {
		value, exists := current.data[c]

		if !exists {
			return false
		}

		current = value
	}

	return true
}

// func main() {
// 	word := "apply"
// 	obj := Constructor1()
// 	obj.Insert(word)
// 	param_2 := obj.Search(word)
// 	param_3 := obj.StartsWith("app")
// 	fmt.Println(
// 		param_2,
// 		param_3,
// 		obj.StartsWith("apr"),
// 		obj.Search("app"),
// 		obj.StartsWith(""),
// 	)
// }
