package main

import "fmt"
import "math"

type heap struct {
	keys []int
	size, capacity int
}

func HeapNew(capacity int) (h *heap) {
	h = new(heap)
	h.capacity, h.size = capacity, 0
	h.keys = make([]int, h.capacity)
	return
}

func (h *heap) parentPosition(idx int) (pid int, ok bool) {
	if h == nil || idx >= h.size {
		return
	}

	pid, ok = (idx-1)/2, true
	return
}

func (h *heap) leftChildPosition(idx int) (cid int, ok bool) {
	if h == nil || idx >= h.size {
		return
	}

	cid = idx*2+1
	if cid < h.size {
		ok = true
	}

	return
}

func (h *heap) rightChildPosition(idx int) (cid int, ok bool) {
	if h == nil || idx >= h.size {
		return
	}

	cid = idx*2+2
	if cid < h.size {
		ok = true
	}

	return
}

func (h *heap) heapifyUp(idx int) {
	if h == nil || idx <= 0 || idx >= h.size {
		return
	}

	for idx != 0 {
		pidx, _ := h.parentPosition(idx)
		if h.keys[pidx] > h.keys[idx] {
			h.keys[pidx], h.keys[idx] = h.keys[idx], h.keys[pidx]
			idx = pidx
		} else {
			break
		}
	}
}

func (h *heap) heapifyDown(idx int) {
	if idx >= h.size {
		return
	}

	min := idx

	cid, ok := h.leftChildPosition(idx)
	if ok && h.keys[cid] < h.keys[min] {
		min = cid

		cid, ok = h.rightChildPosition(idx)
		if ok && h.keys[cid] < h.keys[min] {
			min = cid
		}
	}

	if min != idx {
		h.keys[idx], h.keys[min] = h.keys[min], h.keys[idx]
		h.heapifyDown(min)
	}
}

func (h *heap) Insert(key int) bool {
	if h.size == h.capacity {
		return false
	}

	h.keys[h.size] = key
	h.size++

	h.heapifyUp(h.size-1)

	return true
}

func (h *heap) Delete(idx int) bool {
	if h == nil || idx >= h.size {
		return false
	}

	h.DecreaseKey(idx, math.MinInt32)
	h.ExtractMin()

	return true
}

func (h *heap) DecreaseKey(idx, key int) bool {
	if h == nil || idx >= h.size || h.keys[idx] <= key {
		return false
	}

	h.keys[idx] = key
	h.heapifyUp(idx)

	return true
}

func (h *heap) ExtractMin() (min int, ok bool) {
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
	h.heapifyDown(0)

	return
}

func (h *heap) Print() {
	fmt.Println(h.keys)
}

func buildHeap(keys ...int) (h *heap) {
	h = HeapNew(len(keys))
	for _, k := range keys {
		h.Insert(k)
	}
	return
}

func main() {
	h := HeapNew(11)
	h.Insert(3)
	h.Print()
	h.Insert(2)
	h.Print()
	h.Delete(1)
	h.Print()
	h.Insert(15)
	h.Print()
	h.Insert(5)
	h.Print()
	h.Insert(4)
	h.Print()
	h.Insert(45)
	h.Print()
	h.ExtractMin()
	h.Print()
	h.DecreaseKey(2, 1)
	h.Print()
}
