package main

import "fmt"

func main() {
	root := node{5, nil, nil}
	tree := tree{&root}
	tree.add(6)
	tree.add(2)
	tree.add(7)
	tree.add(4)
	tree.add(1)
	fmt.Println(find(tree.root, 4))
}

type tree struct {
	root *node
}

type node struct {
	value int
	left  *node
	right *node
}

func addRecursive(current *node, value int) *node {
	if current == nil {
		return &node{value, nil, nil}
	}
	if value < current.value {
		current.left = addRecursive(current.left, value)
	} else if value > current.value {
		current.right = addRecursive(current.right, value)
	} else {
		return current
	}
	return current
}

func (tree tree) add(value int) {
	tree.root = addRecursive(tree.root, value)
}

func dfs(current *node) {
	if current == nil {
		return
	}
	dfs(current.left)
	fmt.Println(current.value)
	dfs(current.right)
}

func find(current *node, value int) bool {
	if current == nil {
		return false
	}
	if current.value == value {
		return true
	}
	if value < current.value {
		return find(current.left, value)
	} else {
		return find(current.right, value)
	}
}

func bfs(current *node) {
	if current == nil {
		return
	}
	var queue []*node
	queue = append(queue, current)
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		fmt.Println(current.value)
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}
