package main

import "fmt"

func main() {
	g := map[int][]int{0: {1, 2, 5}, 1: {0}, 2: {0, 3}, 3: {2}, 4: {5}, 5: {0, 4, 6}, 6: {5}}
	rootTree(g, 0)
	fmt.Println("hello")
}

type treeNode struct {
	id       int
	parent   *treeNode
	children []*treeNode
}

func rootTree(g map[int][]int, rootId int) treeNode {
	root := treeNode{rootId, nil, []*treeNode{}}
	return buildTree(g, root, treeNode{})
}

func buildTree(g map[int][]int, node treeNode, parent treeNode) treeNode {
	for childId := range g[node.id] {
		if childId == parent.id {
			continue
		}
		child := treeNode{childId, &node, []*treeNode{}}
		node.children = append(node.children, &child)
		buildTree(g, child, node)
	}
	return node
}
