package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	clones := map[*Node]*Node{}

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		_, exists := clones[n]

		if exists {
			return clones[n]
		}

		clones[n] = &Node{Val: n.Val, Neighbors: []*Node{}}

		for _, nei := range n.Neighbors {
			clones[n].Neighbors = append(clones[n].Neighbors, dfs(nei))
		}

		return clones[n]
	}

	return dfs(node)
}
