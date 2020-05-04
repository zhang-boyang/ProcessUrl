package heapsort

import (
	"fmt"
	"pucommon"
	"testing"
)
func Test_NewHeap(t *testing.T) {
	heap := NewHeap(10, false)
	for i := 0; i < 20; i++ {
		heap.InsertHeap(&pucommon.KeyValue{"x12",uint64(i)})
	}
	fmt.Println(heap.HeapSlice)
}
