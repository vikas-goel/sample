package main

import "fmt"
import "os"
import "strconv"

type List struct {
	key int
	next *List
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <prog> <pivot> <keys>...")
		os.Exit(1)
	}

	pivot, _ := strconv.Atoi(os.Args[1])

	var head *List
	for i := 2; i < len(os.Args); i++ {
		key, _ := strconv.Atoi(os.Args[i])
		head = head.insert(key, true)
	}

	head.Print()
	head.Partition(pivot)
	fmt.Print(" -> ")
	head.Print()
	fmt.Println()

	nums1 := create(9, 9, 9, 6, 1, 7)
	nums2 := create(4, 9, 5)
	sum12 := sumReverse(nums1, nums2)
	nums1.Print()
	fmt.Print(" + ")
	nums2.Print()
	fmt.Print(" = ")
	sum12.Print()
	fmt.Println()

	nums1 = create(7, 1, 6)
	nums2 = create(5, 9, 4)
	sum12 = sumForward(nums1, nums2)
	nums1.Print()
	fmt.Print(" + ")
	nums2.Print()
	fmt.Print(" = ")
	sum12.Print()
	fmt.Println()

	nums1 = create(7, 99, 6, 5, 2, 11, 42, 53)
	nums2 = create(11, 42, 22)
	nums2.next.next.next = nums1.next
	inter := intersect(nums1, nums2)

	nums1.Print()
	fmt.Print(" and ")
	nums2.Print()
	fmt.Print(" interect at ")
	if inter != nil {
		fmt.Println(inter.key)
	} else {
		fmt.Println("nil")
	}
}

func intersect(head1, head2 *List) *List {
	len1, len2 := 0, 0

	// Find length of each list.
	for n := head1; n != nil; n, len1 = n.next, len1+1 {}
	for n := head2; n != nil; n, len2 = n.next, len2+1 {}

	// No intersection if either list is empty.
	if len1 == 0 || len2 == 0 {
		return nil
	}

	// Find difference in the lengths and the longer list.
	diff := len1 - len2
	list1, list2 := head1, head2
	if diff < 0 {
		list1, list2 = head2, head1
		diff *= -1
	}

	// Move the longer list forward to match the length of the other.
	for ; diff != 0; diff-- {
		list1 = list1.next
	}

	// Traverse both lists in parallel to find the intersection.
	for ; list1 != nil; list1, list2 = list1.next, list2.next {
		if list1 == list2 {
			return list1
		}
	}

	return nil
}

func sumForward(head1, head2 *List) (head *List) {
	sum, factor := 0, 1
	for head1 != nil && head2 != nil {
		factor *= 10
		sum = sum*10 + head1.key + head2.key
		head1, head2 = head1.next, head2.next
	}

	for initial := true; factor > 0; factor /= 10 {
		digit := sum / factor
		if initial && digit == 0 {
			initial = false
			continue
		}

		initial = false
		sum %= factor
		head = head.insert(digit, false)
	}

	return
}

func sumReverse(head1, head2 *List) (head *List) {
	carry := 0
	for head1 != nil && head2 != nil {
		sum := head1.key + head2.key + carry
		carry = sum/10
		sum = sum % 10
		head = head.insert(sum, false)

		head1, head2 = head1.next, head2.next
	}

	rem := head1
	if head2 != nil {
		rem = head2
	}

	for ; rem != nil; rem = rem.next {
		sum := rem.key + carry
		carry = sum/10
		sum = sum % 10
		head = head.insert(sum, false)
	}

	if carry > 0 {
		head = head.insert(carry, false)
	}

	return
}

func sumInts(nums... int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return
}

func create(keys... int) (head *List) {
	for _, k := range keys {
		head = head.insert(k, true)
	}
	return
}

func (head *List) Partition(pivot int) {
	// Traverse while the first key is greater-than or equal-to the pivot.
	n1 := head
	for ; n1.key < pivot; n1 = n1.next {}

	// Start traversing from the next node.
	n2 := n1
	if n2.next != nil {
		n2 = n2.next
	}

	// Search for key less than pivot and swap with n1.
	for ; n2 != nil; n2 = n2.next {
		if n2.key < pivot {
			n1.key, n2.key = n2.key, n1.key
			n1 = n1.next
		}
	}
}

func (head *List) insert(key int, front bool) (newHead *List) {
	if front {
		newHead = newNode(key, head)
		return
	}

	newHead = newNode(key, nil)
	if head == nil {
		return
	}

	temp := head
	for ; temp.next != nil; temp = temp.next {}
	temp.next = newHead
	newHead = head

	return
}

func newNode(key int, next *List) (node *List) {
	node = new(List)
	node.key = key
	node.next = next
	return
}

func (head *List) Print() {
	fmt.Print("[")
	n := head
	for ; n != nil && n.next != nil; n = n.next {
		fmt.Print(n.key, " ")
	}

	if n != nil {
		fmt.Print(n.key)
	}

	fmt.Print("]")
}
