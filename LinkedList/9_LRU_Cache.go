package main

type DequeElement struct {
	Value int
	Next  *DequeElement
	Prev  *DequeElement
}

type Deque struct {
	Head   *DequeElement
	Tail   *DequeElement
	Length int
}

func (deque *Deque) Unshift(value int) {
	newElement := &DequeElement{Value: value}
	if deque.Head == nil {
		deque.Head = newElement
		deque.Tail = deque.Head
		return
	}

	deque.Tail.Next = newElement
	newElement.Prev = deque.Tail
	deque.Tail = newElement
	deque.Length++
}

func (deque *Deque) Shift() *int {
	if deque.Head == nil {
		return nil
	}
	oldHead := deque.Head
	deque.Head = deque.Head.Next
	if deque.Head == nil {
		deque.Tail = nil
	} else {
		deque.Head.Prev = nil
	}
	return &oldHead.Value
}

type LRUCache struct {
	cache    map[int]int
	keyList  *Deque
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    map[int]int{},
		keyList:  &Deque{},
	}
}

func (this *LRUCache) Get(key int) int {
	value, exists := this.cache[key]
	if !exists {
		return -1
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {
	if this.keyList.Length == this.capacity {
		last := this.keyList.Shift()
		if last != nil {
			delete(this.cache, *last)
		}
	}
	this.keyList.Unshift(key)
	this.cache[key] = value
}

func runCommands(commands []string, args [][]int) []any {
	var obj LRUCache
	results := []any{}
	for i := range commands {
		args := args[i]
		switch commands[i] {
		case "get":
			results = append(results, obj.Get(args[0]))
		case "put":
			obj.Put(args[0], args[1])
			results = append(results, nil)
		case "LRUCache":
			obj = Constructor(args[0])
			results = append(results, nil)
		}
	}
	return results
}

// func main() {
// 	fmt.Println(
// 		runCommands(
// 			[]string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
// 			[][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
// 		)...,
// 	)
// }
