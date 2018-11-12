package main

import "fmt"
import "math"

type heap struct {
	keys []int
	size, capacity int
}

func heapNew(capacity int) (h *heap) {
	h = new(heap)
	h.capacity, h.size = capacity, 0
	h.keys = make([]int, h.capacity)
	return
}

func heapParent(h *heap, idx int) (pid int, ok bool) {
	if h == nil || idx >= h.size {
		return
	}

	pid, ok = (idx-1)/2, true
	return
}

func heapLeftChild(h *heap, idx int) (pid int, ok bool) {
	if h == nil || idx >= h.size {
		return
	}

	pid = idx*2+1
	if pid < h.size {
		ok = true
	}

	return
}

func heapRightChild(h *heap, idx int) (pid int, ok bool) {
	if h == nil || idx >= h.size {
		return
	}

	pid = idx*2+2
	if pid < h.size {
		ok = true
	}

	return
}

func heapAdjust(h *heap, idx int) {
	if h == nil || idx <= 0 || idx >= h.size {
		return
	}

	for i := idx; i != 0; {
		pidx, _ := heapParent(h, i)
		if h.keys[pidx] > h.keys[i] {
			h.keys[pidx], h.keys[i] = h.keys[i], h.keys[pidx]
			i = pidx
		} else {
			break
		}
	}
}

func heapify(h *heap, idx int) {
	if idx >= h.size {
		return
	}

	min := idx

	cid, ok := heapLeftChild(h, idx)
	if ok && h.keys[cid] < h.keys[min] {
		min = cid
	}

	cid, ok = heapRightChild(h, idx)
	if ok && h.keys[cid] < h.keys[min] {
		min = cid
	}

	if min != idx {
		h.keys[idx], h.keys[min] = h.keys[min], h.keys[idx]
		heapify(h, min)
	}
}

func heapInsert(h *heap, key int) bool {
	if h.size == h.capacity {
		return false
	}

	h.keys[h.size] = key
	h.size++

	heapAdjust(h, h.size-1)

	return true
}

func heapDelete(h *heap, idx int) bool {
	if h == nil || idx >= h.size {
		return false
	}

	heapDecreaseKey(h, idx, math.MinInt32)
	heapExtractMin(h)

	return true
}

func heapDecreaseKey(h *heap, idx, key int) bool {
	if h == nil || idx >= h.size || h.keys[idx] <= key {
		return false
	}

	h.keys[idx] = key
	heapAdjust(h, idx)

	return true
}

func heapExtractMin(h *heap) (min int, ok bool) {
	if h == nil || h.size == 0 {
		return
	} else if h.size == 1 {
		h.size--
		min = h.keys[h.size]
		ok = true
		return
	}

	h.size--
	min = h.keys[0]
	h.keys[0] = h.keys[h.size]
	ok = true
	heapify(h, 0)

	return
}

func heapPrint(h *heap) {
	fmt.Println(h.keys)
}

func buildHeap(keys ...int) (h *heap) {
	h = heapNew(len(keys))
	for _, k := range keys {
		heapInsert(h, k)
	}
	return
}

func main() {
	h := heapNew(11)
	heapInsert(h, 3)
	heapPrint(h)
	heapInsert(h, 2)
	heapPrint(h)
	heapDelete(h, 1)
	heapPrint(h)
	heapInsert(h, 15)
	heapPrint(h)
	heapInsert(h, 5)
	heapPrint(h)
	heapInsert(h, 4)
	heapPrint(h)
	heapInsert(h, 45)
	heapPrint(h)
	heapExtractMin(h)
	heapPrint(h)
	heapDecreaseKey(h, 2, 1)
	heapPrint(h)
}
