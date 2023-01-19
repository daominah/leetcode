package structure

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	want := []int{2, 3, 5, 8}
	list := NewListNode(want)
	if got := list.ToSlice(); !reflect.DeepEqual(got, want) {
		t.Errorf("error linked list to slice: %v, want: %v", got, want)
	}

	list2 := NewListNode([]int{})
	//t.Logf("list2: %+v", list2) // Output: nil
	if got := len(list2.ToSlice()); got != 0 {
		t.Errorf("error empty linked list to slice: %v, want: %v", got, 0)
	}
}
