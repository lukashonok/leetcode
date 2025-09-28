package main

type MyHashSet struct {
	List [1000001]bool
}

// func Constructor() MyHashSet {
// 	return MyHashSet{List: [1000001]bool{}}
// }

func (this *MyHashSet) Add(key int) {
	this.List[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.List[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.List[key]
}

// func main() {
// 	myHashSet := Constructor()
// 	myHashSet.Add(1)                   // set = [1]
// 	myHashSet.Add(2)                   // set = [1, 2]
// 	fmt.Println(myHashSet.Contains(1)) // return True
// 	fmt.Println(myHashSet.Contains(3)) // return False, (not found)
// 	myHashSet.Add(2)                   // set = [1, 2]
// 	fmt.Println(myHashSet.Contains(2)) // return True
// 	myHashSet.Remove(2)                // set = [1]
// 	fmt.Println(myHashSet.Contains(2)) // return False, (already removed)
// }
