package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	buffer := ""

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current == nil {
			buffer += ","
			continue
		}

		buffer += strconv.Itoa(current.Val) + ","

		left, right := current.Left, current.Right

		queue = append(queue, left)
		queue = append(queue, right)
	}

	i := len(buffer) - 1
	for ; i > 0; i-- {
		if buffer[i] != ',' {
			break
		}
	}
	buffer = buffer[:i+1]

	return buffer
}

// 1,2,3,n,4,n,n
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	elements := strings.Split(data, ",")

	rootValue, _ := strconv.Atoi(elements[0])
	root := &TreeNode{Val: rootValue}
	queue := []*TreeNode{root}

	for i := 1; i < len(elements); i += 2 {
		current := queue[0]
		queue = queue[1:]

		if elements[i] != "" {
			value, _ := strconv.Atoi(elements[i])
			current.Left = &TreeNode{Val: value}
			queue = append(queue, current.Left)
		}

		if i+1 < len(elements) && elements[i+1] != "" {
			value, _ := strconv.Atoi(elements[i+1])
			current.Right = &TreeNode{Val: value}
			queue = append(queue, current.Right)
		}
	}

	return root
}

func main() {
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(decodeTree([]any{1, 2, 3, nil, 4, nil, nil}))
	ans := deser.deserialize(data)
	fmt.Println(ans)
}
