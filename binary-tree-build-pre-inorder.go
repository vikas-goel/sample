// https://www.geeksforgeeks.org/construct-tree-from-given-inorder-and-preorder-traversal/

package main

import "fmt"

type node struct {
	data rune
	left, right *node
}

func indexOf(elem rune, inArray []rune, start, end int) int {
	for i := start; i <= end; i++ {
		if inArray[i] == elem {
			return i
		}
	}

	return -1
}

func buildTree(inOrder, preOrder []rune, start, end int, current *int) *node {
	if start > end {
		return nil
	}

	root := new(node)
	root.data = preOrder[*current]
	*current += 1

	if start == end {
		return root
	}

	pos := indexOf(root.data, inOrder, start, end)
	root.left = buildTree(inOrder, preOrder, start, pos-1, current)
	root.right = buildTree(inOrder, preOrder, pos+1, end, current)

	return root
}

func (root *node) printInorder() {
	if root == nil {
		return
	}

	root.left.printInorder()
	fmt.Printf("%c ", root.data)
	root.right.printInorder()
}

func (root *node) printPreorder() {
	if root == nil {
		return
	}

	fmt.Printf("%c ", root.data)
	root.left.printPreorder()
	root.right.printPreorder()
}

func (root *node) printPostorder() {
	if root == nil {
		return
	}

	root.left.printPostorder()
	root.right.printPostorder()
	fmt.Printf("%c ", root.data)
}

func main() {
	current := 0
	inOrder := []rune{'D', 'B', 'E', 'A', 'F', 'C' }
	preOrder := []rune{ 'A', 'B', 'D', 'E', 'C', 'F' }

	tree := buildTree(inOrder, preOrder, 0, len(inOrder)-1, &current)
	tree.printInorder()
	fmt.Println()
	tree.printPreorder()
	fmt.Println()
	tree.printPostorder()
}
