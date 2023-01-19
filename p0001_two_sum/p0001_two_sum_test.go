package p0001_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	for i, c := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	} {
		got := twoSum(c.nums, c.target)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error twoSum i %v got %v, want %v", i, got, c.want)
		}
	}
}
