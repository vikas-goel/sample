package main

import "fmt"
import "math"
import "sort"

type node struct {
	key int
	left, right, peer *node
}

func max(n ...int) (m int) {
	m = 1 >> 32
	for _, i := range n {
		if i > m {
			m = i
		}
	}

	return
}

func (root *node) height(diameter *int) int {
	if root == nil {
		return 0
	}

	lheight := root.left.height(diameter)
	rheight := root.right.height(diameter)

	if diameter != nil {
		thisDiag := 1 + lheight + rheight
		if thisDiag > *diameter {
			*diameter = thisDiag
		}
	}

	return 1 + max(lheight, rheight)
}

func (root *node) diameter() (diam int) {
	if root == nil {
		return 0
	}

	root.height(&diam)
	return
}

func (root *node) connectPeers() {
	if root == nil {
		return
	}

	pos, length := 0, 1
	queue := make([]*node, int(math.Pow(2, float64(root.height(nil))))-1)
	queue[pos] = root

	for pos < length {
		count := length - pos
		fmt.Print("{ ")
		for ; count > 0; count-- {
			fmt.Print(queue[pos].key, " ")

			if count > 1 {
				queue[pos].peer = queue[pos+1]
			}

			if queue[pos].left != nil {
				queue[length] = queue[pos].left
				length++
			}

			if queue[pos].right != nil {
				queue[length] = queue[pos].right
				length++
			}

			pos++
		}
		fmt.Print("}")
	}
}

func (this *node) connectPeersPerfect() {
	if this == nil {
		return
	}

	for first = this; first.left != nil; first = first.left {
		for current, last := first, nil; current != nil; current = current.peer {
			if last != nil {
				last.peer = current.left
			}

			current.left.peer = current.right
			last = current.right
		}
	}
}

func (root *node) printPeers() {
	for ; root != nil; root = root.peer {
		fmt.Print(root.key, " ")
	}
}

func (root *node) convertToDLL(origin bool) *node {
	if root == nil {
		return root
	}

	if root.left != nil {
		// Get left subtree converted.
		left := root.left.convertToDLL(false)

		// Adjust root and its left subtree relation.
		for ; left.right != nil; left = left.right {}
		left.right = root
		root.left = left
	}

	if root.right != nil {
		// Get right subtree converted.
		right := root.right.convertToDLL(false)

		// Adjust root and its right subtree relation.
		for ; right.left != nil; right = right.left {}
		right.left = root
		root.right = right
	}

	if origin {
		// Move root pointer to the beginning of the list.
		for ; root.left != nil; root = root.left {}
	}

	return root
}

func (root *node) maxWidthByOrder(width *[]int, level int, mwidth *int) int {
	if root == nil {
		return 0
	}

	if width == nil {
		width = new([]int)
		*width = make([]int, root.height(nil))
	}

	if mwidth == nil {
		mwidth = new(int)
		*mwidth = 0
	}

	(*width)[level]++
	if (*width)[level] > *mwidth {
		*mwidth = (*width)[level]
	}

	if root.left != nil {
		root.left.maxWidthByOrder(width, level+1, mwidth)
	}

	if root.right != nil {
		root.right.maxWidthByOrder(width, level+1, mwidth)
	}

	return *mwidth
}

func (root *node) maxWidthByQueue() (mwidth int) {
	if root == nil {
		return
	}

	type nheight struct {
		n *node
		h int
	}

	height := root.height(nil)
	hwidth := make([]int, height)
	tqueue := make([]nheight, int(math.Pow(2, float64(height)))-1)

	pos, length := 0, 1
	tqueue[pos].n = root
	tqueue[pos].h = 0

	for pos = 0; pos < length; pos++ {
		// Increase the width of this tree height.
		hwidth[tqueue[pos].h]++
		if hwidth[tqueue[pos].h] > mwidth {
			mwidth = hwidth[tqueue[pos].h]
		}

		if tqueue[pos].n.left != nil {
			tqueue[length].n = tqueue[pos].n.left
			tqueue[length].h = tqueue[pos].h+1
			length++
		}

		if tqueue[pos].n.right != nil {
			tqueue[length].n = tqueue[pos].n.right
			tqueue[length].h = tqueue[pos].h+1
			length++
		}
	}

	return
}

