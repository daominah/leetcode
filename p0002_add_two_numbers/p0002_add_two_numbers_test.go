package p0002_add_two_numbers

import (
	"reflect"
	"testing"

	"github.com/daominah/leetcode/structure"
)

func TestAddTwoNumbers(t *testing.T) {
	for i, c := range []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{list1: []int{2, 4, 3}, list2: []int{5, 6, 4}, want: []int{7, 0, 8}},
		{list1: []int{0}, list2: []int{0}, want: []int{0}},
		{list1: []int{9, 9, 9, 9, 9, 9, 9}, list2: []int{9, 9, 9, 9}, want: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	} {
		list1 := structure.NewListNode(c.list1)
		list2 := structure.NewListNode(c.list2)
		got := addTwoNumbers(list1, list2).ToSlice()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error twoSum i %v got %v, want %v", i, got, c.want)
		}
	}
}
