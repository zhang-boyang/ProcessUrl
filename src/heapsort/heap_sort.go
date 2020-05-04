package heapsort

import (
	"pucommon"
)

/* Heap headSize is the capacity of heap,
 * HeapSlice store the kv item, capacity is heapSize
 * currentCnt shows that how many item in HeapSlice now
 * isBigHeap is true means a big heap, false is little heap
 */
type Heap struct {
	heapSize   uint32
	HeapSlice  []pucommon.KeyValue
	currentCnt uint32
	isBigHeap  bool
}

//ROOT is index of heap root
const ROOT uint32 = 0

//NewHeap Alloc a head
func NewHeap(hsize uint32, isbighead bool) *Heap {
	if hsize <= 0 {
		panic("heap size can not <= 0")
	}
	rt := &Heap{
		heapSize:   hsize,
		HeapSlice:  make([]pucommon.KeyValue, hsize),
		currentCnt: 0,
		isBigHeap:  isbighead,
	}
	return rt
}

func (h *Heap) getFather(NodeNo uint32) (uint32, bool) {
	if NodeNo == 0 {
		return 0, false
	}
	fatherNo := (NodeNo - 1) / 2
	if fatherNo >= h.currentCnt {
		return 0, false
	}
	return fatherNo, true

}

func (h *Heap) getLeftChild(NodeNo uint32) (uint32, bool) {
	childNo := NodeNo*2 + 1
	if childNo >= h.currentCnt {
		return 0, false
	}
	return childNo, true
}

func (h *Heap) getRightChild(NodeNo uint32) (uint32, bool) {
	childNo := NodeNo*2 + 2
	if childNo >= h.currentCnt {
		return 0, false
	}
	return childNo, true
}

func (h *Heap) compare(a *pucommon.KeyValue, b *pucommon.KeyValue) bool {
	if h.isBigHeap {
		return a.Value > b.Value
	}
	return a.Value < b.Value
}

func (h *Heap) itemSwap(a *uint32, b *uint32) {
	h.HeapSlice[*a], h.HeapSlice[*b] = h.HeapSlice[*b], h.HeapSlice[*a]
	*a, *b = *b, *a
}

//siftUp when insert a new one into heap and the capacity is less than
// heap size and lift the item up
func (h *Heap) siftUp(curNo uint32) {
	for {
		fno, succ := h.getFather(curNo)
		if !succ {
			break
		}
		if h.compare(&h.HeapSlice[curNo], &h.HeapSlice[fno]) {
			h.itemSwap(&curNo, &fno)
		} else {
			break
		}
	}
}

//siftDown when the root is not suitable for the requirement
//and sink down
func (h *Heap) siftDown() {
	for curNodeNo := ROOT; curNodeNo < h.heapSize; {
		lc, lsuc := h.getLeftChild(curNodeNo)
		rc, rsuc := h.getRightChild(curNodeNo)
		if !lsuc && !rsuc {
			break
		}
		if !lsuc && rsuc {
			if h.compare(&h.HeapSlice[curNodeNo], &h.HeapSlice[rc]) {
				break
			}
			h.itemSwap(&curNodeNo, &rc)
			continue
		}

		if lsuc && !rsuc {
			if h.compare(&h.HeapSlice[curNodeNo], &h.HeapSlice[lc]) {
				break
			}
			h.itemSwap(&curNodeNo, &lc)
			continue
		}
		if h.compare(&h.HeapSlice[lc], &h.HeapSlice[rc]) {
			if h.compare(&h.HeapSlice[curNodeNo], &h.HeapSlice[lc]) {
				break
			}
			h.itemSwap(&curNodeNo, &lc)
		} else {
			if h.compare(&h.HeapSlice[curNodeNo], &h.HeapSlice[rc]) {
				break
			}
			h.itemSwap(&curNodeNo, &rc)
		}
	}
}

//InsertHeap insert an item into heap, maybe cause
// heap illegal and it will be adjusted itself to be a decent heap
func (h *Heap) InsertHeap(item *pucommon.KeyValue) {
	if h.currentCnt < h.heapSize {
		h.HeapSlice[h.currentCnt] = *item
		curNodeNo := h.currentCnt
		h.currentCnt++
		h.siftUp(curNodeNo)
		return
	}
	if !h.compare(&h.HeapSlice[ROOT], item) {
		return
	}
	h.HeapSlice[ROOT] = *item
	h.siftDown()
}