func (root *node) printInorder() {
	if root == nil {
		return
	}

	root.left.printInorder()
	fmt.Print(root.key, " ")
	root.right.printInorder()
}

func (root *node) printPreorder() {
	if root == nil {
		return
	}

	fmt.Print(root.key, " ")
	root.left.printPreorder()
	root.right.printPreorder()
}

func (root *node) printPostorder() {
	if root == nil {
		return
	}

	root.left.printPostorder()
	root.right.printPostorder()
	fmt.Print(root.key, " ")
}

func (root *node) printLevelorder() {
	if root == nil {
		return
	}

	queue := make([]*node, int(math.Pow(2, float64(root.height(nil))))-1)
	queue[0] = root
	pos, length := 0, 1

	for pos < length {
		fmt.Print(queue[pos].key, " ")

		if queue[pos].left != nil {
			queue[length] = queue[pos].left
			length++
		}

		if queue[pos].right != nil {
			queue[length] = queue[pos].right
			length++
		}

		pos++
	}
}

func (root *node) printDLL() {
	for ; root != nil; root = root.right {
		fmt.Print(root.key, " ")
	}
}

func (root *node) printAncestors(key int) bool {
	if root == nil {
		return false
	}

	if root.key == key {
		return true
	}

	if root.left.printAncestors(key) || root.right.printAncestors(key) {
		fmt.Print(root.key, " ")
		return true
	}

	return false
}

func (root *node) printVerticals(verticals *map[int][]int, distance int) {
	if root == nil {
		return
	}

	print := false
	if verticals == nil {
		verticals = new(map[int][]int)
		*verticals = make(map[int][]int)
		print = true
	}

	(*verticals)[distance] = append((*verticals)[distance], root.key)
	root.left.printVerticals(verticals, distance-1)
	root.right.printVerticals(verticals, distance+1)

	if !print {
		return
	}

	var keys []int
	for k := range (*verticals) {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range(keys) {
		fmt.Print((*verticals)[key], " ")
	}
}

func newNode(key int) (n *node) {
	n = new(node)
	n.key = key
	return
}

func (root *node) isBST(left, right *node) bool {
	if root == nil {
		return true
	}

	if left != nil && root.key <= left.key {
		return false
	}

	if right != nil && root.key > right.key {
		return false
	}

	return root.left.isBST(left, root) && root.right.isBST(root, right)
}

func (root *node) searchLCA(key1, key2 int) (lca int) {
	/*
	if root == nil {
		return
	}
	*/

	if key1 > key2 {
		key1, key2 = key2, key1
	}

	for root != nil {
		if key2 < root.key {
			root = root.left
		} else if key1 > root.key {
			root = root.right
		} else {
			lca = root.key
			break
		}
	}

	/*
	if key2 < root.key {
		lca = root.left.searchLCA(key1, key2)
	} else if key1 > root.key {
		lca = root.right.searchLCA(key1, key2)
	} else {
		lca = root.key
	}
	*/

	return
}

func (root *node) sumPair(sum int) (key1, key2 int, found bool) {
	if root == nil {
		return
	}

	length1, length2 := 0, 0
	curr1, curr2 := root, root
	done1, done2 := false, false
	stack1 := make([]*node, root.height(nil))
	stack2 := make([]*node, root.height(nil))

	for !found {
		for !done1 {
			if curr1 != nil {
				stack1[length1] = curr1
				length1++
				curr1 = curr1.left
			} else {
				if length1 != 0 {
					length1--
					key1 = stack1[length1].key
					curr1 = stack1[length1].right

				}
				done1 = true
			}
		}

		for !done2 {
			if curr2 != nil {
				stack2[length2] = curr2
				length2++
				curr2 = curr2.right
			} else {
				if length2 != 0 {
					length2--
					key2 = stack2[length2].key
					curr2 = stack2[length2].left
				}
				done2 = true
			}
		}

		if key1 >= key2 {
			break
		} else if key1 + key2 == sum {
			found = true
		} else if key1 + key2 < sum {
			done1 = false
		} else if key1 + key2 > sum {
			done2 = false
		}
	}

	return
}

func (root *node) correction(first, middle, last, prev **node) {
	if root == nil {
		return
	}

	swap := false
	if first == nil {
		first, middle, last = new(*node), new(*node), new(*node)
		prev = new(*node)
		swap = true
	}

	root.left.correction(first, middle, last, prev)

	if *prev != nil && (*prev).key > root.key {
		if *first == nil {
			*first = *prev
			*middle = root
		} else {
			*last = root
		}
	}
	*prev = root

	root.right.correction(first, middle, last, prev)

	if !swap {
		return
	}

	if *first != nil && *last != nil {
		(*first).key, (*last).key = (*last).key, (*first).key
	} else if *first != nil && *middle != nil {
		(*first).key, (*middle).key = (*middle).key, (*first).key
	}
}

func (root *node) search(key int) (bool, *node, *node) {
	if root == nil {
		return false, nil, nil
	} else if key == root.key {
		return true, root, nil
	}

	if key < root.key {
		if root.left == nil {
			return false, nil, root
		} else if root.left.key == key {
			return true, root.left, root
		}
		return root.left.search(key)
	} else {
		if root.right == nil {
			return false, nil, root
		} else if root.right.key == key {
			return true, root.right, root
		}
		return root.right.search(key)
	}
}

func (root *node) successor() (*node, *node) {
	if root == nil || root.left == nil {
		return root, nil
	}

	for ; root.left.left != nil; root = root.left {}

	return root.left, root
}

func (root *node) insert(key int) (bool, *node) {
	present, n, p := root.search(key)

	if present || p == nil {
		return false, n
	}

	if key < p.key {
		p.left = newNode(key)
		return true, p.left
	} else {
		p.right = newNode(key)
		return true, p.right
	}
}

func (root *node) delete(key int) (bool, *node) {
	if root == nil {
		return false, nil
	} else if key == root.key {
		if root.left == nil {
			return true, root.right
		} else if root.right == nil {
			return true, root.left
		}

		succ, psucc := root.right.successor()
		succ.left = root.left
		if psucc != nil {
			psucc.left = succ.right
		}

		if root.right != succ {
			succ.right = root.right
		}
		root.left, root.right = nil, nil
		return true, succ
	}

	if key < root.key {
		ok, newroot := root.left.delete(key)
		if ok {
			root.left = newroot
		}
	} else if key > root.key {
		ok, newroot := root.right.delete(key)
		if ok {
			root.right = newroot
		}
	}

	return true, root
}

func binTree1() (root *node) {
	root = newNode(10)
	root.left = newNode(12)
	root.right = newNode(15)
	root.left.left = newNode(25)
	root.left.right = newNode(30)
	root.right.left = newNode(36)
	return
}

func binTree2() (root *node) {
	root = newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.left.right.left = newNode(6)
	root.left.right.right = newNode(7)
	root.left.left.right = newNode(8)
	root.left.left.right.left = newNode(9)
	return
}

func binTree3() (root *node) {
	root = newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.left.left.left = newNode(7)
	return
}

func binTree4() (root *node) {
	root = newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.right.right = newNode(8)
	root.right.right.left = newNode(6)
	root.right.right.right = newNode(7)
	return
}

func binTree5() (root *node) {
	root = newNode(3)
	root.left = newNode(2)
	root.right = newNode(5)
	root.left.left = newNode(1)
	root.left.right = newNode(4)
	return
}

func binTree6() (root *node) {
	root = newNode(6)
	root.left = newNode(10)
	root.right = newNode(2)
	root.left.left = newNode(1)
	root.left.right = newNode(3)
	root.right.right = newNode(12)
	root.right.left = newNode(7)
	return
}

func binTree7() (root *node) {
	root = newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.right.left = newNode(6)
	root.right.right = newNode(7)
	root.right.left.right = newNode(8)
	root.right.right.right = newNode(9)
	return
}

func bsTree(keys ...int) (root *node) {
	if len(keys) == 0 {
		keys = []int{50, 30, 20, 40, 70, 60, 80}
	}

	root = newNode(keys[0])
	for _, i := range keys[1:] {
		root.insert(i)
	}
	return
}

func main() {
	roots := []*node{binTree1(), binTree2(), binTree3(), binTree4(), binTree7(), bsTree()}

	for _, root := range roots {
		fmt.Printf("%11v: %v\n", "Height", root.height(nil))
		fmt.Printf("%11v: %v\n", "Diameter", root.diameter())
		fmt.Printf("%11v: %v\n", "Max Q Width", root.maxWidthByQueue())

		fmt.Printf("%11v: ", "Max O Width")
		fmt.Println(root.maxWidthByOrder(nil, 0, nil))

		fmt.Printf("%11v: %v\n", "Is BST", root.isBST(nil, nil))

		fmt.Printf("%11v: { ", "Ancestor(7)")
		root.printAncestors(7)
		fmt.Println("}")

		fmt.Printf("%11v: ", "All Peers")
		root.connectPeers()
		fmt.Println()

		fmt.Printf("%11v: { ", "Peers")
		root.printPeers()
		fmt.Print("}{ ")
		root.left.printPeers()
		fmt.Print("}{ ")
		root.left.left.printPeers()
		fmt.Println("}")

		fmt.Printf("%11v: { ", "Level-order")
		root.printLevelorder()
		fmt.Println("}")

		fmt.Printf("%11v: { ", "Pre-order")
		root.printPreorder()
		fmt.Println("}")

		fmt.Printf("%11v: { ", "Post-order")
		root.printPostorder()
		fmt.Println("}")

		fmt.Printf("%11v: { ", "In-order")
		root.printInorder()
		fmt.Println("}")

		fmt.Printf("%11v: { ", "Verticals")
		root.printVerticals(nil, 0)
		fmt.Println("}")

		fmt.Printf("%11v: { ", "List")
		root.convertToDLL(true).printDLL()
		fmt.Println("}")
		fmt.Println("---")

	}

	root := bsTree()
	for _, k := range []int{0, 70, 20, 50, 80, 30, 60, 40} {
		fmt.Printf("%8v %2v: { ", "Delete", k)
		_, root = root.delete(k)
		root.printInorder()
		fmt.Println("}")
	}

	fmt.Println("---")
	root = bsTree(4, 2, 5, 1, 3)
	fmt.Printf("%11v: { ", "In-order")
	root.printInorder()
	fmt.Println("}")
	fmt.Printf("%11v: %v\n", "Is BST", root.isBST(nil, nil))

	root = binTree5()
	fmt.Printf("%11v: { ", "In-order")
	root.printInorder()
	fmt.Println("}")
	fmt.Printf("%11v: %v\n", "Is BST", root.isBST(nil, nil))
	fmt.Println("---")

	root = bsTree(20, 8, 22, 4, 12, 10, 14)
	fmt.Printf("%11v: { ", "In-order")
	root.printInorder()
	fmt.Println("}")
	fmt.Printf("%11v: %v\n", "LCA(10, 14)", root.searchLCA(10, 14))
	fmt.Printf("%11v: %v\n", "LCA(14, 08)", root.searchLCA(14, 8))
	fmt.Printf("%11v: %v\n", "LCA(10, 22)", root.searchLCA(10, 22))
	fmt.Println("---")

	root = bsTree(15, 10, 20, 8, 12, 16, 25)
	fmt.Printf("%11v: { ", "In-order")
	root.printInorder()
	fmt.Println("}")
	key1, key2, ok := root.sumPair(35)
	fmt.Printf("%11v: ", "Sum Pair 35")
	if ok {
		fmt.Printf("(%v, %v)\n", key1, key2)
	} else {
		fmt.Printf("Not found.\n")
	}
	fmt.Println("---")

	root = binTree6()
	fmt.Printf("%11v: { ", "In-order")
	root.printInorder()
	fmt.Println("}")
	root.correction(nil, nil, nil, nil)
	fmt.Printf("%11v: { ", "Correction")
	root.printInorder()
	fmt.Println("}")
	fmt.Println("---")
}
